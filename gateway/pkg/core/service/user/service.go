package user

import (
	"context"
	"fmt"

	"github.com/isutare412/meetup/gateway/pkg/core/domain"
	"github.com/isutare412/meetup/gateway/pkg/core/dto"
	"github.com/isutare412/meetup/gateway/pkg/core/port"
	"github.com/isutare412/meetup/gateway/pkg/pkgerr"
)

type Service struct {
	repoSess port.RepositorySession
	userRepo port.UserRepository
}

func NewService(repoSession port.RepositorySession, userRepo port.UserRepository) *Service {
	return &Service{
		repoSess: repoSession,
		userRepo: userRepo,
	}
}

func (s *Service) Create(ctx context.Context, req *dto.CreateUserReq) (user *domain.User, err error) {
	if err := req.Validate(); err != nil {
		return nil, err
	}

	ctxWithTx, commit, rollback := s.repoSess.BeginTx(ctx)
	defer func() { err = pkgerr.TryRollback(err, rollback) }()

	user = req.IntoUser()
	if err := s.userRepo.Create(ctxWithTx, user); err != nil {
		return nil, fmt.Errorf("creating user: %w", err)
	}

	if err := commit(); err != nil {
		return nil, err
	}
	return user, nil
}

func (s *Service) GetByID(ctx context.Context, id int64) (user *domain.User, err error) {
	ctxWithTx, commit, rollback := s.repoSess.BeginTx(ctx)
	defer func() { err = pkgerr.TryRollback(err, rollback) }()

	user, err = s.userRepo.GetByID(ctxWithTx, id)
	if err != nil {
		return nil, fmt.Errorf("getting user by id(%d): %w", id, err)
	}

	if err := commit(); err != nil {
		return nil, err
	}
	return user, nil
}

func (s *Service) DeleteByID(ctx context.Context, id int64) (err error) {
	ctxWithTx, commit, rollback := s.repoSess.BeginTx(ctx)
	defer func() { err = pkgerr.TryRollback(err, rollback) }()

	if err := s.userRepo.DeleteByID(ctxWithTx, id); err != nil {
		return fmt.Errorf("deleting user by id(%d): %w", id, err)
	}

	if err := commit(); err != nil {
		return err
	}
	return nil
}
