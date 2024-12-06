package rest

import (
	"context"
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/mattismoel/konnekt/internal/password"
	"github.com/mattismoel/konnekt/internal/service"
)

type AuthService interface {
	GenerateSessionToken() (service.SessionToken, error)
	CreateSession(context.Context, service.SessionToken, int64) (service.Session, error)
	ValidateSessionToken(context.Context, service.SessionToken) (service.Session, service.User, error)
	InvalidateSession(context.Context, service.SessionID) error
	HasPermission(ctx context.Context, userID int64, perm string) error
	Login(ctx context.Context, w http.ResponseWriter, email string, password []byte) (service.User, error)
	Register(ctx context.Context, w http.ResponseWriter, r *http.Request, user service.User, password password.Password, passwordConfirm password.Password) (service.User, error)
}

func (s server) createAuthRoutes() http.Handler {
	r := chi.NewRouter()

	r.Post("/register", s.handleRegister())
	r.Post("/login", s.handleLogin())
	r.Post("/logout", s.handleLogout())

	return r
}

type RegisterLoad struct {
	service.User
	Password        string `json:"password"`
	PasswordConfirm string `json:"passwordConfirm"`
}

func (s server) handleRegister() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var load RegisterLoad

		err := readJSON(r, &load)
		if err != nil {
			Error(w, r, err)
			return
		}

		user, err := s.authService.Register(r.Context(), w, r, load.User, password.Password(load.Password), password.Password(load.PasswordConfirm))
		if err != nil {
			Error(w, r, err)
			return
		}

		writeJSON(w, http.StatusOK, user)
	}
}

type LoginLoad struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (s server) handleLogin() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var loginLoad LoginLoad

		err := readJSON(r, &loginLoad)
		if err != nil {
			Error(w, r, err)
			return
		}

		user, err := s.authService.Login(r.Context(), w, loginLoad.Email, []byte(loginLoad.Password))
		if err != nil {
			Error(w, r, err)
			return
		}

		fmt.Println(user)

		writeJSON(w, http.StatusOK, user)
	}
}

func (s server) handleLogout() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {}
}
