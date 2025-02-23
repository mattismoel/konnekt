package service

import (
	"context"

	"github.com/mattismoel/konnekt/internal/domain/artist"
)

type ArtistService struct {
	artistRepo artist.Repository
}

func NewArtistService(artistRepo artist.Repository) (*ArtistService, error) {
	return &ArtistService{
		artistRepo: artistRepo,
	}, nil
}

type CreateArtist struct {
	Name        string
	Description string
	ImageURL    string
	GenreIDs    []int64
	Socials     []string
}

type ArtistQuery struct {
	Page    int
	PerPage int
}

type ArtistListResult struct {
	Page       int
	PerPage    int
	PageCount  int
	TotalCount int
	Artists    []artist.Artist
}

func (s ArtistService) List(ctx context.Context, query ArtistQuery) (ArtistListResult, error) {
	artists, totalCount, err := s.artistRepo.List(ctx)
	if err != nil {
		return ArtistListResult{}, err
	}

	pageCount := (totalCount + query.PerPage - 1) / query.PerPage

	return ArtistListResult{
		Artists:    artists,
		TotalCount: totalCount,
		Page:       query.Page,
		PerPage:    query.PerPage,
		PageCount:  pageCount,
	}, nil
}

func (s ArtistService) Create(ctx context.Context, load CreateArtist) (int64, error) {
	socials := make([]artist.Social, 0)
	for _, social := range load.Socials {
		s, err := artist.NewSocial(social)
		if err != nil {
			return 0, err
		}
		socials = append(socials, s)
	}

	genres := make([]artist.Genre, 0)
	for _, genreID := range load.GenreIDs {
		genre, err := s.artistRepo.GenreByID(ctx, genreID)
		if err != nil {
			return 0, err
		}

		genres = append(genres, genre)
	}

	a, err := artist.NewArtist(load.Name, load.Description, load.ImageURL, socials, genres)
	if err != nil {
		return 0, err
	}

	artistID, err := s.artistRepo.Insert(ctx, a)
	if err != nil {
		return 0, err
	}

	return artistID, nil
}
