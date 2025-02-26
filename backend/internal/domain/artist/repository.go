package artist

import (
	"context"
)

type Repository interface {
	Insert(ctx context.Context, a Artist) (int64, error)
	List(ctx context.Context, offset, limit int) ([]Artist, int, error)
	ByID(ctx context.Context, artistID int64) (Artist, error)
	Delete(ctx context.Context, artistID int64) error
	GenreByID(ctx context.Context, genreID int64) (Genre, error)
	ListGenres(ctx context.Context, offset, limit int) ([]Genre, int, error)
}
