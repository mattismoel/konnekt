package query_test

import (
	"errors"
	"testing"

	"github.com/mattismoel/konnekt/internal/query"
)

var defaultQuery = query.ListQuery{
	Page:    query.DEFAULT_PAGE,
	PerPage: query.DEFAULT_PER_PAGE,
	Filters: make(query.FilterCollection, 0),
	OrderBy: make(map[string]query.Order),
}

func TestNewListQuery(t *testing.T) {
	type modFunc func(q query.ListQuery) query.ListQuery

	t.Logf("DEFAULT %+v\n", defaultQuery)

	type test struct {
		cfgs         []query.CfgFunc
		wantQueryMod modFunc
		wantErr      error
	}

	tests := map[string]test{
		"With valid page": {
			cfgs: []query.CfgFunc{
				query.WithPage(4),
			},
			wantQueryMod: func(q query.ListQuery) query.ListQuery {
				q.Page = 4
				return q
			},
			wantErr: nil,
		},
		"With page zero": {
			cfgs: []query.CfgFunc{
				query.WithPage(0),
			},
			wantQueryMod: func(q query.ListQuery) query.ListQuery {
				q.Page = query.DEFAULT_PAGE
				return q
			},
			wantErr: nil,
		},
		"With negative page": {
			cfgs: []query.CfgFunc{
				query.WithPage(-1),
			},
			wantQueryMod: func(q query.ListQuery) query.ListQuery {
				q.Page = query.DEFAULT_PAGE
				return q
			},
			wantErr: nil,
		},
		"With valid per page": {
			cfgs: []query.CfgFunc{
				query.WithPerPage(16),
			},
			wantQueryMod: func(q query.ListQuery) query.ListQuery {
				q.PerPage = 16
				return q
			},
			wantErr: nil,
		},
		"With zero per page": {
			cfgs: []query.CfgFunc{
				query.WithPerPage(0),
			},
			wantQueryMod: func(q query.ListQuery) query.ListQuery {
				q.PerPage = query.DEFAULT_PER_PAGE
				return q
			},
			wantErr: nil,
		},
		"With negative per page": {
			cfgs: []query.CfgFunc{
				query.WithPerPage(-1),
			},
			wantQueryMod: func(q query.ListQuery) query.ListQuery {
				q.PerPage = query.DEFAULT_PER_PAGE
				return q
			},
			wantErr: nil,
		},
		"With valid limit": {
			cfgs: []query.CfgFunc{
				query.WithLimit(4),
			},
			wantQueryMod: func(q query.ListQuery) query.ListQuery {
				q.Limit = 4
				return q
			},
			wantErr: nil,
		},
		"With negative limit": {
			cfgs: []query.CfgFunc{
				query.WithLimit(-1),
			},
			wantQueryMod: func(q query.ListQuery) query.ListQuery {
				q.Limit = query.DEFAULT_LIMIT
				return q
			},
			wantErr: nil,
		},
		"With valid filters": {
			cfgs: []query.CfgFunc{
				query.WithFilters(map[string][]query.Filter{
					"prop_a": {{Cmp: query.GreaterThan, Value: "2"}},
					"prop_b": {{Cmp: query.Equal, Value: "hello_world"}},
					"prop_c": {{Cmp: query.LessThan, Value: "d"}},
				}),
			},
			wantQueryMod: func(q query.ListQuery) query.ListQuery {
				q.Filters = map[string][]query.Filter{
					"prop_a": {{Cmp: query.GreaterThan, Value: "2"}},
					"prop_b": {{Cmp: query.Equal, Value: "hello_world"}},
					"prop_c": {{Cmp: query.LessThan, Value: "d"}},
				}
				return q
			},
			wantErr: nil,
		},
		"With invalid filter key": {
			cfgs: []query.CfgFunc{
				query.WithFilters(map[string][]query.Filter{
					"": {{Cmp: query.GreaterThan, Value: "2"}},
				}),
			},
			wantQueryMod: func(q query.ListQuery) query.ListQuery {
				return query.ListQuery{}
			},
			wantErr: query.ErrFilterKeyInvalid,
		},
		"With invalid filter cmp": {
			cfgs: []query.CfgFunc{
				query.WithFilters(map[string][]query.Filter{
					"prop_a": {{Cmp: query.Comparator(""), Value: "2"}},
				}),
			},
			wantQueryMod: func(q query.ListQuery) query.ListQuery {
				return query.ListQuery{}
			},
			wantErr: query.ErrFilterCmpInvalid,
		},
		"With invalid filter value": {
			cfgs: []query.CfgFunc{
				query.WithFilters(map[string][]query.Filter{
					"prop_a": {{Cmp: query.GreaterThan, Value: ""}},
				}),
			},
			wantQueryMod: func(q query.ListQuery) query.ListQuery {
				return query.ListQuery{}
			},
			wantErr: query.ErrFilterValueInvalid,
		},
		"With valid ordering": {
			cfgs: []query.CfgFunc{
				query.WithOrders(map[string]query.Order{
					"prop_a": query.OrderDescending,
				}),
			},
			wantQueryMod: func(q query.ListQuery) query.ListQuery {
				q.OrderBy = map[string]query.Order{"prop_a": query.OrderDescending}
				return q
			},
			wantErr: nil,
		},
		"With invalid order": {
			cfgs: []query.CfgFunc{
				query.WithOrders(map[string]query.Order{
					"prop_a": query.Order(""),
				}),
			},
			wantQueryMod: func(q query.ListQuery) query.ListQuery {
				q.OrderBy = map[string]query.Order{
					"prop_a": query.OrderAscending,
				}
				return q
			},
			wantErr: nil,
		},
	}

	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			query, err := query.NewListQuery(tt.cfgs...)
			if !errors.Is(err, tt.wantErr) {
				t.Fatalf("got %v, want %v", err, tt.wantErr)
			}

			wantQuery := tt.wantQueryMod(defaultQuery)
			if !query.Equals(wantQuery) {
				t.Fatalf("\ngot:\t%+v,\nwant:\t%+v", query, wantQuery)
			}
		})
	}
}
