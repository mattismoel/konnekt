package query

const (
	// The default page to query for.
	DEFAULT_PAGE = 1
	// The default amount of results on a page.
	DEFAULT_PER_PAGE = 8
	// The max amount to query for on a page.
	MAX_PER_PAGE = 100
)

type ListQuery struct {
	// The page to be retrived.
	Page int
	// The maximum amount of results on a retrieved page.
	PerPage int
	// The maximum amount of total results.
	Limit int
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

// Creates a new list query.
func NewListQuery(page, perPage, limit int) ListQuery {
	if page <= 0 {
		page = DEFAULT_PAGE
	}

	if perPage <= 0 {
		perPage = DEFAULT_PER_PAGE
	}

	if perPage > MAX_PER_PAGE {
		perPage = MAX_PER_PAGE
	}

	if limit <= 0 {
		limit = 1
	}

	return ListQuery{
		Page:    page,
		PerPage: perPage,
		Limit:   limit,
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
