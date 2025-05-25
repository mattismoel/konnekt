package sqlite

import (
	"context"
	"database/sql"

	sq "github.com/Masterminds/squirrel"
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
