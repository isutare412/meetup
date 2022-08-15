package postgres

import (
	"context"
	"errors"
	"fmt"

	"github.com/isutare412/meetup/gateway/pkg/core/domain"
	pkgerr "github.com/isutare412/meetup/gateway/pkg/error"
	"gorm.io/gorm"
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
		return err
	}
	return nil
}

func (r *UserRepository) GetByID(ctx context.Context, id int64) (*domain.User, error) {
	db := r.cli.extractTxOrDB(ctx)

	var user = domain.User{ID: id}
	if err := db.First(&user).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, pkgerr.Known{
				Errno:  pkgerr.ErrnoEntityNotFound,
				Source: fmt.Errorf("user(id=%d) not found", id),
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
		return err
	}
	if res.RowsAffected != 1 {
		return pkgerr.Known{
			Errno:  pkgerr.ErrnoEntityNotFound,
			Source: fmt.Errorf("user(id=%d) not found", id),
		}
	}
	return nil
}
