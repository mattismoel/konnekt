package rest

import (
	"context"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/mattismoel/konnekt"
	"github.com/mattismoel/konnekt/internal/service"
)

type SessionService interface {
	GenerateSessionToken() (konnekt.SessionToken, error)
	CreateSession(context.Context, konnekt.SessionToken, int64) (konnekt.Session, error)
	ValidateSessionToken(context.Context, konnekt.SessionToken) (konnekt.Session, service.User, error)
	InvalidateSession(context.Context, konnekt.SessionID) error
}

func (s server) createAuthRoutes() http.Handler {
	r := chi.NewRouter()

	r.Post("/register", s.handleRegister())
	r.Post("/login", s.handleLogin())
	r.Post("/logout", s.handleLogout())

	return r
}

func (s server) handleRegister() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {}
}

func (s server) handleLogin() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {}
}

func (s server) handleLogout() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {}
}
