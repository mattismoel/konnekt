package auth

import (
	"context"
	"time"
)

type Repository interface {
	Session(ctx context.Context, sessionID SessionID) (Session, error)
	InsertSession(ctx context.Context, s Session) error
	DeleteUserSession(ctx context.Context, userID int64) error
	SetSessionExpiry(ctx context.Context, sessionID SessionID, newExpiry time.Time) (Session, error)

	// Roles(ctx context.Context) ([]Role, error)
	UserRoles(ctx context.Context, userID int64) ([]Role, error)
	RolePermissions(ctx context.Context, roleID int64) (PermissionCollection, error)
}
