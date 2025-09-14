package psql

import (
	sq "github.com/Masterminds/squirrel"
	"github.com/mattismoel/konnekt/backend/api"
	"github.com/mattismoel/konnekt/backend/order"
)

type Pagination struct {
	Offset int
	Limit  int
}

func paginationFromListRequest(lr api.ListRequest) Pagination {
	return Pagination{
		Limit:  lr.Limit,
		Offset: lr.Offset(),
	}
}

func applyPagination(sb sq.SelectBuilder, pg Pagination) sq.SelectBuilder {
	sb = sb.
		Offset(uint64(pg.Offset))

	if pg.Limit > 0 {
		sb = sb.Limit(uint64(pg.Limit))
	}

	return sb
}

func applyOrdering(sb sq.SelectBuilder, orderMap order.Map) sq.SelectBuilder {
	for key, order := range orderMap {
		sb = sb.OrderBy(key, string(order))
	}

	return sb
}
