package konnekt

import (
	"slices"
	"strings"
	"time"
)

const DATE_PRECISION = time.Minute

type Event struct {
	ID          int64     `json:"id"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	FromDate    time.Time `json:"fromDate"`
	ToDate      time.Time `json:"toDate"`
	Address     Address   `json:"address"`
	Genres      []Genre   `json:"genres"`
}

func (e Event) Validate() error {
	if e.ID < 0 {
		return Errorf(ERRINVALID, "ID must not be negative")
	}

	if strings.TrimSpace(e.Title) == "" {
		return Errorf(ERRINVALID, "Title must be set")
	}

	if strings.TrimSpace(e.Description) == "" {
		return Errorf(ERRINVALID, "Description must be set")
	}

	if e.FromDate.IsZero() {
		return Errorf(ERRINVALID, "FromDate must be set")
	}
	if e.ToDate.IsZero() {
		return Errorf(ERRINVALID, "ToDate must be set")
	}

	if e.Genres != nil && len(e.Genres) == 0 {
		return Errorf(ERRINVALID, "Genre count must be at least 1")
	}

	return nil
}

type EventFilter struct {
	ID         *int64    `json:"id"`
	ArtistName *string   `json:"artistName"`
	MinDate    time.Time `json:"minDate"`
	MaxDate    time.Time `json:"maxDate"`

	Offset int `json:"offset"`
	Limit  int `json:"limit"`
}

type EventUpdate struct {
	Title       *string        `json:"title"`
	Description *string        `json:"description"`
	FromDate    time.Time      `json:"fromDate"`
	ToDate      time.Time      `json:"toDate"`
	Address     *AddressUpdate `json:"address"`
	GenreNames  []string       `json:"genreNames"`
}

func (e Event) Equals(a Event) bool {
	if e.Title != a.Title {
		return false
	}

	if e.Description != a.Description {
		return false
	}

	if !e.FromDate.Truncate(DATE_PRECISION).Equal(a.FromDate.Truncate(DATE_PRECISION)) {
		return false
	}

	if !e.ToDate.Truncate(DATE_PRECISION).Equal(a.ToDate.Truncate(DATE_PRECISION)) {
		return false
	}

	if !e.Address.Equals(a.Address) {
		return false
	}

	if !slices.Equal(e.Genres, a.Genres) {
		return false
	}

	return true
}
