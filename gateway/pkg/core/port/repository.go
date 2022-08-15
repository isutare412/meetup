package port

import (
	"context"
	"database/sql"

	"github.com/isutare412/meetup/gateway/pkg/core/domain"
)

type RepositorySession interface {
	BeginTx(ctx context.Context, opts ...*sql.TxOptions) (ctxWithTx context.Context, commit, rollback func() error)
}

type UserRepository interface {
	Create(ctx context.Context, user *domain.User) error
	GetByID(ctx context.Context, id int64) (*domain.User, error)
	DeleteByID(ctx context.Context, id int64) error
}
