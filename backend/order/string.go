package order

import "strings"

// Creates an order map from the input order_by string.
//
// Example: "prop_a, prop_b desc, prop_c asc"
// -> { "prop_a": "ASC", "prop_b": "DESC", "prop_c": "ASC" }
func MapFromString(s string) Map {
	orderMap := make(Map)
	if s == "" {
		return orderMap
	}

	orderStrs := strings.Split(s, ",")
	for _, orderStr := range orderStrs {
		parts := strings.Split(orderStr, " ")
		key := parts[0]

		if len(parts) == 1 {
			orderMap[key] = ASCENDING
			continue
		}

		order := Order(strings.ToUpper(parts[1]))
		if order != ASCENDING && order != DESCENDING {
			orderMap[key] = ASCENDING
			continue
		}

		orderMap[key] = order
	}

	return orderMap
}
