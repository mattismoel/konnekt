package server

import (
	"errors"
	"fmt"
	"net/http"
	"slices"

	konnekt "github.com/mattismoel/konnekt/backend"
	"github.com/mattismoel/konnekt/backend/api"
	"github.com/mattismoel/konnekt/backend/auth"
)

func (s Server) withPermissions(h http.Handler, reqPerms ...string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()

		sessCookie, err := sessionCookie(r)
		if err != nil {
			if errors.Is(err, http.ErrNoCookie) {
				api.WriteError(w, r, api.UnauthorisedError(r))
				return
			}

			api.WriteError(w, r, err)
			return
		}

		sessionToken := auth.SessionToken(sessCookie.Value)
		sessionID, err := sessionToken.SessionID()
		if err != nil {
			api.WriteError(w, r, api.UnauthorisedError(r))
			return
		}

		fmt.Println("SESSION ID", sessionID)

		session, err := s.SessionRepo.GetSession(ctx, sessionID)
		if err != nil {
			api.WriteError(w, r, api.UnauthorisedError(r))
			return
		}
		fmt.Printf("SESSION: %+v\n", session)

		if err := sessionToken.Validate(session.SecretHash); err != nil {
			api.WriteError(w, r, api.UnauthorisedError(r))
			return
		}

		mtResult, err := s.AuthRepo.MemberTeams(ctx, session.MemberID)
		if err != nil {
			api.WriteError(w, r, api.UnauthorisedError(r))
			return
		}

		memberPerms := make([]konnekt.Permission, 0)
		for _, mt := range mtResult.Records {
			tpResult, err := s.AuthRepo.TeamPermissions(ctx, mt.ID)
			if err != nil {
				api.WriteError(w, r, err)
				return
			}

			memberPerms = append(memberPerms, tpResult.Records...)
		}

		if !hasAllPermissions(memberPerms, reqPerms...) {
			api.WriteError(w, r, api.UnauthorisedError(r))
			return
		}

		h.ServeHTTP(w, r)
	}
}

func hasAllPermissions(perms []konnekt.Permission, reqPerms ...string) bool {
	memberPermNames := make([]string, 0)
	for _, p := range perms {
		memberPermNames = append(memberPermNames, p.Name)
	}

	for _, rpName := range reqPerms {
		if !slices.Contains(memberPermNames, rpName) {
			return false
		}
	}

	return true
}
