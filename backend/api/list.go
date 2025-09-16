package api

import (
	"net/http"

	"github.com/mattismoel/konnekt/backend/order"
	"github.com/mattismoel/konnekt/backend/urlutil"
)

const (
	DEFAULT_PAGE      = 1
	DEFAULT_PAGE_SIZE = 12
)

type ListRequest struct {
	Page     int
	PageSize int
	Limit    int
	OrderMap order.Map
}

type ListResponse[T any] struct {
	Records []T `json:"records"`
}

func NewListRequest(r *http.Request) ListRequest {
	page, err := urlutil.QueryInt(r, "page")
	if err != nil {
		page = DEFAULT_PAGE
	}

	pageSize, err := urlutil.QueryInt(r, "page_size")
	if err != nil {
		pageSize = DEFAULT_PAGE_SIZE
	}

	limit, _ := urlutil.QueryInt(r, "limit")

	orderStr := r.URL.Query().Get("order_by")
	orderMap := order.MapFromString(orderStr)

	return ListRequest{
		Page:     page,
		Limit:    limit,
		PageSize: pageSize,
		OrderMap: orderMap,
	}
}

func (lr ListRequest) Offset() int {
	return (lr.Page - 1) * lr.Limit
}
