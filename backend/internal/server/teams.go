package server

import (
	"encoding/json"
	"net/http"
	"github.com/mattismoel/konnekt/internal/service"
)

func (s Server) handleListTeams() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()

		query, err := NewListQueryFromURL(r.URL.Query())
		if err != nil {
			writeError(w, err)
			return
		}

		result, err := s.teamService.List(ctx, query)
		if err != nil {
			writeError(w, err)
			return
		}

		writeJSON(w, http.StatusOK, result)
	}
}

func (s Server) handleListMemberTeams() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()

		memberID, err := paramID("memberID", r)
		if err != nil {
			writeError(w, err)
			return
		}

		teams, err := s.teamService.MemberTeams(ctx, memberID)
		if err != nil {
			writeError(w, err)
			return
		}

		writeJSON(w, http.StatusOK, teams)
	}
}

func (s Server) handleCreateTeam() http.HandlerFunc {
	type CreateTeamLoad struct {
		Name        string `json:"name"`
		DisplayName string `json:"displayName"`
		Description string `json:"description"`
	}

	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()

		var load CreateTeamLoad

		err := json.NewDecoder(r.Body).Decode(&load)
		if err != nil {
			writeError(w, err)
			return
		}

		team, err := s.teamService.Create(ctx, service.CreateTeam{
			Name:        load.Name,
