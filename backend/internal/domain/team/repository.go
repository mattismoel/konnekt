package team

import (
	"context"

	"github.com/mattismoel/konnekt/internal/query"
)

type Repository interface {
	Insert(ctx context.Context, r Team) (int64, error)
	List(ctx context.Context, query query.ListQuery) (query.ListResult[Team], error)
	Delete(ctx context.Context, teamID int64) error
	ByID(ctx context.Context, id int64) (Team, error)
	ByName(ctx context.Context, name string) (Team, error)

	MemberTeams(ctx context.Context, memberID int64) (TeamCollection, error)
	AddMemberTeams(ctx context.Context, memberID int64, teamIDs ...int64) error
}
