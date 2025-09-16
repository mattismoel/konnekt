package server

import "net/http"

func (s Server) setupRouter(mux *http.ServeMux) {
	// |=> EVENT ROUTES.
	mux.HandleFunc("GET /events", s.handleListEvents())
	mux.HandleFunc("GET /events/{eventID}", s.handleGetEventByID())
	mux.HandleFunc("PATCH /events/{eventID}", s.withPermissions(s.handleUpdateEvent(), "edit:event"))
	mux.HandleFunc("POST /events", s.withPermissions(s.handleCreateEvent(), "edit:event"))

	// |=> VENUE ROUTES.
	mux.HandleFunc("POST /venues", s.withPermissions(s.handleCreateVenue(), "edit:venue"))
	mux.HandleFunc("GET /venues", s.handleListVenues())

	// |=> ARTIST ROUTES.
	mux.HandleFunc("GET /artists", s.handleListArtists())
	mux.HandleFunc("POST /artists", s.withPermissions(s.handleCreateArtist(), "edit:artist"))
	mux.HandleFunc("GET /artists/{artistID}", s.handleGetArtistByID())
	mux.HandleFunc("PATCH /artists/{artistID}", s.withPermissions(s.handleUpdateArtist(), "edit:artist"))
	mux.HandleFunc("DELETE /artists/{artistID}", s.withPermissions(s.handleDeleteArtist(), "delete:artist"))
	mux.HandleFunc("POST /artists/genres", s.withPermissions(s.handleCreateGenre(), "edit:genre"))

	// |=> AUTHENTICATION ROUTES.
	mux.HandleFunc("POST /auth/register", s.handleRegisterMember())
	mux.HandleFunc("POST /auth/login", s.handleLoginMember())

	// |=> AUTHORIZATION ROUTES.
	mux.HandleFunc("GET /members/{memberID}/teams", s.handleGetMemberTeams())
	mux.HandleFunc("GET /teams/{teamID}/permissions", s.withPermissions(s.handleGetTeamPermissions(), "view:permission"))
}
