package server

import (
	"context"
	"fmt"
	"io"
	"net"
	"net/http"
	"strconv"

	"github.com/mattismoel/konnekt/backend/auth"
)

type ObjectStore interface {
	Insert(context.Context, string, io.Reader) (string, error)
	Delete(context.Context, string) error
}

type Server struct {
	ArtistRepo ArtistRepo
	VenueRepo  VenueRepo
	MemberRepo MemberRepo
	EventRepo  EventRepo
	AuthRepo   AuthRepo

	SessionRepo SessionRepo
	ObjectStore ObjectStore
}

func (s Server) Start(host string, port int) error {
	mux := http.NewServeMux()
	s.setupRouter(mux)

	httpServer := &http.Server{
		Addr:    net.JoinHostPort(host, strconv.Itoa(port)),
		Handler: Chain(mux, Logger),
	}

	return httpServer.ListenAndServe()
}

func (s Server) requestMemberID(r *http.Request) (int64, error) {
	ctx := r.Context()

	sc, err := sessionCookie(r)
	if err != nil {
		return 0, err
	}

	token := auth.SessionToken(sc.Value)
	sessionID, err := token.SessionID()

	session, err := s.SessionRepo.GetSession(ctx, sessionID)
	if err != nil {
		return 0, fmt.Errorf("Could not get session: %v", err)
	}

	if err := token.Validate(session.SecretHash); err != nil {
		return 0, err
	}

	return session.MemberID, nil
}
