package query

import (
	"errors"
	"strings"
)

var (
	ErrFilterKeyInvalid   = errors.New("Filter key must be a valid non-empty string")
	ErrFilterCmpInvalid   = errors.New("Filter comparable must be { <, >, <=, >=, = or != }")
	ErrFilterValueInvalid = errors.New("Filter value must not be empty")
)

type Comparator string

const (
	GreaterThan      = Comparator(">")
	LessThan         = Comparator("<")
	LessThanEqual    = Comparator("<=")
	GreaterThanEqual = Comparator(">=")
	Equal            = Comparator("=")
	NotEqual         = Comparator("!=")
)

type Filter struct {
	Key   string
	Cmp   Comparator
	Value string
}

type FilterCollection []Filter

func NewFilter(key string, cmp Comparator, value string) (Filter, error) {
	if err := validateFilterKey(key); err != nil {
		return Filter{}, err
	}

	if !cmp.valid() {
		return Filter{}, ErrFilterCmpInvalid
	}

	if err := validateFilterValue(value); err != nil {
		return Filter{}, err
	}

	return Filter{Key: key, Cmp: cmp, Value: value}, nil
}

// Returns whether or not the input filter key is valid, i.e. a single property.
// If not, a validation error is returned.
func validateFilterKey(key string) error {
	key = strings.TrimSpace(key)
	if key == "" {
		return ErrFilterKeyInvalid
	}

	keyParts := strings.Split(key, " ")
	if len(keyParts) > 1 {
		return ErrFilterKeyInvalid
	}

	keyParts = strings.Split(key, ",")
	if len(keyParts) > 1 {
		return ErrFilterKeyInvalid
	}

	return nil
}

// Returns whether or not the input filter key is valid, i.e. a single property.
// If not, a validation error is returned.
func validateFilterValue(value string) error {
	value = strings.TrimSpace(value)
	if value == "" {
		return ErrFilterValueInvalid
	}

	valueParts := strings.Split(value, " ")
	if len(valueParts) > 1 {
		return ErrFilterValueInvalid
	}

	valueParts = strings.Split(value, ",")
	if len(valueParts) > 1 {
		return ErrFilterValueInvalid
	}

	return nil
}

func WithFilters(filters FilterCollection) CfgFunc {
	return func(q *ListQuery) error {
		for _, f := range filters {
			if err := validateFilterKey(f.Key); err != nil {
				return err
			}

			if err := validateFilterValue(f.Value); err != nil {
				return err
			}

			if !f.Cmp.valid() {
				return ErrFilterCmpInvalid
			}
		}

		q.Filters = filters

		return nil
	}
}

func (c Comparator) valid() bool {
	return c == GreaterThan || c == LessThan || c == GreaterThanEqual || c == LessThanEqual || c == Equal
}

// Checks whether the filter collection contains a filter entry with the given
// key.
//
// If found, the filter is returned, as well as a true boolean.
// If not found, an uninitialized filter, and a false boolean is returned.
func (fs FilterCollection) Contains(key string) (Filter, bool) {
	for _, f := range fs {
		if f.Key == key {
			return f, true
		}
	}

	return Filter{}, false
}

func (f1 Filter) Equals(f2 Filter) bool {
	if f1.Key != f2.Key {
		return false
	}

	if f1.Value != f2.Value {
		return false
	}

	if f1.Cmp != f2.Cmp {
		return false
	}

	return true
}
