package konnekt

import (
	"net"
	"net/http"
	"strconv"
)

type Server struct {
	ArtistRepo ArtistRepo
}

func (s Server) Start(host string, port int) error {
	mux := http.NewServeMux()
	s.setupRouter(mux)

	httpServer := http.Server{
		Handler: mux,
		Addr:    net.JoinHostPort(host, strconv.Itoa(port)),
	}

	return httpServer.ListenAndServe()
}
