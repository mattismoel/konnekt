package service

import (
	"context"
	"errors"
	"fmt"
	"image"
	"io"
	"net/url"
	"path"
	"strconv"
	"strings"

	"github.com/mattismoel/konnekt/internal/domain/artist"
	"github.com/mattismoel/konnekt/internal/domain/event"
	"github.com/mattismoel/konnekt/internal/object"
	"github.com/mattismoel/konnekt/internal/query"
)

const ARTIST_IMAGE_WIDTH_PX = 2048

var (
	AllowedImageFiletypes = []string{".png", ".jpeg", ".jpg"}
)

var (
	ErrInvalidImageFiletype = errors.New(fmt.Sprintf("Image file must be of format %s", strings.Join(AllowedImageFiletypes, ", ")))
	ErrArtistInEvent        = errors.New("Artist must not be part of an event to be deleted")
)

type ArtistService struct {
	artistRepo  artist.Repository
	eventRepo   event.Repository
	objectStore object.Store
}

func NewArtistService(artistRepo artist.Repository, eventRepo event.Repository, objectStore object.Store) (*ArtistService, error) {
	return &ArtistService{
		artistRepo:  artistRepo,
		eventRepo:   eventRepo,
		objectStore: objectStore,
	}, nil
}

type CreateArtist struct {
	Name        string
	Description string
	ImageURL    string
	PreviewURL  string
	GenreIDs    []int64
	Socials     []string
}

type UpdateArtist struct {
	Name        string
	Description string
	ImageURL    string
	PreviewURL  string
	GenreIDs    []int64
	Socials     []string
}

func (s ArtistService) ByID(ctx context.Context, artistID int64) (artist.Artist, error) {
	a, err := s.artistRepo.ByID(ctx, artistID)
	if err != nil {
		return artist.Artist{}, err
	}

	return a, nil
}

func (s ArtistService) List(ctx context.Context, q artist.Query) (query.ListResult[artist.Artist], error) {
	result, err := s.artistRepo.List(ctx, q)
	if err != nil {
		return query.ListResult[artist.Artist]{}, err
	}

	return result, nil
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
		artist.WithImageURL(load.ImageURL),
		artist.WithGenres(genres...),
		artist.WithSocials(socials...),
	)

	if err != nil {
		return 0, err
	}

	if strings.TrimSpace(load.PreviewURL) != "" {
		err := a.WithCfgs(artist.WithPreviewURL(load.PreviewURL))
		if err != nil {
			return 0, err
		}
	}

	artistID, err := s.artistRepo.Insert(ctx, *a)
	if err != nil {
		return 0, err
	}

	return artistID, nil
}

func (s ArtistService) Update(ctx context.Context, artistID int64, load UpdateArtist) (artist.Artist, error) {
	prevArtist, err := s.ByID(ctx, artistID)
	if err != nil {
		return artist.Artist{}, err
	}

	socials := make([]artist.Social, 0)
	for _, social := range load.Socials {
		s, err := artist.NewSocial(social)
		if err != nil {
			return artist.Artist{}, err
		}

		socials = append(socials, s)
	}

	genres := make([]artist.Genre, 0)
	for _, genreID := range load.GenreIDs {
		genre, err := s.artistRepo.GenreByID(ctx, genreID)
		if err != nil {
			return artist.Artist{}, err
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
		return artist.Artist{}, err
	}

	if strings.TrimSpace(load.PreviewURL) != "" {
		err := a.WithCfgs(artist.WithPreviewURL(load.PreviewURL))
		if err != nil {
			return artist.Artist{}, err
		}
	}

	if strings.TrimSpace(load.ImageURL) != "" {
		// Delete previous artist image from object store.
		url, err := url.Parse(prevArtist.ImageURL)
		if err != nil {
			return artist.Artist{}, err
		}

		if err := s.objectStore.Delete(ctx, url.Path); err != nil {
			return artist.Artist{}, err
		}

		// Set the new artist image url.
		if err := a.WithCfgs(artist.WithImageURL(load.ImageURL)); err != nil {
			return artist.Artist{}, err
		}
	}

	err = s.artistRepo.Update(ctx, artistID, *a)
	if err != nil {
		return artist.Artist{}, nil
	}

	return *a, nil
}

func (s ArtistService) Delete(ctx context.Context, artistID int64) error {
	artistEventsResult, err := s.ArtistEvents(ctx, artistID)
	if err != nil {
		return err
	}

	if len(artistEventsResult.Records) > 0 {
		return ErrArtistInEvent
	}

	a, err := s.artistRepo.ByID(ctx, artistID)
	if err != nil {
		return err
	}

	url, err := url.Parse(a.ImageURL)
	if err != nil {
		return err
	}

	err = s.objectStore.Delete(ctx, url.Path)
	if err != nil {
		return err
	}

	err = s.artistRepo.Delete(ctx, artistID)
	if err != nil {
		return err
	}

	return nil
}

func (s ArtistService) ListGenres(ctx context.Context, q artist.GenreQuery) (query.ListResult[artist.Genre], error) {
	result, err := s.artistRepo.ListGenres(ctx, q)
	if err != nil {
		return query.ListResult[artist.Genre]{}, err
	}

	return result, nil
}

func (s ArtistService) CreateGenre(ctx context.Context, name string) (int64, error) {
	genreID, err := s.artistRepo.InsertGenre(ctx, name)
	if err != nil {
		return 0, err
	}

	return genreID, nil
}

func (s ArtistService) UploadImage(ctx context.Context, r io.Reader) (string, error) {
	img, _, err := image.Decode(r)
	if err != nil {
		return "", err
	}

	fileName := createRandomImageFileName("jpeg")

	// Resize if too high resolution.
	if img.Bounds().Max.X > ARTIST_IMAGE_WIDTH_PX {
		resizedImage, err := resizeImage(img, ARTIST_IMAGE_WIDTH_PX, 0)
		if err != nil {
			return "", err
		}

		url, err := s.objectStore.Upload(ctx, path.Join("/artists", fileName), resizedImage)
		if err != nil {
			return "", err
		}

		return url, nil
	}

	url, err := s.objectStore.Upload(ctx, path.Join("/artists", fileName), r)
	if err != nil {
		return "", err
	}

	return url, nil
}

func (s ArtistService) ArtistEvents(ctx context.Context, artistID int64) (query.ListResult[event.Event], error) {
	artistFilter, err := query.NewFilter(query.Equal, strconv.Itoa(int(artistID)))
	if err != nil {
		return query.ListResult[event.Event]{}, err
	}

	q, err := query.NewListQuery(query.WithFilters(query.FilterCollection{
		"artist_id": []query.Filter{artistFilter},
	}))

	result, err := s.eventRepo.List(ctx, q)
	if err != nil {
		return query.ListResult[event.Event]{}, err
	}

	return result, nil
}
