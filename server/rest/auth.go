package rest

import (
	"net/http"

	"github.com/go-chi/chi/v5"
)

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
