package server

import (
	"fmt"
	"log/slog"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/cors"
	"github.com/mattismoel/konnekt/internal/cfg"
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

	contentService *service.ContentService
	authService    *service.AuthService
	teamService    *service.TeamService
	eventService   *service.EventService
	artistService  *service.ArtistService
	memberService  *service.MemberService
	venueService   *service.VenueService
}

func New(cfgs ...cfg.Func[Server]) (*Server, error) {
	s := &Server{
		mux: chi.NewMux(),
	}

	for _, cfg := range cfgs {
		if err := cfg(s); err != nil {
			return nil, fmt.Errorf("Could not use server config: %v", err)
		}
	}

	s.setupRoutes()

	return s, nil
}

func WithCORSOrigins(allowedOrigins ...string) cfg.Func[Server] {
	return func(s *Server) error {
		s.mux.Use(cors.Handler(cors.Options{
			AllowedOrigins:   allowedOrigins,
			AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
			AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
			ExposedHeaders:   []string{"Link"},
			AllowCredentials: true,
			MaxAge:           300,
		}))

		return nil
	}
}

func WithContentService(contentService *service.ContentService) cfg.Func[Server] {
	return func(s *Server) error {
		s.contentService = contentService
		return nil
	}
}

func WithTeamService(teamService *service.TeamService) cfg.Func[Server] {
	return func(s *Server) error {
		s.teamService = teamService
		return nil
	}
}

func WithAuthService(authService *service.AuthService) cfg.Func[Server] {
	return func(s *Server) error {
		s.authService = authService
		return nil
	}
}

func WithEventService(eventService *service.EventService) cfg.Func[Server] {
	return func(s *Server) error {
		s.eventService = eventService
		return nil
	}
}

func WithArtistService(artistService *service.ArtistService) cfg.Func[Server] {
	return func(s *Server) error {
		s.artistService = artistService
		return nil
	}
}

func WithMemberService(memberService *service.MemberService) cfg.Func[Server] {
	return func(s *Server) error {
		s.memberService = memberService
		return nil
	}
}

func WithVenueService(venueService *service.VenueService) cfg.Func[Server] {
	return func(s *Server) error {
		s.venueService = venueService
		return nil
	}
}

func WithAddress(addr string) cfg.Func[Server] {
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
