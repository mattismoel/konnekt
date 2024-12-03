package rest

import (
	"context"
	"fmt"
	"net"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
	"github.com/mattismoel/konnekt"
)

type EventService interface {
	FindEventByID(context.Context, int64) (konnekt.Event, error)
	FindEvents(context.Context, konnekt.EventFilter) ([]konnekt.Event, error)
	CreateEvent(context.Context, konnekt.Event) (konnekt.Event, error)
	UpdateEvent(context.Context, int64, konnekt.EventUpdate) (konnekt.Event, error)
	DeleteEvent(context.Context, int64) error
}

type server struct {
	server *http.Server
	mux    *chi.Mux

	eventService EventService
}

func NewServer(cfg Cfg) (*server, error) {
	if err := cfg.Validate(); err != nil {
		return nil, err
	}

	addr := net.JoinHostPort(cfg.Host, strconv.Itoa(cfg.Port))
	fmt.Println("ADDR", addr)

	s := &server{
		server: &http.Server{
			Addr: addr,
		},
		mux: chi.NewRouter(),

		eventService: cfg.EventService,
	}

	s.server.Handler = http.HandlerFunc(s.mux.ServeHTTP)

	s.mux.Mount("/events", s.createEventsRoutes())
	s.mux.Mount("/auth", s.createAuthRoutes())

	return s, nil
}

func (s server) Start() error {
	return s.server.ListenAndServe()
}
