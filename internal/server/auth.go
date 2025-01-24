package server

import (
	"encoding/json"
	"errors"
	"net/http"
	"time"

	"github.com/mattismoel/konnekt/internal/domain/auth"
	"github.com/mattismoel/konnekt/internal/domain/user"
)

const (
	SESSION_COOKIE_NAME = "konnekt-session"
)

var (
	ErrPasswordsNoMatch   = APIError{Message: "Passwords do not match", Status: http.StatusBadRequest}
	ErrUserAlreadyExists  = APIError{Message: "User already exists", Status: http.StatusConflict}
	ErrInvalidCredentials = APIError{Message: "User credentials are invalid", Status: http.StatusBadRequest}
)

func (s Server) handleRegister() http.HandlerFunc {
	type RegisterLoad struct {
		Email           string `json:"email"`
		FirstName       string `json:"firstName"`
		LastName        string `json:"lastName"`
		Password        string `json:"password"`
		PasswordConfirm string `json:"passwordConfirm"`
	}

	return func(w http.ResponseWriter, r *http.Request) {
		var load RegisterLoad

		if err := json.NewDecoder(r.Body).Decode(&load); err != nil {
			writeError(w, err)
			return
		}

		session, token, err := s.authService.Register(r.Context(),
			load.Email,
			[]byte(load.Password),
			[]byte(load.PasswordConfirm),
			load.FirstName,
			load.LastName,
		)

		if err != nil {
			switch {
			case errors.Is(err, user.ErrAlreadyExists):
				writeError(w, ErrUserAlreadyExists)
			case errors.Is(err, auth.ErrPasswordsNoMatch):
				writeError(w, ErrPasswordsNoMatch)
			default:
				writeError(w, err)
			}
			return
		}

		writeSessionCookie(w, token, session.ExpiresAt)
	}
}

func (s Server) handleLogin() http.HandlerFunc {
	type LoginLoad struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	return func(w http.ResponseWriter, r *http.Request) {
		var load LoginLoad

		err := json.NewDecoder(r.Body).Decode(&load)
		if err != nil {
			writeError(w, err)
			return
		}

		ctx := r.Context()

		token, expiry, err := s.authService.Login(ctx, load.Email, []byte(load.Password))
		if err != nil {
			switch {
			case errors.Is(err, user.ErrNotFound):
				writeError(w, ErrInvalidCredentials)
			case errors.Is(err, auth.ErrPasswordsNoMatch):
				writeError(w, ErrInvalidCredentials)
			default:
				writeError(w, err)
			}
			return
		}

		writeSessionCookie(w, token, expiry)
	}
}

func writeSessionCookie(w http.ResponseWriter, token auth.SessionToken, expiresAt time.Time) {
	http.SetCookie(w, &http.Cookie{
		Name:     SESSION_COOKIE_NAME,
		Value:    string(token),
		HttpOnly: true,
		Path:     "/",
		Secure:   true,
		SameSite: http.SameSiteLaxMode,
		Expires:  expiresAt,
	})
}
