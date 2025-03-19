package server

import (
	"errors"
	"fmt"
	"net/url"
	"strconv"
	"strings"

	"github.com/mattismoel/konnekt/internal/query"
)

func NewListQueryFromURL(vals url.Values) (query.ListQuery, error) {
	page := parsePage(vals)
	perPage := parsePerPage(vals)
	limit := parseLimit(vals)

	orderMap, err := orderMapFromRequest(vals)
	if err != nil {
		return query.ListQuery{}, err
	}

	filters, err := parseFilters(vals)
	if err != nil {
		return query.ListQuery{}, err
	}

	q, err := query.NewListQuery(
		query.WithPage(page),
		query.WithPerPage(perPage),
		query.WithLimit(limit),
		query.WithOrders(orderMap),
		query.WithFilters(filters),
	)

	if err != nil {
		return query.ListQuery{}, err
	}

	return q, nil
}

func orderMapFromRequest(vals url.Values) (map[string]query.Order, error) {
	m := make(map[string]query.Order)

	orderBy := vals.Get("order_by")
	orderParts := strings.Split(orderBy, ",") // ["title", "created_at desc"]

	isOrderByPresent := len(orderParts) > 0 && strings.TrimSpace(orderParts[0]) != ""

	if isOrderByPresent {
		for _, part := range orderParts { // each | "prop order"|
			key, order, err := orderByStrToPair(part)
			if err != nil {
				return nil, err
			}

			m[key] = order
		}
	}

	return m, nil
}

// Converts a search param order by pair to a usable order key and order.
//
// Examples:
//
//	orderByStrToPair("title") -> "title", query.OrderAscending, nil
//	orderByStrToPair("description desc") -> "description", query.OrderDescending, nil
//	orderByStrToPair("") -> "", "", <error>
func orderByStrToPair(s string) (string, query.Order, error) {
	if strings.TrimSpace(s) == "" {
		return "", "", errors.New(fmt.Sprintf("Order pair must be valid. Got %q", s))
	}

	key, orderStr := func() (string, string) {
		parts := strings.Split(strings.TrimSpace(s), " ")
		if len(parts) == 1 {
			return parts[0], ""
		}

		return parts[0], parts[1]
	}()

	order := query.Order(orderStr)
	if !order.Valid() {
		return key, query.OrderAscending, nil
	}

	return key, order, nil
}

// "filter=id!=4,title=what, hello>world"
func parseFilters(vals url.Values) ([]query.Filter, error) {
	filterStr := vals.Get("filter")
	if filterStr == "" {
		return make([]query.Filter, 0), nil
	}

	filters := make([]query.Filter, 0)
	filterPairs := strings.SplitSeq(filterStr, ",")

	for pair := range filterPairs {
		var cmp query.Comparator
		switch {
		case strings.Contains(pair, "!="):
			cmp = query.NotEqual
		case strings.Contains(pair, ">="):
			cmp = query.GreaterThanEqual
		case strings.Contains(pair, "<="):
			cmp = query.LessThanEqual
		case strings.Contains(pair, ">"):
			cmp = query.GreaterThan
		case strings.Contains(pair, "<"):
			cmp = query.LessThan
		case strings.Contains(pair, "="):
			cmp = query.Equal
		default:
			return make([]query.Filter, 0), fmt.Errorf("Invalid filter format in pair %q", pair)
		}

		parts := strings.SplitN(pair, string(cmp), 2)
		if len(parts) != 2 {
			return make([]query.Filter, 0), fmt.Errorf("Invalid filter format in pair %q", pair)
		}

		filter, err := query.NewFilter(parts[0], cmp, parts[1])
		if err != nil {
			return make([]query.Filter, 0), err
		}

		filters = append(filters, filter)
	}

	return filters, nil
}

// Returns the page entry of a url. If not found, zero is returned
func parsePage(vals url.Values) int {
	page, _ := strconv.Atoi(vals.Get("page"))
	return page
}

// Returns the perPage entry of a url. If not found, zero is returned.
func parsePerPage(vals url.Values) int {
	perPage, _ := strconv.Atoi(vals.Get("perPage"))
	return perPage
}

func parseLimit(vals url.Values) int {
	limit, _ := strconv.Atoi(vals.Get("limit"))
	return limit
}
