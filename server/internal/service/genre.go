package service

import (
	"context"
	"strings"

	"github.com/mattismoel/konnekt"
	"github.com/mattismoel/konnekt/internal/storage"
)

type Genre string

type genreRepository interface {
	InsertGenre(ctx context.Context, name string) (storage.Genre, error)
	DeleteGenre(ctx context.Context, id int64) error
	FindGenres(ctx context.Context, filter GenreFilter) ([]storage.Genre, error)
}

type GenreFilter struct {
	ID      *int64  `json:"id"`
	Name    *string `json:"name"`
	EventID *int64  `json:"eventId"`

	Limit  int `json:"limit"`
	Offset int `json:"offset"`
}

type genreService struct {
	repo genreRepository
}

func NewGenreService(repo genreRepository) *genreService {
	return &genreService{repo: repo}
}

func (s genreService) CreateGenre(ctx context.Context, name string) (Genre, error) {
	if strings.TrimSpace(name) == "" {
		return "", konnekt.Errorf(konnekt.ERRINVALID, "Genre name must not be empty")
	}

	genre, err := s.repo.InsertGenre(ctx, name)
	if err != nil {
		return "", err
	}

	return Genre(genre.Name), nil
}

func (s genreService) DeleteGenre(ctx context.Context, id int64) error {
	err := s.repo.DeleteGenre(ctx, id)
	if err != nil {
		return err
	}

	return nil
}

func (s genreService) FindGenres(ctx context.Context, filter GenreFilter) ([]Genre, error) {
	genres, err := s.repo.FindGenres(ctx, filter)
	if err != nil {
		return nil, err
	}

	names := []Genre{}

	for _, genre := range genres {
		names = append(names, Genre(genre.Name))
	}

	return names, nil
}

func (g Genre) Validate() error {
	if strings.TrimSpace(string(g)) == "" {
		return konnekt.Errorf(konnekt.ERRINVALID, "Name must not be empty")
	}

	return nil
}

func (g Genre) Equals(a Genre) bool {
	if g != a {
		return false
	}

	return true
}
