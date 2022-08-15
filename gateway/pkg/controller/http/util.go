package http

import "github.com/gin-gonic/gin"

func responseError(c *gin.Context, code int, err error) {
	c.Set(ctxKeyError, err)
	c.JSON(code, &errorResp{
		Msg: err.Error(),
	})
}
