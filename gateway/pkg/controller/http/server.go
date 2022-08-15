package http

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/isutare412/meetup/gateway/pkg/config"
)

type Server struct {
	srv *http.Server
}

func NewServer(cfg *config.HTTPServerConfig) *Server {
	gin.SetMode(gin.ReleaseMode)
	root := gin.New()

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
