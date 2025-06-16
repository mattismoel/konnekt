package artist

import (
	"errors"
	"strings"
)

var (
	ErrInvalidGenreName = errors.New("Genre name must not be valid")
)

type Genre struct {
	ID   int64  `json:"id"`
	Name string `json:"name"`
}

func NewGenre(name string) (Genre, error) {
	if strings.TrimSpace(name) == "" {
		return Genre{}, ErrInvalidGenreName
	}

	return Genre{Name: name}, nil
}
