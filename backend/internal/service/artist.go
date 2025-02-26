package service

import (
	"context"

	"github.com/mattismoel/konnekt/internal/domain/artist"
	"github.com/mattismoel/konnekt/internal/query"
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

type ArtistListQuery struct {
	query.ListQuery
}

func (s ArtistService) List(ctx context.Context, q ArtistListQuery) (query.ListResult[artist.Artist], error) {
	artists, totalCount, err := s.artistRepo.List(ctx, q.Offset(), q.Limit)
	if err != nil {
		return query.ListResult[artist.Artist]{}, err
	}

	return query.ListResult[artist.Artist]{
		Records:    artists,
		TotalCount: totalCount,
		Page:       q.Page,
		PerPage:    q.PerPage,
		PageCount:  q.PageCount(totalCount),
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

func (s ArtistService) Delete(ctx context.Context, artistID int64) error {
	err := s.artistRepo.Delete(ctx, artistID)
	if err != nil {
		return err
	}

	return nil
}
