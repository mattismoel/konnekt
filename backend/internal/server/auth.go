package server

import (
	"encoding/json"
	"errors"
	"net/http"
	"strconv"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/mattismoel/konnekt/internal/domain/auth"
	"github.com/mattismoel/konnekt/internal/domain/member"
)

const (
	SESSION_COOKIE_NAME = "konnekt-session"
)

var (
	ErrPasswordsNoMatch         = APIError{Message: "Passwords do not match", Status: http.StatusBadRequest}
	ErrMemberAlreadyExists      = APIError{Message: "Member already exists", Status: http.StatusConflict}
	ErrMemberInvalidCredentials = APIError{Message: "Member credentials are invalid", Status: http.StatusBadRequest}
	ErrUnauthorized             = APIError{Message: "Member unauthorized", Status: http.StatusUnauthorized}
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

		err := s.authService.Register(r.Context(),
			load.Email,
			[]byte(load.Password),
			[]byte(load.PasswordConfirm),
			load.FirstName,
			load.LastName,
		)

		if err != nil {
			switch {
			case errors.Is(err, member.ErrAlreadyExists):
				writeError(w, ErrMemberAlreadyExists)
			case errors.Is(err, auth.ErrPasswordsNoMatch):
				writeError(w, ErrPasswordsNoMatch)
			default:
				writeError(w, err)
			}
			return
		}
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
			case errors.Is(err, member.ErrNotFound):
				writeError(w, ErrMemberInvalidCredentials)
			case errors.Is(err, auth.ErrPasswordsNoMatch):
				writeError(w, ErrMemberInvalidCredentials)
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

func (s Server) handleGetSession() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()

		sessionCookie, err := r.Cookie(SESSION_COOKIE_NAME)
		if err != nil {
			switch {
			case errors.Is(err, http.ErrNoCookie):
				writeError(w, ErrUnauthorized)
			default:
				writeError(w, err)
			}
			return
		}

		token := auth.SessionToken(sessionCookie.Value)

		newExpiry, err := s.authService.ValidateSession(ctx, token)
		if err != nil {
			writeError(w, err)
			return
		}

		session, err := s.authService.Session(ctx, token.SessionID())
		if err != nil {
			writeError(w, err)
			return
		}

		member, err := s.memberService.ByID(ctx, session.MemberID)
		if err != nil {
			writeError(w, err)
			return
		}

		writeSessionCookie(w, token, newExpiry)

		writeJSON(w, http.StatusOK, member)
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

func (s Server) handleListTeamPermissions() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		teamID, err := strconv.Atoi(chi.URLParam(r, "teamID"))
		if err != nil {
			writeError(w, err)
			return
		}

		ctx := r.Context()

		perms, err := s.teamService.TeamPermissions(ctx, int64(teamID))
		if err != nil {
			writeError(w, err)
			return
		}

		if err := writeJSON(w, http.StatusOK, perms); err != nil {
			writeError(w, err)
			return
		}
	}
}
