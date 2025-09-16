package server

import (
	"context"
	"errors"
	"fmt"
	"net/http"

	konnekt "github.com/mattismoel/konnekt/backend"
	"github.com/mattismoel/konnekt/backend/api"
	"github.com/mattismoel/konnekt/backend/auth"
	"golang.org/x/crypto/bcrypt"
)

const (
	SESSION_COOKIE_NAME = "konnekt-session"
)

type SessionRepo interface {
	InsertSession(context.Context, auth.Session) (int64, error)
	GetSession(context.Context, auth.SessionID) (auth.Session, error)
	DeleteSession(context.Context, auth.SessionID) error
}

type AuthRepo interface {
	TeamPermissions(context.Context, int64) (api.ListResponse[konnekt.Permission], error)
	MemberTeams(context.Context, int64) (api.ListResponse[konnekt.Team], error)
}

type Register struct {
	Email           string `json:"email"`
	FirstName       string `json:"firstName"`
	LastName        string `json:"lastName"`
	AvatarURL       string `json:"avatarUrl"`
	Password        string `json:"password"`
	PasswordConfirm string `json:"passwordConfirm"`
}

type Login struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (s Server) handleRegisterMember() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()

		var load Register
		if err := api.ReadJSON(r.Body, &load); err != nil {
			api.WriteError(w, r, err)
			return
		}

		_, err := s.MemberRepo.MemberByEmail(ctx, load.Email)
		if err == nil {
			api.WriteError(w, r, api.BadRequestError(r, "Member already exists"))
			return
		}

		if err := auth.ValidatePasswordConfirm([]byte(load.Password), []byte(load.PasswordConfirm)); err != nil {
			api.WriteError(w, r, api.BadRequestError(r, "Passwords do not match"))
			return
		}

		cm := konnekt.CreateMember{
			Email:     load.Email,
			FirstName: load.FirstName,
			LastName:  load.LastName,
			AvatarURL: load.AvatarURL,
			Password:  load.Password,
		}

		memberID, err := s.MemberRepo.InsertMember(ctx, cm)
		if err != nil {
			api.WriteError(w, r, err)
			return
		}

		member, err := s.MemberRepo.MemberByID(ctx, memberID)
		if err != nil {
			api.WriteError(w, r, err)
			return
		}

		session, err := auth.CreateSession(memberID)
		if err != nil {
			api.WriteError(w, r, err)
			return
		}

		_, err = s.SessionRepo.InsertSession(ctx, session)
		if err != nil {
			api.WriteError(w, r, err)
			return
		}

		sessionCookie := createSessionCookie(session)
		fmt.Printf("%+v\n", sessionCookie)
		http.SetCookie(w, sessionCookie)

		if err := api.WriteJSON(w, member, int(http.StatusCreated)); err != nil {
			api.WriteError(w, r, err)
			return
		}
	}
}

func (s Server) handleLoginMember() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()

		var load Login
		if err := api.ReadJSON(r.Body, &load); err != nil {
			api.WriteError(w, r, err)
			return
		}

		member, err := s.MemberRepo.MemberByEmail(ctx, load.Email)
		if err != nil {
			if errors.Is(err, konnekt.ErrResourceNotFound) {
				api.WriteError(w, r, api.UnauthorisedError(r))
				return
			}
			api.WriteError(w, r, err)
			return
		}

		if !member.Approved {
			api.WriteError(w, r, api.UnauthorisedError(r))
			return
		}

		passwordHash, err := s.MemberRepo.MemberPasswordHash(ctx, member.ID)
		if err != nil {
			api.WriteError(w, r, err)
			return
		}

		if err := bcrypt.CompareHashAndPassword(passwordHash, []byte(load.Password)); err != nil {
			api.WriteError(w, r, err)
			return
		}

		sessCookie, err := sessionCookie(r)
		if err != nil && !errors.Is(err, http.ErrNoCookie) {
			api.WriteError(w, r, err)
			return
		}

		// If session token already exists, delete previous session.
		if sessCookie != nil {
			token := auth.SessionToken(sessCookie.Value)
			sessionID, err := token.SessionID()
			if err != nil {
				api.WriteError(w, r, err)
				return
			}

			if err := s.SessionRepo.DeleteSession(ctx, sessionID); err != nil {
				api.WriteError(w, r, err)
				return
			}
		}

		session, err := auth.CreateSession(member.ID)
		if err != nil {

		}

		_, err = s.SessionRepo.InsertSession(ctx, session)
		if err != nil {
			api.WriteError(w, r, err)
			return
		}

		sessionCookie := createSessionCookie(session)
		http.SetCookie(w, sessionCookie)
	}
}

func sessionCookie(r *http.Request) (*http.Cookie, error) {
	cookie, err := r.Cookie(SESSION_COOKIE_NAME)
	if err != nil {
		if errors.Is(err, http.ErrNoCookie) {
			return nil, err
		}

		return nil, fmt.Errorf("Could not get session cookie: %v", err)
	}

	return cookie, nil
}

func createSessionCookie(s auth.Session) *http.Cookie {
	return &http.Cookie{
		Name:     SESSION_COOKIE_NAME,
		Value:    string(s.Token),
		MaxAge:   int(auth.SESSION_LIFETIME.Seconds()),
		HttpOnly: true,
		Secure:   true,
		Path:     "/",
		SameSite: http.SameSiteLaxMode,
	}
}
