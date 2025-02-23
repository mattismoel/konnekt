package event

import (
	"context"
	"time"
)

type Repository interface {
	Insert(ctx context.Context, e Event) (int64, error)
	List(ctx context.Context, from, to time.Time, offset, limit int) ([]Event, int, error)
	ByID(ctx context.Context, eventID int64) (Event, error)
}
