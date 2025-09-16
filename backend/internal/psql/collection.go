package psql

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5"
)

type Collection[T Assembler[T, K], K any] []T

func (c Collection[T, K]) Assemble(ctx context.Context, tx pgx.Tx) ([]K, error) {
	items := make([]K, 0)
	for _, dbItem := range c {
		item, err := dbItem.Assemble(ctx, tx)
		if err != nil {
			return nil, fmt.Errorf("Could not assemble DB item: %v", err)
		}

		items = append(items, item)
	}

	return items, nil
}
