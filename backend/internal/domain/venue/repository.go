package venue

import (
	"context"

	"github.com/mattismoel/konnekt/internal/query"
)

type Repository interface {
	Insert(ctx context.Context, v Venue) (int64, error)
	ByID(ctx context.Context, venueID int64) (Venue, error)
	List(ctx context.Context, q Query) (query.ListResult[Venue], error)
	Delete(ctx context.Context, venueID int64) error
}
