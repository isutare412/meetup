package main

import (
	"context"
	"fmt"

	"github.com/isutare412/meetup/gateway/pkg/config"
	"github.com/isutare412/meetup/gateway/pkg/controller/http"
	"github.com/isutare412/meetup/gateway/pkg/core/service/user"
	"github.com/isutare412/meetup/gateway/pkg/infrastructure/postgres"
	"github.com/isutare412/meetup/gateway/pkg/logger"
)

type components struct {
	pgClient   *postgres.Client
	httpServer *http.Server
}

func newComponents(ctx context.Context, cfg *config.Config) (*components, error) {
	success := make(chan *components, 1)
	fails := make(chan error, 1)
	go func() {
		logger.S().Info("Dependency injection start")

		var pgClient *postgres.Client
		pgClient, err := postgres.NewClient(cfg.Postgres)
		if err != nil {
			fails <- err
			return
		}
		userRepo := postgres.NewUserRepository(pgClient)

		userService := user.NewService(pgClient, userRepo)

		var httpServer = http.NewServer(cfg.Server.HTTP, userService)

		logger.S().Info("Dependency injection complete")
		success <- &components{
			pgClient:   pgClient,
			httpServer: httpServer,
		}
	}()

	var components *components
	select {
	case <-ctx.Done():
		return nil, fmt.Errorf("time out while injecting dependencies")
	case err := <-fails:
		return nil, err
	case components = <-success:
	}
	return components, nil
}

func (c *components) init(ctx context.Context) error {
	logger.S().Info("Component initialization start")

	if err := c.pgClient.MigrateSchema(ctx); err != nil {
		return err
	}
	logger.S().Info("DB schema migration complete")

	logger.S().Info("Component initialization complete")
	return nil
}

func (c *components) run(ctx context.Context) <-chan error {
	failMux := make(chan error)
	go func() {
		httpServerFails := c.httpServer.Run()
		logger.S().Infof("Run http server at %s", c.httpServer.Addr())

		select {
		case err := <-httpServerFails:
			failMux <- err
		}
	}()
	return failMux
}

func (c *components) shutdown(ctx context.Context) {
	logger.S().Info("Graceful shutdown start")

	if err := c.httpServer.Shutdown(ctx); err != nil {
		logger.S().Errorf("Error while shutting down http server: %v", err)
	}

	logger.S().Info("Graceful shutdown complete")
}
