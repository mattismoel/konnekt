package konnekt

import (
	"context"

	"github.com/mattismoel/konnekt/backend/api"
)

type Team struct {
	ID          int64  `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	DisplayName string `json:"displayName"`
}

type Permission struct {
	ID          int64  `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	DisplayName string `json:"displayName"`
}

type AuthRepo interface {
	TeamPermissions(context.Context, int64) (api.ListResponse[Permission], error)
	MemberTeams(context.Context, int64) (api.ListResponse[Team], error)
}
