package server

import (
	"encoding/json"
	"net/http"
	"strings"

	"github.com/mattismoel/konnekt/internal/domain/member"
)

func (s Server) handleListMembers() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		query, err := NewListQueryFromURL(r.URL.Query())
		if err != nil {
			writeError(w, err)
			return
		}

		ctx := r.Context()

		result, err := s.memberService.List(ctx, query)
		if err != nil {
			writeError(w, err)
			return
		}

		err = writeJSON(w, http.StatusOK, result)
		if err != nil {
			writeError(w, err)
			return
		}
	}
}

func (s Server) handleApproveMember() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		memberID, err := paramID("memberID", r)
		if err != nil {
			writeError(w, err)
			return
		}

		ctx := r.Context()

		err = s.memberService.Approve(ctx, memberID)
		if err != nil {
			writeError(w, err)
			return
		}

		w.WriteHeader(http.StatusOK)
	}
}

func (s Server) handleDeleteMember() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		memberID, err := paramID("memberID", r)
		if err != nil {
			writeError(w, err)
			return
		}

		ctx := r.Context()

		err = s.memberService.Delete(ctx, memberID)
		if err != nil {
			writeError(w, err)
			return
		}

		w.WriteHeader(http.StatusOK)
	}
}

func (s Server) handleMemberByID() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		memberID, err := paramID("memberID", r)
		if err != nil {
			writeError(w, err)
			return
		}

		ctx := r.Context()

		m, err := s.memberService.ByID(ctx, memberID)
		if err != nil {
			writeError(w, err)
			return
		}

		if err := writeJSON(w, http.StatusOK, m); err != nil {
			writeError(w, err)
			return
		}
	}
}

func (s Server) handleUpdateMember() http.HandlerFunc {
	type UpdateMemberLoad struct {
		Email             string  `json:"email"`
		FirstName         string  `json:"firstName"`
		LastName          string  `json:"lastName"`
		ProfilePictureURL string  `json:"profilePictureUrl"`
		MemberTeamIDs     []int64 `json:"memberTeams"`
	}

	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()

		session, err := s.memberSession(ctx, w, r)
		if err != nil {
			writeError(w, ErrUnauthorized)
			return
		}

		memberID, err := paramID("memberID", r)
		if err != nil {
			writeError(w, err)
			return
		}

		reqMember, err := s.memberService.ByID(ctx, session.MemberID)
		if err != nil {
			writeError(w, err)
			return
		}

		isAdmin := false
		for _, team := range reqMember.Teams {
			if team.Name == "admin" {
				isAdmin = true
			}
		}

		if (session.MemberID != memberID) && !isAdmin {
			writeError(w, ErrUnauthorized)
			return
		}

		var load UpdateMemberLoad

		err = json.NewDecoder(r.Body).Decode(&load)
		if err != nil {
			writeError(w, err)
			return
		}

		m, err := member.NewMember(
			member.WithEmail(load.Email),
			member.WithFirstName(load.FirstName),
			member.WithLastName(load.LastName),
		)

		if err != nil {
			writeError(w, err)
			return
		}

		if strings.TrimSpace(load.ProfilePictureURL) != "" {
			err := m.WithCfgs(member.WithProfilePictureURL(load.ProfilePictureURL))
			if err != nil {
				writeError(w, err)
				return
			}
		}

		if err := s.memberService.Update(ctx, memberID, m); err != nil {
			writeError(w, err)
			return
		}

		err = s.memberService.SetMemberTeams(ctx, memberID, load.MemberTeamIDs...)
		if err != nil {
			writeError(w, err)
			return
		}

		updatedMember, err := s.memberService.ByID(ctx, memberID)
		if err != nil {
			writeError(w, err)
			return
		}

		if err := writeJSON(w, http.StatusOK, updatedMember); err != nil {
			writeError(w, err)
		}
	}
}

func (srv Server) handleSetMemberTeams() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()

		memberID, err := paramID("memberID", r)
		if err != nil {
			writeError(w, err)
			return
		}

		var teamIDs []int64

		if err := json.NewDecoder(r.Body).Decode(&teamIDs); err != nil {
			writeError(w, err)
			return
		}

		err = srv.memberService.SetMemberTeams(ctx, memberID, teamIDs...)
		if err != nil {
			writeError(w, err)
			return
		}

		w.WriteHeader(http.StatusOK)
	}
}

func (srv Server) handleUploadMemberProfilePicture() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		file, _, err := r.FormFile("file")
		if err != nil {
			writeError(w, err)
			return
		}

		defer file.Close()

		ctx := r.Context()

		profilePictureUrl, err := srv.memberService.UploadProfilePicture(ctx, file)

		if err != nil {
			writeError(w, err)
			return
		}

		if err := writeText(w, http.StatusOK, profilePictureUrl); err != nil {
			writeError(w, err)
			return
		}

		w.WriteHeader(http.StatusOK)
	}
}
