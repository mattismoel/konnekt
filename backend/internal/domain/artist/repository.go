package artist

import (
	"context"
)

type Repository interface {
	Insert(ctx context.Context, a Artist) (int64, error)
	List(ctx context.Context) ([]Artist, int, error)
	ByID(ctx context.Context, artistID int64) (Artist, error)
	GenreByID(ctx context.Context, genreID int64) (Genre, error)
}
