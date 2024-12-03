package konnekt

import (
	"context"
	"strings"
)

type Genre struct {
	ID   int64  `json:"id"`
	Name string `json:"name"`
}

type GenreService interface {
	GenreByID(context.Context, int64) (Genre, error)
	FindGenres(context.Context, GenreFilter) ([]Genre, error)
	CreateGenre(context.Context, Genre) (int64, error)
	UpdateGenre(context.Context, int64, GenreUpdate) (Genre, error)
	DeleteGenre(context.Context, int64) error
}

type GenreFilter struct {
	ID      *int64  `json:"id"`
	Name    *string `json:"name"`
	EventID *int64  `json:"eventId"`

	Limit  int `json:"limit"`
	Offset int `json:"offset"`
}

type GenreUpdate struct {
	Name *string `json:"name"`
}

func (g Genre) Validate() error {
	if g.ID < 0 {
		return Errorf(ERRINVALID, "ID must be a positive integer")
	}

	if strings.TrimSpace(g.Name) == "" {
		return Errorf(ERRINVALID, "Name must not be empty")
	}

	return nil
}

func (g Genre) Equals(a Genre) bool {
	if g.ID != a.ID {
		return false
	}

	if g.Name != a.Name {
		return false
	}

	return true
}
