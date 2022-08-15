package http

import (
	"time"

	"github.com/gin-gonic/gin"
	"github.com/isutare412/meetup/gateway/pkg/logger"
	"go.uber.org/zap"
)

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

	if errAny, exists := c.Get(ctxKeyError); exists {
		log = log.With(zap.NamedError("errorMsg", errAny.(error)))
	}

	log.Info("HTTP access")
}
