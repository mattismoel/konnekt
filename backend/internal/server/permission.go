package server

import (
	"errors"
	"net/http"
	"slices"

	konnekt "github.com/mattismoel/konnekt/backend"
	"github.com/mattismoel/konnekt/backend/api"
	"github.com/mattismoel/konnekt/backend/auth"
	"github.com/mattismoel/konnekt/backend/urlutil"
)

func (s Server) handleGetTeamPermissions() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()

		teamID, err := urlutil.PathInt(r, "teamID")
		if err != nil {
			api.WriteError(w, r, api.BadRequestError(r, "Invalid ID"))
			return
		}

		result, err := s.AuthRepo.TeamPermissions(ctx, int64(teamID))
		if err != nil {
			api.WriteError(w, r, err)
			return
		}

		if err := api.WriteJSON(w, result, int(http.StatusOK)); err != nil {
			api.WriteError(w, r, err)
			return
		}
	}
}

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

		session, err := s.SessionRepo.GetSession(ctx, sessionID)
		if err != nil {
			api.WriteError(w, r, api.UnauthorisedError(r))
			return
		}

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
