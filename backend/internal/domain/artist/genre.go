package artist

import (
	"errors"
	"strings"
)

type Genre struct {
	ID   int64  `json:"id"`
	Name string `json:"name"`
}

var (
	ErrEmptyGenreName = errors.New("Genre name must not be empty")
)

func NewGenre(name string) (Genre, error) {
	if strings.TrimSpace(name) == "" {
		return Genre{}, ErrEmptyGenreName
	}

	return Genre{Name: name}, nil
}
