package event

import (
	"context"
)

type Repository interface {
	Insert(ctx context.Context, e Event) (int64, error)
	Update(ctx context.Context, eventID int64, e Event) error
	List(ctx context.Context, q Query) ([]Event, int, error)
	ByID(ctx context.Context, eventID int64) (Event, error)
	SetImageURL(ctx context.Context, eventID int64, imageURL string) error
}
