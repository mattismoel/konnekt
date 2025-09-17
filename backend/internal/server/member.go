package server

import (
	"context"
	"net/http"

	konnekt "github.com/mattismoel/konnekt/backend"
	"github.com/mattismoel/konnekt/backend/api"
)

type MemberRepo interface {
	ListMembers(context.Context, api.ListRequest) (api.ListResponse[konnekt.Member], error)
	MemberByID(context.Context, int64) (konnekt.Member, error)
	MemberByEmail(context.Context, string) (konnekt.Member, error)
	InsertMember(context.Context, konnekt.CreateMember) (int64, error)
	MemberPasswordHash(context.Context, int64) ([]byte, error)
}

func (s Server) handleListMembers() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()

		lr := api.NewListRequest(r)

		result, err := s.MemberRepo.ListMembers(ctx, lr)
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
