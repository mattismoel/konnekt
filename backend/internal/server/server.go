package server

import (
	"net"
	"net/http"
	"strconv"

	konnekt "github.com/mattismoel/konnekt/backend"
	"github.com/mattismoel/konnekt/backend/auth"
)

type Server struct {
	ArtistRepo  konnekt.ArtistRepo
	VenueRepo   konnekt.VenueRepo
	MemberRepo  konnekt.MemberRepo
	EventRepo   konnekt.EventRepo
	SessionRepo auth.SessionRepo
}

func (s Server) Start(host string, port int) error {
	mux := http.NewServeMux()
	s.setupRouter(mux)

	httpServer := &http.Server{
		Addr:    net.JoinHostPort(host, strconv.Itoa(port)),
		Handler: mux,
	}

	return httpServer.ListenAndServe()
}
