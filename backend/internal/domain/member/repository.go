package member

import (
	"context"

	"github.com/mattismoel/konnekt/internal/query"
)

type Repository interface {
	ByEmail(ctx context.Context, email string) (Member, error)
	ByID(ctx context.Context, memberID int64) (Member, error)
	PasswordHash(ctx context.Context, memberID int64) (PasswordHash, error)
	Insert(ctx context.Context, m Member) (int64, error)
	SetProfilePictureURL(ctx context.Context, memberID int64, url string) error
}
