package server

import (
	"context"
	"net/http"
	"time"

	"github.com/begenov/real-time-forum/internal/config"
)

type Server struct {
	httpServer *http.Server
}

func NewServer(cfg *config.Config, handler http.Handler) *Server {

	return &Server{
		httpServer: &http.Server{
			Addr:           ":" + cfg.Server.Port,
			Handler:        handler,
			ReadTimeout:    time.Duration(cfg.Server.ReadTimeout) * time.Second,
			WriteTimeout:   time.Duration(cfg.Server.WriteTimeout) * time.Second,
			MaxHeaderBytes: cfg.Server.MaxHeaderMegabytes << 20,
		},
	}
}

func (s *Server) Run() error {
	return s.httpServer.ListenAndServe()
}

func (s *Server) Stop(ctx context.Context) error {
	return s.httpServer.Shutdown(ctx)
}
