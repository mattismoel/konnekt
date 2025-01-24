package server

import (
	"log/slog"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/mattismoel/konnekt/internal/service"
)

type Server struct {
	mux  *chi.Mux
	addr string

	authService *service.AuthService
}

type CfgFunc func(s *Server) error

func New(cfgs ...CfgFunc) (*Server, error) {
	s := &Server{
		mux: chi.NewMux(),
	}

	for _, cfg := range cfgs {
		if err := cfg(s); err != nil {
			return nil, err
		}
	}

	s.setupRoutes()

	return s, nil
}

func WithAuthService(authService *service.AuthService) CfgFunc {
	return func(s *Server) error {
		s.authService = authService
		return nil
	}
}

func WithAddress(addr string) CfgFunc {
	return func(s *Server) error {
		s.addr = addr
		return nil
	}
}

func (srv Server) Start() error {
	slog.Info("Server started", "address", srv.addr)
	httpServer := http.Server{
		Addr:    srv.addr,
		Handler: srv.mux,
	}

	return httpServer.ListenAndServe()
}
