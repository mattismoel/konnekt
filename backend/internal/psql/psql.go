package psql

import (
	"context"

	sq "github.com/Masterminds/squirrel"
	"github.com/jackc/pgx/v5"
)

var psql = sq.StatementBuilder.PlaceholderFormat(sq.Dollar)

type Assembler[T any, K any] interface {
	Assemble(context.Context, pgx.Tx) (K, error)
}
