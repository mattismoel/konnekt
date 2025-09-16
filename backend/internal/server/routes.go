package server

import "net/http"

func (s Server) setupRouter(mux *http.ServeMux) {
	// mux.HandleFunc("GET /events/", handleUnimplemented())
	// mux.HandleFunc("GET /events/{eventID}/", handleUnimplemented())
	// mux.HandleFunc("PATCH /events/{eventID}/", handleUnimplemented())
	// mux.HandleFunc("DELETE /events/{eventID}/", handleUnimplemented())
	mux.HandleFunc("POST /events", s.withPermissions(s.handleCreateEvent(), "edit:event"))

	mux.HandleFunc("POST /venues", s.withPermissions(s.handleCreateVenue(), "edit:venue"))

	mux.HandleFunc("GET /artists", s.handleListArtists())
	mux.HandleFunc("POST /artists", s.withPermissions(s.handleCreateArtist(), "edit:artist"))
	mux.HandleFunc("GET /artists/{artistID}", s.handleGetArtistByID())
	mux.HandleFunc("PATCH /artists/{artistID}", s.withPermissions(s.handleUpdateArtist(), "edit:artist"))
	mux.HandleFunc("DELETE /artists/{artistID}", s.withPermissions(s.handleDeleteArtist(), "delete:artist"))
	mux.HandleFunc("POST /artists/genres", s.withPermissions(s.handleCreateGenre(), "edit:genre"))

	mux.HandleFunc("POST /auth/register", s.handleRegisterMember())
	mux.HandleFunc("POST /auth/login", s.handleLoginMember())
	mux.HandleFunc("GET /members/{memberID}/teams", s.handleGetMemberTeams())
	mux.HandleFunc("GET /teams/{teamID}/permissions", s.withPermissions(s.handleGetTeamPermissions(), "view:permission"))
}
