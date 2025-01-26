package artist

import (
	"errors"
	"strings"
)

type Genre string

var (
	ErrEmptyGenreName = errors.New("Genre name must not be empty")
)

func NewGenre(name string) (Genre, error) {
	if strings.TrimSpace(name) == "" {
		return "", ErrEmptyGenreName
	}

	return Genre(name), nil
}
