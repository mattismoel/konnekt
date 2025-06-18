package concert

import (
	"errors"
	"time"

	"github.com/mattismoel/konnekt/internal/domain/artist"
)

var (
	ErrInvalidID               = errors.New("Concert ID must be a positive integer")
	ErrInvalidDateRelationship = errors.New("Concert dates must be concecutive")
	ErrInvalidDate             = errors.New("One or more dates are invalid or empty")
)

type CfgFunc func(c *Concert) error

type Concert struct {
	ID     int64         `json:"id"`
	From   time.Time     `json:"from"`
	To     time.Time     `json:"to"`
	Artist artist.Artist `json:"artist"`
}

func (c *Concert) WithCfgs(cfgs ...CfgFunc) error {
	for _, cfg := range cfgs {
		if err := cfg(c); err != nil {
			return err
		}
	}

	return nil
}

func NewConcert(cfgs ...CfgFunc) (Concert, error) {
	c := &Concert{}

	if err := c.WithCfgs(cfgs...); err != nil {
		return Concert{}, err
	}

	return *c, nil
}

func WithID(id int64) CfgFunc {
	return func(c *Concert) error {
		if id <= 0 {
			return ErrInvalidID
		}

		c.ID = id
		return nil
	}
}

func WithArtist(a artist.Artist) CfgFunc {
	return func(c *Concert) error {
		c.Artist = a
		return nil
	}
}

func WithFrom(from time.Time) CfgFunc {
	return func(c *Concert) error {
		if from.IsZero() {
			return ErrInvalidDate
		}

		if from.After(c.To) && !c.To.IsZero() {
			return ErrInvalidDateRelationship
		}

		c.From = from

		return nil
	}
}

func WithTo(to time.Time) CfgFunc {
	return func(c *Concert) error {
		if to.IsZero() {
			return ErrInvalidDate
		}

		if to.Before(c.From) && !c.From.IsZero() {
			return ErrInvalidDateRelationship
		}

		c.To = to

		return nil
	}
}
