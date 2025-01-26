package artist

import "context"

type Repository interface {
	Insert(ctx context.Context, a Artist) (int64, error)
	ByID(ctx context.Context, artistID int64) (Artist, error)
}
