package auth

import (
	"context"
	"time"

	"github.com/mattismoel/konnekt/internal/query"
)

type Repository interface {
	Session(ctx context.Context, sessionID SessionID) (Session, error)
	InsertSession(ctx context.Context, s Session) error
	DeleteMemberSession(ctx context.Context, memberID int64) error
	SetSessionExpiry(ctx context.Context, sessionID SessionID, newExpiry time.Time) error

	ListPermissions(ctx context.Context, q query.ListQuery) (query.ListResult[Permission], error)
	TeamPermissions(ctx context.Context, teamID int64) (PermissionCollection, error)
}
