package auth

import (
	"context"
	"time"

	"github.com/mattismoel/konnekt/internal/query"
)

type Repository interface {
	Session(ctx context.Context, sessionID SessionID) (Session, error)
	InsertSession(ctx context.Context, s Session) error
	DeleteUserSession(ctx context.Context, userID int64) error
	SetSessionExpiry(ctx context.Context, sessionID SessionID, newExpiry time.Time) error

	InsertRole(ctx context.Context, r Role) (int64, error)
	ListRoles(ctx context.Context, query query.ListQuery) (query.ListResult[Role], error)
	DeleteRole(ctx context.Context, roleID int64) error
	RoleByID(ctx context.Context, id int64) (Role, error)
	RoleByName(ctx context.Context, name string) (Role, error)

	UserRoles(ctx context.Context, userID int64) ([]Role, error)
	AddUserRoles(ctx context.Context, userID int64, roleIDs ...int64) error

	ListPermissions(ctx context.Context, q query.ListQuery) (query.ListResult[Permission], error)
	RolePermissions(ctx context.Context, roleID int64) (PermissionCollection, error)
}
