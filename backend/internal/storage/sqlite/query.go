package sqlite

import (
	"database/sql"
	"errors"
	"strings"
)

var (
	ErrInvalidBaseQuery    = errors.New("Base query must not include filters, offsets or limits")
	ErrInvalidFilterString = errors.New("Filter string must be of format 'filter [=,<,>,<=,>=,!=] ?.")
)

type QueryParams struct {
	// The offset of which to apply to the query.
	Offset int
	// The limit of which to apply to the query.
	Limit int
}

// A SQLite query builder instance.
type Query struct {
	baseQuery string
	args      []any
	filters   map[string]any
	offset    int
	limit     int
}

// Builds a new query.
//
// The query must not have filters, offsets or limits.
func NewQuery(baseQuery string) (Query, error) {
	if strings.TrimSpace(baseQuery) == "" {
		return Query{}, ErrInvalidBaseQuery
	}

	// Check if base query contains limit, offset or filters.
	if strings.Contains(strings.ToUpper(baseQuery), " OFFSET ") {
		return Query{}, ErrInvalidBaseQuery
	}

	if strings.Contains(strings.ToUpper(baseQuery), " LIMIT ") {
		return Query{}, ErrInvalidBaseQuery
	}

	if strings.Contains(strings.ToUpper(baseQuery), " WHERE ") {
		return Query{}, ErrInvalidBaseQuery
	}

	return Query{
		baseQuery: strings.TrimSpace(baseQuery + "\n" + "WHERE 1=1"),
		args:      make([]any, 0),
		filters:   make(map[string]any),
	}, nil
}

// Applies a limit to the query.
func (q *Query) WithLimit(limit int) error {
	if limit < 0 {
		return errors.New("Limit must be non-negative")
	}

	q.limit = limit

	return nil
}

// Applies an offset to the query.
func (q *Query) WithOffset(offset int) error {
	if offset < 0 {
		return errors.New("Offset must be non-negative")
	}

	q.offset = offset

	return nil
}

// Applies filters to the query. The filter key is the filter string, e.g.
//
// a = ?
// a != ?
// a >= ?
// a <= ?
//
// The filter value is the value to replace the placeholder '?' with.
func (q *Query) WithFilters(filters map[string]any) error {
	for filterString, value := range filters {
		err := q.AddFilter(filterString, value)
		if err != nil {
			return err
		}
	}

	return nil
}

// Appends a single filter to the query. The filterString must be of format:
// {property} {conditional} ?
// i.e. "a >= ?"
//
// The value is the value of which to replace the placeholder '?' with when
// the query is built with Query.Build().
func (q *Query) AddFilter(filterString string, value any) error {
	if !isValidFilterString(filterString) {
		return ErrInvalidFilterString
	}

	q.filters[filterString] = value

	return nil
}

// Adds a string line to the query.
func (q *Query) AddLine(s string) {
	q.baseQuery = strings.TrimSpace(q.baseQuery)
	q.baseQuery += "\n"
	q.baseQuery += s
}

// Builds the query into a query string and the arguments.
//
// Example usage:
//
//	query, _ := NewQuery("SELECT a, b FROM table")
//	_ = query.WithFilters(map[string]any{"a > ?": 10})
//
//	queryString, args := query.Build()
//
//	db.QueryContext(ctx, queryString, args...)
func (q Query) Build() (string, []any) {
	if len(q.filters) > 0 {
		for filterString, value := range q.filters {
			q.AddLine("AND " + filterString)
			q.args = append(q.args, value)
		}
	}

	if q.limit > 0 {
		q.AddLine("LIMIT @limit")
		q.args = append(q.args, sql.Named("limit", q.limit))
	}

	if q.offset >= 0 && q.limit > 0 {
		q.AddLine("OFFSET @offset")
		q.args = append(q.args, sql.Named("offset", q.offset))
	}

	return q.baseQuery, q.args
}

func isValidFilterString(s string) bool {
	return strings.ContainsRune(s, '?')
}
