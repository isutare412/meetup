package dto

import (
	"fmt"

	"github.com/isutare412/meetup/gateway/pkg/core/domain"
	"github.com/isutare412/meetup/gateway/pkg/pkgerr"
)

type CreateUserReq struct {
	Nickname string
}

func (r *CreateUserReq) Validate() error {
	if r.Nickname == "" {
		return pkgerr.Known{Simple: fmt.Errorf("empty nickname is not allowed")}
	}
	return nil
}

func (r *CreateUserReq) IntoUser() *domain.User {
	return &domain.User{
		Nickname: &r.Nickname,
	}
}
