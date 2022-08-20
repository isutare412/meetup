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
	Nickname string `json:"nickname" example:"redshore"`
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
	ID        int64     `json:"id" example:"412"`
	Nickname  string    `json:"nickname" example:"redshore"`
	CreatedAt time.Time `json:"createdAt" example:"2022-08-20T18:54:53.965295+09:00"`
	UpdatedAt time.Time `json:"updatedAt" example:"2022-08-20T18:54:53.965295+09:00"`
}

func (resp *createUserResp) fromUser(user *domain.User) {
	resp.ID = user.ID
	resp.Nickname = *user.Nickname
	resp.CreatedAt = user.CreatedAt
	resp.UpdatedAt = user.UpdatedAt
}

type getUserResp struct {
	ID        int64     `json:"id" example:"412"`
	Nickname  string    `json:"nickname" example:"redshore"`
	CreatedAt time.Time `json:"createdAt" example:"2022-08-20T18:54:53.965295+09:00"`
	UpdatedAt time.Time `json:"updatedAt" example:"2022-08-20T18:54:53.965295+09:00"`
}

func (resp *getUserResp) fromUser(user *domain.User) {
	resp.ID = user.ID
	resp.Nickname = *user.Nickname
	resp.CreatedAt = user.CreatedAt
	resp.UpdatedAt = user.UpdatedAt
}

type deleteUserResp struct {
	ID int64 `json:"id" example:"412"`
}
