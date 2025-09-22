package server

import "net/http"

func (s Server) setupRouter(mux *http.ServeMux) {
	// |=> EVENT ROUTES.
	mux.HandleFunc("GET /events", s.handleListEvents())
	mux.HandleFunc("GET /events/{eventID}", s.handleGetEventByID())
	mux.HandleFunc("POST /events", s.withPermissions(s.handleCreateEvent(), "edit:event"))
	mux.HandleFunc("PATCH /events/{eventID}", s.withPermissions(s.handleUpdateEvent(), "edit:event"))
	mux.HandleFunc("DELETE /events/{eventID}", s.withPermissions(s.handleDeleteEvent(), "delete:event"))

	// |=> VENUE ROUTES.
	mux.HandleFunc("GET /venues", s.handleListVenues())
	mux.HandleFunc("GET /venues/{venueID}", s.handleGetVenueByID())
	mux.HandleFunc("POST /venues", s.withPermissions(s.handleCreateVenue(), "edit:venue"))
	mux.HandleFunc("PATCH /venues/{venueID}", s.withPermissions(s.handleUpdateVenue(), "edit:venue"))
	mux.HandleFunc("DELETE /venues/{venueID}", s.withPermissions(s.handleDeleteVenue(), "delete:venue"))

	// |=> ARTIST ROUTES.
	mux.HandleFunc("GET /artists", s.handleListArtists())
	mux.HandleFunc("GET /artists/{artistID}", s.handleGetArtistByID())
	mux.HandleFunc("GET /artists/{artistID}/events", s.handleGetArtistEvents())
	mux.HandleFunc("POST /artists", s.withPermissions(s.handleCreateArtist(), "edit:artist"))
	mux.HandleFunc("POST /artists/genres", s.withPermissions(s.handleCreateGenre(), "edit:genre"))
	mux.HandleFunc("PATCH /artists/{artistID}", s.withPermissions(s.handleUpdateArtist(), "edit:artist"))
	mux.HandleFunc("DELETE /artists/{artistID}", s.withPermissions(s.handleDeleteArtist(), "delete:artist"))

	// |=> AUTHENTICATION ROUTES.
	mux.HandleFunc("POST /auth/register", s.handleRegisterMember())
	mux.HandleFunc("POST /auth/login", s.handleLoginMember())

	// |=> AUTHORIZATION ROUTES.
	mux.HandleFunc("GET /members", s.handleListMembers())
	mux.HandleFunc("GET /members/{memberID}/teams", s.handleGetMemberTeams())
	mux.HandleFunc("POST /members/avatars", s.handleUploadAvatar())
	mux.HandleFunc("GET /teams/{teamID}/permissions", s.withPermissions(s.handleGetTeamPermissions(), "view:permission"))
	mux.HandleFunc("GET /session", s.handleGetSession())
}
