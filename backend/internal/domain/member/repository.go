package member

import (
	"context"

	"github.com/mattismoel/konnekt/internal/query"
)

type Repository interface {
	ByEmail(ctx context.Context, email string) (Member, error)
	ByID(ctx context.Context, memberID int64) (Member, error)
	List(ctx context.Context, query query.ListQuery) (query.ListResult[Member], error)
	PasswordHash(ctx context.Context, memberID int64) (PasswordHash, error)
	Insert(ctx context.Context, m Member) (int64, error)
	Update(ctx context.Context, memberID int64, m Member) error
	SetMemberTeams(ctx context.Context, memberID int64, teamIDs ...int64) error
	Approve(ctx context.Context, memberID int64) error
	Delete(ctx context.Context, memberID int64) error
	SetProfilePictureURL(ctx context.Context, memberID int64, url string) error
}
