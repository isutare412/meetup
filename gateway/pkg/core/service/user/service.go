package user

import (
	"context"

	"github.com/isutare412/meetup/gateway/pkg/core/domain"
	"github.com/isutare412/meetup/gateway/pkg/core/dto"
	"github.com/isutare412/meetup/gateway/pkg/core/port"
	perror "github.com/isutare412/meetup/gateway/pkg/error"
)

type Service struct {
	repoSession port.RepositorySession
	userRepo    port.UserRepository
}

func (s *Service) Create(ctx context.Context, req *dto.CreateUserReq) (user *domain.User, err error) {
	if err := req.IsValid(); err != nil {
		return nil, err
	}

	ctxWithTx, commit, rollback := s.repoSession.BeginTx(ctx)
	defer func() { err = perror.TryRollback(err, rollback) }()

	user = req.IntoUser()
	if err := s.userRepo.Create(ctxWithTx, user); err != nil {
		return nil, err
	}

	if err := commit(); err != nil {
		return nil, err
	}
	return user, nil
}

func (s *Service) GetByID(ctx context.Context, id int64) (*domain.User, error) {
	var user *domain.User
	user, err := s.userRepo.GetByID(ctx, id)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (s *Service) DeleteByID(ctx context.Context, id int64) error {
	if err := s.userRepo.DeleteByID(ctx, id); err != nil {
		return err
	}
	return nil
}
