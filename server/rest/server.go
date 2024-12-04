package rest

import (
	"net"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
)

type server struct {
	server *http.Server
	mux    *chi.Mux

	eventService EventService
	userService  UserService
}

func NewServer(cfg Cfg) (*server, error) {
	if err := cfg.Validate(); err != nil {
		return nil, err
	}

	addr := net.JoinHostPort(cfg.Host, strconv.Itoa(cfg.Port))

	s := &server{
		server: &http.Server{
			Addr: addr,
		},
		mux: chi.NewRouter(),

		eventService: cfg.EventService,
		userService:  cfg.UserService,
	}

	s.server.Handler = http.HandlerFunc(s.mux.ServeHTTP)

	s.mux.Mount("/events", s.createEventsRoutes())
	s.mux.Mount("/auth", s.createAuthRoutes())
	s.mux.Mount("/users", s.createUserRoutes())

	return s, nil
}

func (s server) Start() error {
	return s.server.ListenAndServe()
}
