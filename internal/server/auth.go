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
	ErrUnauthorized       = APIError{Message: "User unauthorized", Status: http.StatusUnauthorized}
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

		token, expiry, err := s.authService.Register(r.Context(),
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

		writeSessionCookie(w, token, expiry)
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

func (s Server) handleLogOut() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		sessionCookie, err := r.Cookie(SESSION_COOKIE_NAME)
		if err != nil {
			switch {
			case errors.Is(err, http.ErrNoCookie):
				break
			default:
				writeError(w, err)
				return
			}
		}

		token := sessionCookie.Value

		err = s.authService.LogOut(r.Context(), auth.SessionToken(token))
		if err != nil {
			switch {
			case errors.Is(err, auth.ErrNoSession):
				break
			default:
				writeError(w, err)
				return
			}
		}

		clearSessionCookie(w)
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

func clearSessionCookie(w http.ResponseWriter) {
	http.SetCookie(w, &http.Cookie{
		Name:     SESSION_COOKIE_NAME,
		Value:    "",
		MaxAge:   0,
		Path:     "/",
		HttpOnly: true,
		Secure:   true,
		SameSite: http.SameSiteLaxMode,
	})
}
