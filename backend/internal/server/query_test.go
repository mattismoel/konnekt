package server_test

import (
	"errors"
	"net/url"
	"strconv"
	"testing"

	"github.com/mattismoel/konnekt/internal/query"
	"github.com/mattismoel/konnekt/internal/server"
)

const BASE_URL = "https://example.com"

func TestNewListQueryFromRequest(t *testing.T) {
	baseQuery := query.ListQuery{
		Page:    query.DEFAULT_PAGE,
		PerPage: query.DEFAULT_PER_PAGE,
		Limit:   query.DEFAULT_LIMIT,
		OrderBy: make(map[string]query.Order),
		Filters: make(query.FilterCollection, 0),
	}

	baseUrl, err := url.Parse(BASE_URL)
	if err != nil {
		t.Fatal(err)
	}

	type modFunc func(q query.ListQuery) query.ListQuery

	type test struct {
		params       map[string]string
		wantQueryMod modFunc
		wantErr      error
	}

	tests := map[string]test{
		"Valid query string (With all base params)": {
			params: map[string]string{
				"limit":   strconv.Itoa(12),
				"page":    strconv.Itoa(2),
				"perPage": strconv.Itoa(4),
				"orderBy": "prop_a,prop_b desc",
				"filter":  "prop_a=4,prop_b>=3,prop_c!=2",
			},
			wantQueryMod: func(q query.ListQuery) query.ListQuery {
				q.Limit = 12
				q.Page = 2
				q.PerPage = 4
				q.OrderBy = map[string]query.Order{
					"prop_a": query.OrderAscending,
					"prop_b": query.OrderDescending,
				}
				q.Filters = map[string][]query.Filter{
					"prop_a": {{Cmp: query.Equal, Value: "4"}},
					"prop_b": {{Cmp: query.GreaterThanEqual, Value: "3"}},
					"prop_c": {{Cmp: query.NotEqual, Value: "2"}},
				}
				return q
			},
			wantErr: nil,
		},
	}

	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			vals := baseUrl.Query()

			for key, value := range tt.params {
				vals.Add(key, value)
			}

			query, err := server.NewListQueryFromURL(vals)
			if !errors.Is(err, tt.wantErr) {
				t.Fatalf("got %v, want %v with url %v", err, tt.wantErr, vals)
			}

			wantQuery := tt.wantQueryMod(baseQuery)
			if !query.Equals(wantQuery) {
				t.Fatalf("\ngot:\t%+v\nwant:\t%+v", query, wantQuery)
			}
		})
	}
}
