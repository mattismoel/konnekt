package event

import (
	"context"
)

type Repository interface {
	Insert(ctx context.Context, e Event) (int64, error)
}
