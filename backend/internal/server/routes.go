package server

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func (s *Server) setupRoutes() {
	s.mux.Use(middleware.Logger)

	s.mux.Route("/auth", func(r chi.Router) {
		r.Post("/login", s.handleLogin())
		r.Post("/register", s.handleRegister())
		r.Post("/log-out", s.handleLogOut())
		r.Get("/session", s.handleGetSession())
	})

	s.mux.Route("/events", func(r chi.Router) {
		r.Post("/", s.withPermissions(s.handleCreateEvent(), "event-create"))
		r.Get("/", s.handleListEvents())
		r.Get("/{eventID}", s.handleEventByID())
	})

	s.mux.Route("/artists", func(r chi.Router) {
		r.Get("/", s.withPermissions(s.handleListArtists(), "artist-list"))
		r.Post("/", s.withPermissions(s.handleCreateArtist(), "artist-create"))
	})

	s.mux.Route("/venues", func(r chi.Router) {
		r.Get("/", s.withPermissions(s.handleListVenues(), "venue-list"))
	})

	s.mux.Route("/users", func(r chi.Router) {
		r.Get("/roles/{userID}", s.withPermissions(s.handleListUserRoles(), "role-list"))
	})
}
