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
}
