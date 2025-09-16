package psql

import (
	"context"
	"fmt"

	sq "github.com/Masterminds/squirrel"
	"github.com/jackc/pgx/v5"
)

var psql = sq.StatementBuilder.PlaceholderFormat(sq.Dollar)

type Assembler[T any, K any] interface {
	Assemble(context.Context, pgx.Tx) (K, error)
}

func checkIfExists(ctx context.Context, tx pgx.Tx, table string, key string, idValue any) (bool, error) {
	query, args, err := psql.
		Select("1").
		Prefix("SELECT EXISTS (").
		From(table).
		Where(sq.Eq{key: idValue}).
		Suffix(")").
		ToSql()

	if err != nil {
		return false, NewQueryBuildError("check if exists", err)
	}

	rows, err := tx.Query(ctx, query, args...)
	if err != nil {
		return false, fmt.Errorf("Could not check if %q with %s=%v exists: %v", table, key, idValue, err)
	}

	exists, err := pgx.CollectExactlyOneRow(rows, pgx.RowTo[bool])
	if err != nil {
		return false, fmt.Errorf("Could not collect 'exists' row: %v", err)
	}

	return exists, nil
}
