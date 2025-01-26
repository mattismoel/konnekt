package concert

import (
	"errors"
	"time"

	"github.com/mattismoel/konnekt/internal/domain/artist"
)

var (
	ErrInvalidDateRelationship = errors.New("Concert dates must be concecutive")
	ErrInvalidDate             = errors.New("One or more dates are invalid or empty")
)

type Concert struct {
	ID     int64         `json:"id"`
	From   time.Time     `json:"from"`
	To     time.Time     `json:"to"`
	Artist artist.Artist `json:"artist"`
}

func NewConcert(a artist.Artist, from time.Time, to time.Time) (Concert, error) {
	if from.After(to) {
		return Concert{}, ErrInvalidDateRelationship
	}

	if from.IsZero() || to.IsZero() {
		return Concert{}, ErrInvalidDate
	}

	return Concert{
		Artist: a,
		From:   from,
		To:     to,
	}, nil
}
