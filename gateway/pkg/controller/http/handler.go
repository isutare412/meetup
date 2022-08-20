package http

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/isutare412/meetup/gateway/pkg/core/domain"
	"github.com/isutare412/meetup/gateway/pkg/core/port"
	"github.com/isutare412/meetup/gateway/pkg/logger"
)

// @Tags        Users
// @Description Create an user.
// @Router      /api/v1/users [POST]
// @Param       request body createUserReq true "Request to create an user"
// @Accept      json
// @Produce     json
// @Success     200     {object} createUserResp
// @Failure     default {object} errorResp "Somethings got wrong"
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
		c.JSON(http.StatusOK, &resp)
	}
}

// @Tags        Users
// @Description Get an user.
// @Router      /api/v1/users/{userId} [GET]
// @Param       userId path number true "User ID"
// @Produce     json
// @Success     200     {object} getUserResp
// @Failure     default {object} errorResp "Somethings got wrong"
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

// @Tags        Users
// @Description Delete an user.
// @Router      /api/v1/users/{userId} [DELETE]
// @Param       userId path number true "User ID"
// @Produce     json
// @Success     200     {object} deleteUserResp
// @Failure     default {object} errorResp "Somethings got wrong"
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
