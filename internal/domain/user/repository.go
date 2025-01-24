package user

import (
	"context"
)

type Repository interface {
	ByEmail(ctx context.Context, email string) (User, error)
	PasswordHash(ctx context.Context, userID int64) (PasswordHash, error)
	Insert(ctx context.Context, email, firstName, lastName string, passwordHash []byte) (int64, error)
}
