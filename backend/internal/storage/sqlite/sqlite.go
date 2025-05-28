package sqlite

import (
	"context"
	"database/sql"
	"fmt"

	sq "github.com/Masterminds/squirrel"
	"github.com/mattismoel/konnekt/internal/query"
)

type Scanner interface {
	Scan(dst ...any) error
}

func count(ctx context.Context, tx *sql.Tx, table string) (int, error) {
	query, args, err := sq.
		Select("COUNT(*)").
		From(table).
		ToSql()

	if err != nil {
		return 0, err
	}

	var count int
	err = tx.QueryRowContext(ctx, query, args...).Scan(&count)
	if err != nil {
		return 0, err
	}

	return count, nil
}

func withPagination(b sq.SelectBuilder, params QueryParams) sq.SelectBuilder {
	if params.Limit > 0 {
		b = b.Limit(uint64(params.Offset))
	}

	if params.Offset > 0 {
		b = b.Offset(uint64(params.Offset))
	}

	return b
}

