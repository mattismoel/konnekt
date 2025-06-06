package server

import "net/http"

func (s Server) handleListMemberPermissions() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()

		memberID, err := paramID("memberID", r)
		if err != nil {
			writeError(w, err)
			return
		}

		perms, err := s.authService.MemberPermissions(ctx, memberID)
		if err != nil {
			writeError(w, err)
			return
		}

		writeJSON(w, http.StatusOK, perms)
	}
}
