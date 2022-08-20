package http

import (
	"context"
	"errors"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	_ "github.com/isutare412/meetup/gateway/api"
	"github.com/isutare412/meetup/gateway/pkg/config"
	"github.com/isutare412/meetup/gateway/pkg/core/port"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

type Server struct {
	srv *http.Server
}

func NewServer(cfg *config.HTTPServerConfig, uSvc port.UserService) *Server {
	gin.SetMode(gin.ReleaseMode)
	root := gin.New()
	root.Use(recovery)
	root.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	api := root.Group("/api/v1")
	api.Use(accessLog)

	api.POST("/users", createUser(uSvc))
	api.GET("/users/:userId", getUser(uSvc))
	api.DELETE("/users/:userId", deleteUser(uSvc))

	api.GET("/playground/:id", playground(uSvc))

	return &Server{
		srv: &http.Server{
			Addr:    fmt.Sprintf("%s:%d", cfg.Host, cfg.Port),
			Handler: root,
		},
	}
}

func (s *Server) Addr() string {
	return s.srv.Addr
}

func (s *Server) Run() <-chan error {
	fails := make(chan error, 1)
	go func() {
		err := s.srv.ListenAndServe()
		if err != nil && !errors.Is(err, http.ErrServerClosed) {
			fails <- fmt.Errorf("listening http server: %w", err)
		}
	}()
	return fails
}

func (s *Server) Shutdown(ctx context.Context) error {
	return s.srv.Shutdown(ctx)
}
