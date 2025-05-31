package server

import (
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func (s *Server) setupRoutes() {
	s.mux.Use(middleware.Logger)
	s.mux.Use(middleware.RequestID)
	s.mux.Use(middleware.RealIP)
	s.mux.Use(middleware.Recoverer)
	s.mux.Use(middleware.Timeout(60 * time.Second))

	s.mux.Get("/sitemap", s.handleGetSitemap())

	s.mux.Get("/health", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	})

	s.mux.Route("/content", func(r chi.Router) {
		r.Route("/landing-images", func(r chi.Router) {
			r.Get("/", s.handleLandingImages())
			r.Post("/", s.withPermissions(s.handleUploadLandingImage(), "edit:content"))
			r.Delete("/{imageID}", s.withPermissions(s.handleDeleteLandingImage(), "delete:content"))
		})
	})

	s.mux.Route("/members", func(r chi.Router) {
		r.Get("/", s.handleListMembers())

		r.Get("/{memberID}", s.withPermissions(s.handleMemberByID(), "view:member"))
		r.Put("/{memberID}", s.handleUpdateMember())
		r.Delete("/{memberID}", s.withPermissions(s.handleDeleteMember(), "delete:member"))

		r.Get("/{memberID}/teams", s.withPermissions(s.handleListMemberTeams(), "view:team"))
		r.Put("/{memberID}/teams", s.withPermissions(s.handleSetMemberTeams(), "view:team", "edit:member"))

		r.Get("/{memberID}/permissions", s.withPermissions(s.handleListMemberPermissions(), "view:permission", "view:member"))

		r.Post("/{memberID}/approve", s.withPermissions(s.handleApproveMember(), "edit:member"))

		r.Post("/picture", s.handleUploadMemberProfilePicture())
		// r.Get("/{memberID}", s.withPermissions(s.handleListUser(), "view:user", "view:team", "view:permission"))
	})

	s.mux.Route("/teams", func(r chi.Router) {
		r.Get("/", s.handleListTeams())
		r.Post("/", s.withPermissions(s.handleCreateTeam(), "edit:team"))

		r.Get("/{teamID}", s.withPermissions(s.handleTeamByID(), "view:team"))
		r.Delete("/{teamID}", s.withPermissions(s.handleDeleteTeam(), "delete:team"))

	})

	s.mux.Route("/auth", func(r chi.Router) {
		r.Post("/login", s.handleLogin())
		r.Post("/register", s.handleRegister())
		r.Post("/log-out", s.handleLogOut())
		r.Get("/session", s.handleGetSession())

		r.Route("/permissions", func(r chi.Router) {
			r.Get("/{teamID}", s.withPermissions(s.handleListTeamPermissions(), "view:team", "view:permission"))
			// r.Get("/", s.withPermissions(s.handleListPermissions(), "view:permission"))
			// r.Get("/{memberID}", s.withPermissions(s.handleListMemberPermissions(), "view:permission"))
		})
	})

	s.mux.Route("/events", func(r chi.Router) {
		r.Get("/", s.handleListEvents())
		r.Get("/{eventID}", s.handleEventByID())

		r.Post("/", s.withPermissions(s.handleCreateEvent(), "edit:event"))
		r.Put("/{eventID}", s.withPermissions(s.handleUpdateEvent(), "edit:event"))
		r.Delete("/{eventID}", s.withPermissions(s.handleDeleteEvent(), "delete:event"))
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
		r.Get("/{venueID}", s.withPermissions(s.handleVenueByID(), "view:venue"))

		r.Post("/", s.withPermissions(s.handleCreateVenue(), "edit:venue"))
		r.Put("/{venueID}", s.withPermissions(s.handleUpdateVenue(), "edit:venue"))
		r.Delete("/{venueID}", s.withPermissions(s.handleDeleteVenue(), "delete:venue"))
	})

	s.mux.Route("/genres", func(r chi.Router) {
		r.Post("/", s.withPermissions(s.handleCreateGenre(), "edit:genre"))
		r.Get("/", s.handleListGenres())
	})
}
