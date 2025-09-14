package konnekt

import (
	"net/http"
)

func (s Server) setupRouter(mux *http.ServeMux) {
	mux.HandleFunc("GET /artists", s.handleListArtists())
	mux.HandleFunc("POST /artists", s.handleCreateArtist())
	mux.HandleFunc("GET /artists/{artistID}", s.handleGetArtistByID())
	mux.HandleFunc("PATCH /artists/{artistID}", s.handleUpdateArtist())
	mux.HandleFunc("POST /artists/genres", s.handleCreateGenre())
	mux.HandleFunc("DELETE /artists/{artistID}", s.handleDeleteArtist())
}
