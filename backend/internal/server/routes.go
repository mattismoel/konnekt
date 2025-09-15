package server

import "net/http"

func (s Server) setupRouter(mux *http.ServeMux) {
	// mux.HandleFunc("GET /events/", handleUnimplemented())
	mux.HandleFunc("POST /events", s.handleCreateEvent())
	// mux.HandleFunc("GET /events/{eventID}/", handleUnimplemented())
	// mux.HandleFunc("PATCH /events/{eventID}/", handleUnimplemented())
	// mux.HandleFunc("DELETE /events/{eventID}/", handleUnimplemented())

	mux.HandleFunc("POST /venues", s.handleCreateVenue())

	mux.HandleFunc("GET /artists", s.handleListArtists())
	mux.HandleFunc("POST /artists", s.handleCreateArtist())
	mux.HandleFunc("GET /artists/{artistID}", s.handleGetArtistByID())
	mux.HandleFunc("PATCH /artists/{artistID}", s.handleUpdateArtist())
	mux.HandleFunc("POST /artists/genres", s.handleCreateGenre())
	mux.HandleFunc("DELETE /artists/{artistID}", s.handleDeleteArtist())

	mux.HandleFunc("POST /auth/register", s.handleRegisterMember())
	mux.HandleFunc("POST /auth/login", s.handleLoginMember())
}
