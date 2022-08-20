package http

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/isutare412/meetup/gateway/pkg/core/domain"
	"github.com/isutare412/meetup/gateway/pkg/core/port"
	"github.com/isutare412/meetup/gateway/pkg/logger"
)

func createUser(uSvc port.UserService) gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx := c.Request.Context()

		var req createUserReq
		if err := c.ShouldBindJSON(&req); err != nil {
			logger.S().Error(err)
			responseError(c, http.StatusBadRequest, err)
			return
		}
		if err := req.validate(); err != nil {
			logger.S().Error(err)
			responseError(c, http.StatusBadRequest, err)
			return
		}

		var user *domain.User
		user, err := uSvc.Create(ctx, req.intoDTO())
		if err != nil {
			logger.S().Error(err)
			responseError(c, http.StatusInternalServerError, err)
			return
		}

		var resp createUserResp
		resp.fromUser(user)
		c.JSON(http.StatusCreated, &resp)
	}
}

func getUser(uSvc port.UserService) gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx := c.Request.Context()

		var params userPathParams
		if err := c.ShouldBindUri(&params); err != nil {
			logger.S().Error(err)
			responseError(c, http.StatusBadRequest, err)
			return
		}
		if err := params.checkUserID(); err != nil {
			logger.S().Error(err)
			responseError(c, http.StatusBadRequest, err)
			return
		}

		var user *domain.User
		user, err := uSvc.GetByID(ctx, params.UserID)
		if err != nil {
			logger.S().Error(err)
			responseError(c, http.StatusInternalServerError, err)
			return
		}

		var resp getUserResp
		resp.fromUser(user)
		c.JSON(http.StatusOK, &resp)
	}
}

func deleteUser(uSvc port.UserService) gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx := c.Request.Context()

		var params userPathParams
		if err := c.ShouldBindUri(&params); err != nil {
			logger.S().Error(err)
			responseError(c, http.StatusBadRequest, err)
			return
		}
		if err := params.checkUserID(); err != nil {
			logger.S().Error(err)
			responseError(c, http.StatusBadRequest, err)
			return
		}

		if err := uSvc.DeleteByID(ctx, params.UserID); err != nil {
			logger.S().Error(err)
			responseError(c, http.StatusInternalServerError, err)
			return
		}

		var resp = deleteUserResp{ID: params.UserID}
		c.JSON(http.StatusOK, &resp)
	}
}

func playground(uSvc port.UserService) gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Param("id")
		if id == "foo" {
			err := fmt.Errorf("foo is invalid path parameter")
			logger.S().Error(err)
			responseError(c, http.StatusBadRequest, err)
			return
		}
		c.Status(http.StatusNoContent)
	}
}
