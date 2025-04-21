package server

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func (s *Server) setupRoutes() {
	s.mux.Use(middleware.Logger)

	s.mux.Route("/members", func(r chi.Router) {
		r.Get("/", s.withPermissions(s.handleListMembers(), "view:member", "view:role", "view:permission"))
		r.Post("/{memberID}/approve", s.withPermissions(s.handleApproveMember(), "edit:member"))
		r.Delete("/{memberID}", s.withPermissions(s.handleDeleteMember(), "delete:member"))
	})

	s.mux.Route("/auth", func(r chi.Router) {
		r.Post("/login", s.handleLogin())
		r.Post("/register", s.handleRegister())
		r.Post("/log-out", s.handleLogOut())
		r.Get("/session", s.handleGetSession())

		r.Route("/roles", func(r chi.Router) {
			r.Get("/", s.withPermissions(s.handleListRoles(), "view:role"))
			r.Get("/{memberID}", s.withPermissions(s.handleListMemberRoles(), "view:role"))
			r.Post("/", s.withPermissions(s.handleCreateRole(), "edit:role"))
			r.Delete("/{roleID}", s.withPermissions(s.handleDeleteRole(), "delete:role"))
		})

		r.Route("/permissions", func(r chi.Router) {
			r.Get("/", s.withPermissions(s.handleListPermissions(), "view:permission"))
			r.Get("/{memberID}", s.withPermissions(s.handleListMemberPermissions(), "view:permission"))
		})
	})

	s.mux.Route("/events", func(r chi.Router) {
		r.Post("/", s.withPermissions(s.handleCreateEvent(), "edit:event"))
		r.Put("/{eventID}", s.withPermissions(s.handleUpdateEvent(), "edit:event"))
		r.Delete("/{eventID}", s.withPermissions(s.handleDeleteEvent(), "delete:event"))
		r.Get("/", s.handleListEvents())
		r.Get("/{eventID}", s.handleEventByID())
		r.Post("/image", s.withPermissions(s.handleUploadEventImage(), "edit:event"))
	})

	s.mux.Route("/artists", func(r chi.Router) {
		r.Get("/{artistID}", s.handleGetArtistByID())
		r.Get("/", s.handleListArtists())
		r.Post("/", s.withPermissions(s.handleCreateArtist(), "edit:artist"))
		r.Put("/{artistID}", s.withPermissions(s.handleUpdateArtist(), "edit:artist"))
		r.Delete("/{artistID}", s.withPermissions(s.handleDeleteArtist(), "delete:artist"))
		r.Put("/image", s.withPermissions(s.handleUploadArtistImage(), "edit:artist"))
	})

	s.mux.Route("/venues", func(r chi.Router) {
		r.Get("/", s.withPermissions(s.handleListVenues(), "view:venue"))
		r.Post("/", s.withPermissions(s.handleCreateVenue(), "edit:venue"))
		r.Put("/{venueID}", s.withPermissions(s.handleUpdateVenue(), "edit:venue"))
		r.Delete("/{venueID}", s.withPermissions(s.handleDeleteVenue(), "delete:venue"))
	})

	s.mux.Route("/genres", func(r chi.Router) {
		r.Post("/", s.withPermissions(s.handleCreateGenre(), "edit:genre"))
		r.Get("/", s.handleListGenres())
	})
}
