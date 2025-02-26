package venue

import "context"

type Repository interface {
	Insert(ctx context.Context, v Venue) (int64, error)
	ByID(ctx context.Context, venueID int64) (Venue, error)
	List(ctx context.Context, limit, offset int) ([]Venue, int, error)
	Delete(ctx context.Context, venueID int64) error
}
