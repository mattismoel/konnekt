package server

import (
	"log/slog"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/cors"
	"github.com/mattismoel/konnekt/internal/service"
)

type ListReponse struct {
	Page       int `json:"page"`
	PerPage    int `json:"perPage"`
	PageCount  int `json:"pageCount"`
	TotalCount int `json:"totalCount"`
	Records    any `json:"records"`
}

type Server struct {
	mux  *chi.Mux
	addr string

	authService   *service.AuthService
	eventService  *service.EventService
	artistService *service.ArtistService
	userService   *service.UserService
	venueService  *service.VenueService
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

func WithCORSOrigins(allowedOrigins ...string) CfgFunc {
	return func(s *Server) error {
		s.mux.Use(cors.Handler(cors.Options{
			AllowedOrigins:   allowedOrigins,
			AllowCredentials: true,
			AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE"},
		}))

		return nil
	}
}

func WithAuthService(authService *service.AuthService) CfgFunc {
	return func(s *Server) error {
		s.authService = authService
		return nil
	}
}

func WithEventService(eventService *service.EventService) CfgFunc {
	return func(s *Server) error {
		s.eventService = eventService
		return nil
	}
}

func WithArtistService(artistService *service.ArtistService) CfgFunc {
	return func(s *Server) error {
		s.artistService = artistService
		return nil
	}
}

func WithUserService(userService *service.UserService) CfgFunc {
	return func(s *Server) error {
		s.userService = userService
		return nil
	}
}

func WithVenueService(venueService *service.VenueService) CfgFunc {
	return func(s *Server) error {
		s.venueService = venueService
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
