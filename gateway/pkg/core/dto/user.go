package dto

import (
	"fmt"

	"github.com/isutare412/meetup/gateway/pkg/core/domain"
)

type CreateUserReq struct {
	Nickname string
}

func (r *CreateUserReq) IsValid() error {
	if r.Nickname == "" {
		return fmt.Errorf("empty nickname is not allowed")
	}
	return nil
}

func (r *CreateUserReq) IntoUser() *domain.User {
	return &domain.User{
		Nickname: &r.Nickname,
	}
}
