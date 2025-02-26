package sqlite

import (
	"database/sql"
	"errors"
	"strings"
)

type QueryParams struct {
	Offset int
	Limit  int
}

type BaseQuery struct {
	query string
	args  []any
}

func NewQuery(query string) (BaseQuery, error) {
	if strings.TrimSpace(query) == "" {
		return BaseQuery{}, errors.New("Query must not be empty")
	}

	return BaseQuery{
		query: query,
		args:  make([]any, 0),
	}, nil
}

func (q BaseQuery) WithLimit(limit int) error {
	if limit < 0 {
		return errors.New("Limit must be non-negative")
	}

	q.AddLine("LIMIT @limit")

	q.args = append(q.args, sql.Named("limit", limit))

	return nil
}

func (q BaseQuery) WithOffset(offset int) error {
	if offset < 0 {
		return errors.New("Offset must be non-negative")
	}

	q.AddLine("OFFSET @offset")

	q.args = append(q.args, sql.Named("offset", offset))

	return nil
}

func (q BaseQuery) AddLine(s string) {
	q.query = strings.TrimSpace(q.query)
	q.query += "\n"
	q.query += s
}

func (q BaseQuery) Query() (string, []any) {
	return q.query, q.args
}
