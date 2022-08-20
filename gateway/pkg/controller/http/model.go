package http

import (
	"fmt"
	"time"

	"github.com/isutare412/meetup/gateway/pkg/core/domain"
	"github.com/isutare412/meetup/gateway/pkg/core/dto"
	"github.com/isutare412/meetup/gateway/pkg/pkgerr"
)

type userPathParams struct {
	UserID int64 `uri:"userId"`
}

func (p *userPathParams) checkUserID() error {
	if p.UserID != 0 {
		return nil
	}
	return pkgerr.Known{Simple: fmt.Errorf("userId required")}
}

type errorResp struct {
	Msg string `json:"msg"`
}

type createUserReq struct {
	Nickname string `json:"nickname"`
}

func (req *createUserReq) validate() error {
	if req.Nickname == "" {
		return pkgerr.Known{Simple: fmt.Errorf("nickname should be given")}
	}
	return nil
}

func (req *createUserReq) intoDTO() *dto.CreateUserReq {
	return &dto.CreateUserReq{
		Nickname: req.Nickname,
	}
}

type createUserResp struct {
	ID        int64     `json:"id"`
	Nickname  string    `json:"nickname"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

func (resp *createUserResp) fromUser(user *domain.User) {
	resp.ID = user.ID
	resp.Nickname = *user.Nickname
	resp.CreatedAt = user.CreatedAt
	resp.UpdatedAt = user.UpdatedAt
}

type getUserResp struct {
	ID        int64     `json:"id"`
	Nickname  string    `json:"nickname"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

func (resp *getUserResp) fromUser(user *domain.User) {
	resp.ID = user.ID
	resp.Nickname = *user.Nickname
	resp.CreatedAt = user.CreatedAt
	resp.UpdatedAt = user.UpdatedAt
}

type deleteUserResp struct {
	ID int64 `json:"id"`
}
