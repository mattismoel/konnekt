package service

import (
	"context"

	"github.com/mattismoel/konnekt/internal/domain/artist"
	"github.com/mattismoel/konnekt/internal/object"
	"github.com/mattismoel/konnekt/internal/query"
)

type ArtistService struct {
	artistRepo  artist.Repository
	objectStore object.Store
}

func NewArtistService(artistRepo artist.Repository, objectStore object.Store) (*ArtistService, error) {
	return &ArtistService{
		artistRepo:  artistRepo,
		objectStore: objectStore,
	}, nil
}

type CreateArtist struct {
	Name        string
	Description string
	GenreIDs    []int64
	Socials     []string
}

type ArtistListQuery struct {
	query.ListQuery
}

type GenreListQuery struct {
	query.ListQuery
}

func (s ArtistService) ByID(ctx context.Context, artistID int64) (artist.Artist, error) {
	a, err := s.artistRepo.ByID(ctx, artistID)
	if err != nil {
		return artist.Artist{}, err
	}

	return a, nil
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

	a, err := artist.NewArtist(
		artist.WithName(load.Name),
		artist.WithDescription(load.Description),
		artist.WithGenres(genres...),
		artist.WithSocials(socials...),
	)

	if err != nil {
		return 0, err
	}

	artistID, err := s.artistRepo.Insert(ctx, *a)
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

func (s ArtistService) ListGenres(ctx context.Context, q GenreListQuery) (query.ListResult[artist.Genre], error) {
	genres, totalCount, err := s.artistRepo.ListGenres(ctx, q.Offset(), q.Limit)
	if err != nil {
		return query.ListResult[artist.Genre]{}, err
	}

	return query.ListResult[artist.Genre]{
		Page:       q.Page,
		PerPage:    q.PerPage,
		PageCount:  q.PageCount(totalCount),
		TotalCount: totalCount,
		Records:    genres,
	}, nil
}

func (s ArtistService) CreateGenre(ctx context.Context, name string) (int64, error) {
	genreID, err := s.artistRepo.InsertGenre(ctx, name)
	if err != nil {
		return 0, err
	}

	return genreID, nil
}
