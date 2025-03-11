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
		r.Post("/", s.withPermissions(s.handleCreateEvent(), "event-edit"))
		r.Put("/{eventID}", s.withPermissions(s.handleUpdateEvent(), "event-edit"))
		r.Get("/", s.handleListEvents())
		r.Get("/{eventID}", s.handleEventByID())
		r.Post("/image", s.withPermissions(s.handleUploadEventCoverImage(), "event-edit"))
	})

	s.mux.Route("/artists", func(r chi.Router) {
		r.Get("/{artistID}", s.handleGetArtistByID())
		r.Get("/", s.handleListArtists())
		r.Post("/", s.withPermissions(s.handleCreateArtist(), "artist-edit"))
		r.Put("/{artistID}", s.withPermissions(s.handleUpdateArtist(), "artist-edit"))
		r.Delete("/{artistID}", s.withPermissions(s.handleDeleteArtist(), "artist-delete"))
		r.Put("/image", s.withPermissions(s.handleUploadArtistImage(), "artist-edit"))
	})

	s.mux.Route("/venues", func(r chi.Router) {
		r.Get("/", s.withPermissions(s.handleListVenues(), "venue-list"))
		r.Post("/", s.withPermissions(s.handleCreateVenue(), "venue-edit"))
		r.Delete("/{venueID}", s.withPermissions(s.handleDeleteVenue(), "venue-delete"))
	})

	s.mux.Route("/users", func(r chi.Router) {
		r.Get("/roles/{userID}", s.withPermissions(s.handleListUserRoles(), "role-list"))
	})

	s.mux.Route("/genres", func(r chi.Router) {
		r.Post("/", s.withPermissions(s.handleCreateGenre(), "genre-edit"))
		r.Get("/", s.handleListGenres())
	})
}
