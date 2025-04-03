package auth

import (
	"context"
	"time"
)

type Repository interface {
	Session(ctx context.Context, sessionID SessionID) (Session, error)
	InsertSession(ctx context.Context, s Session) error
	DeleteUserSession(ctx context.Context, userID int64) error
	SetSessionExpiry(ctx context.Context, sessionID SessionID, newExpiry time.Time) error

	AddUserRoles(ctx context.Context, userID int64, roleIDs ...int64) error
	RoleByName(ctx context.Context, name string) (Role, error)
	UserRoles(ctx context.Context, userID int64) ([]Role, error)
	RolePermissions(ctx context.Context, roleID int64) (PermissionCollection, error)
}
