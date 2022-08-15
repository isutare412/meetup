package http

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/isutare412/meetup/gateway/pkg/logger"
)

func playground(c *gin.Context) {
	id := c.Param("id")
	if id == "foo" {
		err := fmt.Errorf("foo is invalid path parameter")
		logger.S().Error(err)
		responseError(c, http.StatusBadRequest, err)
		return
	}

	c.Status(http.StatusNoContent)
}
