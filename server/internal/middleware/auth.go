package middleware

import (
	"context"
	"errors"
	"net/http"

	"github.com/mattismoel/konnekt/internal/service"
)

type AuthService interface {
	ValidateSessionToken(ctx context.Context, token service.SessionToken) (service.Session, service.User, error)
	SetSessionTokenCookie(ctx context.Context, w http.ResponseWriter, token service.SessionToken)
	DeleteSessionTokenCookie(ctx context.Context, w http.ResponseWriter)
}

func WithAuth(next http.Handler, authService AuthService) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		sessionCookie, err := r.Cookie(service.SESSION_COOKIE_NAME)
		if err != nil {
			if errors.Is(err, http.ErrNoCookie) {
				service.Errorf(service.ERRUNAUTHORIZED, "No session")
				return
			}
		}

		token := []byte(sessionCookie.Value)
		_, _, err = authService.ValidateSessionToken(r.Context(), token)
		if err != nil {
			if errors.Is(err, service.ErrSessionExpired) || errors.Is(err, service.ErrNoSession) {
				authService.DeleteSessionTokenCookie(r.Context(), w)
				service.Errorf(service.ERRUNAUTHORIZED, err.Error())
				return
			}

			service.Errorf(service.ERRINTERNAL, err.Error())
			return
		}

		authService.SetSessionTokenCookie(r.Context(), w, token)

		next.ServeHTTP(w, r)
	})
}
