package http

import (
	"github.com/gin-gonic/gin"
	"github.com/isutare412/meetup/gateway/pkg/pkgerr"
)

const (
	ctxKeyError = "ctx-key-error"
)

func responseError(c *gin.Context, code int, err error) {
	injectError(c, err)

	var errMsg = err.Error()
	if kerr := pkgerr.AsKnown(err); kerr != nil {
		errMsg = kerr.SimpleError()
	}
	c.JSON(code, &errorResp{
		Msg: errMsg,
	})
}

func injectError(c *gin.Context, err error) {
	c.Set(ctxKeyError, err)
}

func extractError(c *gin.Context) error {
	if err, exists := c.Get(ctxKeyError); exists {
		return err.(error)
	}
	return nil
}
