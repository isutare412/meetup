package http

import "github.com/gin-gonic/gin"

const (
	ctxKeyError = "ctx-key-error"
)

func responseError(c *gin.Context, code int, err error) {
	injectError(c, err)
	c.JSON(code, &errorResp{
		Msg: err.Error(),
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
