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
	Cmp   Comparator
	Value string
}

// A collection of key-to-filter entries.
type FilterCollection map[string][]Filter

func NewFilter(cmp Comparator, value string) (Filter, error) {
	if !cmp.valid() {
		return Filter{}, ErrFilterCmpInvalid
	}

	if err := validateFilterValue(value); err != nil {
		return Filter{}, err
	}

	return Filter{
		Cmp:   cmp,
		Value: value,
	}, nil
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

func WithFilters(fc FilterCollection) CfgFunc {
	return func(q *ListQuery) error {
		for key, filters := range fc {
			if err := validateFilterKey(key); err != nil {
				return err
			}

			if err := q.Filters.Add(key, filters...); err != nil {
				return err
			}
		}

		return nil
	}
}

func (c Comparator) valid() bool {
	return c == GreaterThan || c == LessThan || c == GreaterThanEqual || c == LessThanEqual || c == Equal || c == NotEqual
}

// Returns if two filters are equal.
func (f1 Filter) Equals(f2 Filter) bool {
	if f1.Value != f2.Value {
		return false
	}

	if f1.Cmp != f2.Cmp {
		return false
	}

	return true
}

// Adds a new filter to the given key entry of the FilterCollection.
func (fc FilterCollection) Add(key string, filters ...Filter) error {
	for _, f := range filters {
		if err := validateFilterValue(f.Value); err != nil {
			return err
		}

		if !f.Cmp.valid() {
			return ErrFilterCmpInvalid
		}

		fc[key] = append(fc[key], f)
	}

	return nil
}

func (fc1 FilterCollection) Equals(fc2 FilterCollection) bool {
	if len(fc1) != len(fc2) {
		return false
	}

	for key, fs1 := range fc1 {
		if _, ok := fc2[key]; !ok {
			return false
		}

		for i, f := range fs1 {
			if f.Value != fc2[key][i].Value {
				return false
			}
		}
	}

	return true
}
