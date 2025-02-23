package server

import (
	"net/http"
	"strconv"

	"github.com/mattismoel/konnekt/internal/query"
)

func NewListQueryFromRequest(r *http.Request) query.ListQuery {
	page, _ := strconv.Atoi(r.URL.Query().Get("page"))
	perPage, _ := strconv.Atoi(r.URL.Query().Get("perPage"))
	limit, _ := strconv.Atoi(r.URL.Query().Get("limit"))

	return query.NewListQuery(page, perPage, limit)
}
