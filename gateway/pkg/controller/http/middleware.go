package http

import (
	"errors"
	"net"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/isutare412/meetup/gateway/pkg/logger"
	"go.uber.org/zap"
)

func recovery(c *gin.Context) {
	defer func() {
		err := recover()
		if err == nil {
			return
		}

		// Check for a broken connection, as it is not really a condition that
		// warrants a panic stack trace.
		var brokenPipe bool
		if ne, ok := err.(*net.OpError); ok {
			var se *os.SyscallError
			if errors.As(ne, &se) {
				if strings.Contains(strings.ToLower(se.Error()), "broken pipe") ||
					strings.Contains(strings.ToLower(se.Error()), "connection reset by peer") {
					brokenPipe = true
				}
			}
		}

		logger.S().With("panic", err).Error("HTTP recovered from panic")
		if brokenPipe {
			// If the connection is dead, we can't write a status to it.
			c.Abort()
		} else {
			c.AbortWithStatus(http.StatusInternalServerError)
		}
	}()
	c.Next()
}

func accessLog(c *gin.Context) {
	start := time.Now()

	c.Next()

	log := logger.A().With(
		zap.String("remoteAddr", c.ClientIP()),
		zap.Stringer("url", c.Request.URL),
		zap.String("method", c.Request.Method),
		zap.String("contentType", c.ContentType()),
		zap.Int64("contentLength", c.Request.ContentLength),
		zap.String("httpVersion", c.Request.Proto),
		zap.String("userAgent", c.Request.UserAgent()),
		zap.Int("status", c.Writer.Status()),
		zap.Int64("elapsedTime", time.Since(start).Milliseconds()),
	)

	if err := extractError(c); err != nil {
		log = log.With(zap.NamedError("errorMsg", err))
	}

	log.Info("HTTP access")
}
