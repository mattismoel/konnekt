package server

import (
	"errors"
	"net/http"

	"github.com/mattismoel/konnekt/internal/domain/auth"
)

func (s Server) withPermissions(next http.HandlerFunc, perms ...string) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()

		session, err := s.memberSession(ctx, w, r)
		if err != nil {
			writeError(w, ErrUnauthorized)
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
