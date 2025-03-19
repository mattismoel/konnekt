package query

import (
	"slices"
)

const (
	// The default page to query for.
	DEFAULT_PAGE = 1

	// The default amount of results on a page.
	DEFAULT_PER_PAGE = 8

	// The default limit of a list result.
	//
	// Zero-value is regarded as "no limit".
	DEFAULT_LIMIT = 0

	// The max amount to query for on a page.
	MAX_PER_PAGE = 100
)

type Order string

const (
	OrderAscending  = Order("ASC")
	OrderDescending = Order("DESC")
)

type ListQuery struct {
	// The page to be retrived.
	Page int

	// The maximum amount of results on a retrieved page.
	PerPage int

	// The maximum amount of total results.
	Limit int
	// The ordering to apply to the results.
	//
	// Example:
	//	OrderBy: map[string]string{
	//		"created_at": OrderAscending,
	//		"name": OrderDescending,
	//	}
	OrderBy map[string]Order

	// The filters part of this query.
	Filters FilterCollection
}

type ListResult[T any] struct {
	// The page of which the resulting records are from.
	Page int `json:"page"`
	// The amount of records per page.
	PerPage int `json:"perPage"`
	// The total amount of records available for retrieving.
	TotalCount int `json:"totalCount"`
	// The amount of pages available for querying.
	PageCount int `json:"pageCount"`
	// The records of the page.
	Records []T `json:"records"`
}

type CfgFunc func(q *ListQuery) error

// Creates a new list query.
func NewListQuery(cfgs ...CfgFunc) (ListQuery, error) {
	q := &ListQuery{
		Page:    DEFAULT_PAGE,
		PerPage: DEFAULT_PER_PAGE,
		Limit:   DEFAULT_LIMIT,
		OrderBy: map[string]Order{},
		Filters: make([]Filter, 0),
	}

	for _, cfg := range cfgs {
		if err := cfg(q); err != nil {
			return ListQuery{}, err
		}
	}

	return *q, nil
}

func WithPage(page int) CfgFunc {
	return func(q *ListQuery) error {
		if page <= 0 {
			q.Page = DEFAULT_PAGE
			return nil
		}

		q.Page = page

		return nil
	}
}

func WithPerPage(perPage int) CfgFunc {
	return func(q *ListQuery) error {
		if perPage <= 0 {
			q.PerPage = DEFAULT_PER_PAGE
			return nil
		}

		if perPage > MAX_PER_PAGE {
			q.PerPage = MAX_PER_PAGE
			return nil
		}

		q.PerPage = perPage

		return nil
	}
}

func WithLimit(limit int) CfgFunc {
	return func(q *ListQuery) error {
		if limit < 0 {
			q.Limit = DEFAULT_LIMIT
			return nil
		}

		q.Limit = limit
		return nil
	}
}

// Applies ordering to a ListQuery.
//
// Example:
//
//	q.WithOrder(map[string]Order{
//		"name": OrderAscending,
//		"created", OrderDescending,
//	})
func WithOrders(orderMap map[string]Order) CfgFunc {
	return func(q *ListQuery) error {
		for key, order := range orderMap {
			// Set to ascending order on invalid ordering
			if !order.Valid() {
				orderMap[key] = OrderAscending
			}
		}

		q.OrderBy = orderMap

		return nil
	}
}

// Returns the page count based on the total count of records.
func (q ListQuery) PageCount(totalCount int) int {
	return ((totalCount - 1) / q.PerPage) + 1
}

// Returns the offset.
func (q ListQuery) Offset() int {
	return (q.Page - 1) * q.PerPage
}

// Returns whether or not an input ordering is allowed, given the allowed keys.
func IsOrderingAllowed(orderKey string, allowedKeys ...string) bool {
	return slices.Contains(allowedKeys, orderKey)
}

// Checks whether the Order string is valid, i.e. is either "ASC" or "DESC"
func (o Order) Valid() bool {
	return o == OrderAscending || o == OrderDescending
}

// Checks whether or not two list queries are equal.
func (q1 ListQuery) Equals(q2 ListQuery) bool {
	if q1.Page != q2.Page {
		return false
	}

	if q1.PerPage != q2.PerPage {
		return false
	}

	if q1.Limit != q2.Limit {
		return false
	}

	if len(q1.Filters) != len(q2.Filters) {
		return false
	}

	for i, f1 := range q1.Filters {
		if !f1.Equals(q2.Filters[i]) {
			return false
		}
	}

	for key1, o1 := range q1.OrderBy {
		o2, ok := q2.OrderBy[key1]
		if !ok {
			return false
		}

		if o1 != o2 {
			return false
		}
	}

	return true
}
