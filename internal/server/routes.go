package server

import (
	"github.com/go-chi/chi/v5"
)

func (s *Server) setupRoutes() {
	s.mux.Route("/auth", func(r chi.Router) {
		r.Post("/login", s.handleLogin())
		r.Post("/register", s.handleRegister())
		r.Post("/log-out", s.handleLogOut())
	})
}
