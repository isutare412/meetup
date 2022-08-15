package port

import (
	"context"

	"github.com/isutare412/meetup/gateway/pkg/core/domain"
	"github.com/isutare412/meetup/gateway/pkg/core/dto"
)

type UserService interface {
	Create(ctx context.Context, req *dto.CreateUserReq) (*domain.User, error)
	GetByID(ctx context.Context, id int64) (*domain.User, error)
	DeleteByID(ctx context.Context, id int64) error
}
