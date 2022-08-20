package http

import (
	"context"
	"errors"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/isutare412/meetup/gateway/pkg/config"
	"github.com/isutare412/meetup/gateway/pkg/core/port"
)

type Server struct {
	srv *http.Server
}

func NewServer(cfg *config.HTTPServerConfig, uSvc port.UserService) *Server {
	gin.SetMode(gin.ReleaseMode)
	root := gin.New()
	root.Use(recovery)

	api := root.Group("/api/v1")
	api.Use(accessLog)
	api.GET("/playground/:id", playground)

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
