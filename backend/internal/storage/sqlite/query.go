package sqlite

import (
	"database/sql"
	"errors"
	"fmt"
	"strings"

	"github.com/mattismoel/konnekt/internal/query"
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
	// The orderings to apply to the query.
	OrderBy map[string]query.Order
	// The filters to apply to the query.
	Filters query.FilterCollection
}

// A SQLite query builder instance.
type Query struct {
	baseQuery string
	args      []any
	orderMap  map[string]query.Order
	filters   query.FilterCollection
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
		filters:   make([]query.Filter, 0),
		orderMap:  make(map[string]query.Order),
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

func (q *Query) WithOrdering(orderMap map[string]query.Order) error {
	if orderMap == nil {
		return errors.New("The passed order map must not be nil")
	}

	for key, order := range orderMap {
		if strings.ToUpper(string(order)) != "ASC" && strings.ToUpper(string(order)) != "DESC" {
			return errors.New("Order must be 'ASC' or 'DESC' (case-insensitive)")
		}

		if strings.TrimSpace(key) == "" {
			return errors.New("The ordering property must not be empty")
		}
	}

	q.orderMap = orderMap

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
func (q *Query) WithFilters(filters query.FilterCollection) error {
	q.filters = append(q.filters, filters...)

	return nil
}

// Appends a single filter to the query. The filterString must be of format:
// {property} {conditional} ?
// i.e. "a >= ?"
//
// The value is the value of which to replace the placeholder '?' with when
// the query is built with Query.Build().
func (q *Query) AddFilter(filter query.Filter) error {
	q.filters = append(q.filters, filter)
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
		for _, filter := range q.filters {
			filterStr := fmt.Sprintf("filter:%s %s ?", filter.Key, string(filter.Cmp))
			q.AddLine("AND " + filterStr)
			q.args = append(q.args, filter.Value)
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

	if len(q.orderMap) > 0 {
		clauses := []string{}
		for key, order := range q.orderMap {
			clauses = append(clauses, fmt.Sprintf("%s %s", key, order))
		}

		q.AddLine("ORDER BY " + strings.Join(clauses, ", "))
	}

	fmt.Printf("QUERY: %s\n", q.baseQuery)
	return q.baseQuery, q.args
}

func isValidFilterString(s string) bool {
	return strings.ContainsRune(s, '?')
}

// Converts an internal type of map[string]query.Order to an SQLite
// representable map of type map[string]string.
func OrderingMapFromInternal(m map[string]query.Order) map[string]string {
	orderMap := make(map[string]string)

	for key, order := range m {
		orderMap[key] = string(order)
	}

	return orderMap
}
