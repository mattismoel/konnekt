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

func withOrdering(
	b sq.SelectBuilder,
	orderMap query.OrderMap,
	column string,
	table string,
) sq.SelectBuilder {
	order, ok := orderMap[column]
	if !ok {
		return b
	}

	tableColumn := fmt.Sprintf("%s.%s", table, column)
	if order == "ASC" || order == "DESC" {
		orderStr := fmt.Sprintf("%s %s", tableColumn, order)
		b = b.OrderBy(orderStr)
	} else {
		orderStr := fmt.Sprintf("%s ASC", tableColumn)
		b = b.OrderBy(orderStr)
	}

	return b
}

type filterFunc = func(query.Filter) sq.Sqlizer

func withFiltering(b sq.SelectBuilder, fc query.FilterCollection, fm map[string]filterFunc) sq.SelectBuilder {
	for key, fs := range fc {
		applyFn, ok := fm[key]
		if !ok {
			continue
		}

		for _, f := range fs {
			b = b.Where(applyFn(f))
		}
	}

	return b
}

func contains(column string, value any) sq.Like {
	return sq.Like{column: fmt.Sprintf("%%%s%%", value)}
}
