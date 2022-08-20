package postgres

import (
	"context"
	"fmt"

	"github.com/isutare412/meetup/gateway/pkg/core/domain"
	"github.com/isutare412/meetup/gateway/pkg/pkgerr"
)

type UserRepository struct {
	cli *Client
}

func NewUserRepository(cli *Client) *UserRepository {
	return &UserRepository{
		cli: cli,
	}
}

func (r *UserRepository) Create(ctx context.Context, user *domain.User) error {
	db := r.cli.extractTxOrDB(ctx)

	if err := db.Create(user).Error; err != nil {
		if k, v := isErrDuplicateKey(err); k != "" {
			return pkgerr.Known{
				Origin: err,
				Simple: fmt.Errorf("cannot create user due to duplicate(%s=%s)", k, v),
			}
		}
		return pkgerr.Known{
			Origin: err,
			Simple: fmt.Errorf("failed to create user"),
		}
	}
	return nil
}

func (r *UserRepository) GetByID(ctx context.Context, id int64) (*domain.User, error) {
	db := r.cli.extractTxOrDB(ctx)

	var user = domain.User{ID: id}
	if err := db.First(&user).Error; err != nil {
		if isErrRecordNotFound(err) {
			return nil, pkgerr.Known{
				Simple: fmt.Errorf("user(id=%d) not found", id),
			}
		}
		return nil, err
	}
	return &user, nil
}

func (r *UserRepository) DeleteByID(ctx context.Context, id int64) error {
	db := r.cli.extractTxOrDB(ctx)

	res := db.Delete(&domain.User{ID: id})
	if err := res.Error; err != nil {
		return pkgerr.Known{
			Origin: err,
			Simple: fmt.Errorf("cannot delete user"),
		}
	}
	if res.RowsAffected != 1 {
		return pkgerr.Known{
			Simple: fmt.Errorf("user(id=%d) not found", id),
		}
	}
	return nil
}
