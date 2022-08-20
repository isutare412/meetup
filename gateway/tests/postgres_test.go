package tests

import (
	"context"
	"testing"

	"github.com/isutare412/meetup/gateway/pkg/config"
	"github.com/isutare412/meetup/gateway/pkg/core/domain"
	"github.com/isutare412/meetup/gateway/pkg/infrastructure/postgres"
	"github.com/stretchr/testify/require"
)

func TestPostgresCRUD(t *testing.T) {
	cfg, err := loadTestConfig()
	require.NoError(t, err)

	var cli *postgres.Client
	cli, err = postgres.NewClient(context.Background(), cfg.Postgres)
	require.NoError(t, err)

	err = cli.MigrateSchema(context.Background())
	require.NoError(t, err)

	var userRepo = postgres.NewUserRepository(cli)

	ctxWithTx, commit, _ := cli.BeginTx(context.Background())

	userName := "redshore"
	var user = domain.User{Nickname: &userName}
	err = userRepo.Create(ctxWithTx, &user)
	require.NoError(t, err)
	require.NotZero(t, user.ID)

	var foundUser *domain.User
	foundUser, err = userRepo.GetByID(ctxWithTx, user.ID)
	require.NoError(t, err)
	require.Equal(t, &user, foundUser)

	err = userRepo.DeleteByID(ctxWithTx, user.ID)
	require.NoError(t, err)

	err = commit()
	require.NoError(t, err)
}

func loadTestConfig() (*config.Config, error) {
	return config.Load("../configs/local/config.yaml")
}
