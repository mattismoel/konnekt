package server

import (
	"errors"
	"net/http"

	"github.com/mattismoel/konnekt/internal/domain/auth"
)

func (s Server) withPermissions(next http.HandlerFunc, perms ...string) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()

		sessionCookie, err := r.Cookie(SESSION_COOKIE_NAME)
		if err != nil {
			writeError(w, ErrUnauthorized)
			return
		}

		token := auth.SessionToken(sessionCookie.Value)

		session, err := s.authService.Session(ctx, token.SessionID())
		if err != nil {
			writeError(w, err)
			return
		}

		err = s.authService.HasPermission(ctx, session.MemberID, perms...)
		if err != nil {
			switch {
			case errors.Is(err, auth.ErrMissingPermissions):
				writeError(w, ErrUnauthorized)
				return
			default:
				writeError(w, err)
				return
			}
		}

		next(w, r)
	})
}
