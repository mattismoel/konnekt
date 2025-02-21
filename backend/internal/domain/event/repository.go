package event

import (
	"context"
)

type Repository interface {
	Insert(ctx context.Context, e Event) (int64, error)
	List(ctx context.Context, query Query) ([]Event, int, error)
	ByID(ctx context.Context, eventID int64) (Event, error)
}
