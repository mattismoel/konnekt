package artist

import (
	"context"

	"github.com/mattismoel/konnekt/internal/query"
)

type Repository interface {
	Insert(ctx context.Context, a Artist) (int64, error)
	Update(ctx context.Context, artistID int64, a Artist) error
	List(ctx context.Context, q query.ListQuery) (query.ListResult[Artist], error)
	ByID(ctx context.Context, artistID int64) (Artist, error)
	Delete(ctx context.Context, artistID int64) error
	GenreByID(ctx context.Context, genreID int64) (Genre, error)
	ListGenres(ctx context.Context, q GenreQuery) (query.ListResult[Genre], error)
	InsertGenre(ctx context.Context, name string) (int64, error)
	SetImageURL(ctx context.Context, artistID int64, u string) error
}
