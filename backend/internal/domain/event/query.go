package event

import (
	"errors"
	"time"

	"github.com/mattismoel/konnekt/internal/query"
)

type Query struct {
	query.ListQuery
	From      time.Time
	To        time.Time
	ArtistIDs []int64
}

type QueryCfg func(q *Query) error

func (e *Query) WithCfgs(cfgs ...QueryCfg) error {
	for _, cfg := range cfgs {
		if err := cfg(e); err != nil {
			return err
		}
	}

	return nil
}

func NewQuery(cfgs ...QueryCfg) (Query, error) {
	e := &Query{}

	if err := e.WithCfgs(cfgs...); err != nil {
		return Query{}, err
	}

	return *e, nil
}

func WithPagination(query query.ListQuery) QueryCfg {
	return func(q *Query) error {
		q.ListQuery = query
		return nil
	}
}

func WithFromDate(d time.Time) QueryCfg {
	return func(q *Query) error {
		if d.IsZero() {
			return errors.New("Invalid from date")
		}

		q.From = d
		return nil
	}
}

func WithToDate(d time.Time) QueryCfg {
	return func(q *Query) error {
		if d.IsZero() {
			return errors.New("Invalid to date")
		}

		if !q.To.IsZero() && d.Before(q.From) {
			return errors.New("Invalid date relationship")
		}

		q.To = d
		return nil
	}
}

func WithArtistIDs(ids ...int64) QueryCfg {
	return func(q *Query) error {
		q.ArtistIDs = ids
		return nil
	}
}
