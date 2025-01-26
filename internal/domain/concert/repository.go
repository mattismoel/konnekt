package concert

import "context"

type Repository interface {
	Insert(ctx context.Context, c Concert) (int64, error)
}
