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
	"github.com/nfnt/resize"

	_ "image/jpeg"
	_ "image/png"
)

const ARTIST_IMAGE_WIDTH_PX = 2048

var (
	AllowedImageFiletypes = []string{".png", ".jpeg", ".jpg"}
)

var (
	ErrInvalidImageFiletype = fmt.Errorf("Image file must be of format %s", strings.Join(AllowedImageFiletypes, ", "))
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
		return artist.Artist{}, fmt.Errorf("Could not get artist %d: %v", artistID, err)
	}

	return a, nil
}

func (s ArtistService) List(ctx context.Context, q query.ListQuery) (query.ListResult[artist.Artist], error) {
	result, err := s.artistRepo.List(ctx, q)
	if err != nil {
		return query.ListResult[artist.Artist]{}, fmt.Errorf("Could not list artists: %v", err)
	}

	return result, nil
}

func (s ArtistService) Create(ctx context.Context, a artist.Artist) (int64, error) {
	artistID, err := s.artistRepo.Insert(ctx, a)
	if err != nil {
		return 0, fmt.Errorf("Could not insert artist: %v", err)
	}

	return artistID, nil
}

func (s ArtistService) Update(ctx context.Context, artistID int64, load UpdateArtist) error {
	prevArtist, err := s.ByID(ctx, artistID)
	if err != nil {
		return fmt.Errorf("Could not get artist %d: %v", artistID, err)
	}

	socials := make([]artist.Social, 0)
	for _, social := range load.Socials {
		s, err := artist.NewSocial(social)
		if err != nil {
			return fmt.Errorf("Could not create artist social: %v", err)
		}

		socials = append(socials, s)
	}

	genres := make([]artist.Genre, 0)
	for _, genreID := range load.GenreIDs {
		genre, err := s.artistRepo.GenreByID(ctx, genreID)
		if err != nil {
			return fmt.Errorf("Could not get genre %d: %v", genreID, err)
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
		return fmt.Errorf("Could not create artist: %v", err)
	}

	if strings.TrimSpace(load.PreviewURL) != "" {
		err := a.WithCfgs(artist.WithPreviewURL(load.PreviewURL))
		if err != nil {
			return fmt.Errorf("Could not use artist preview URL: %v", err)
		}
	}

	if strings.TrimSpace(load.ImageURL) != "" {
		// Delete previous artist image from object store.
		url, err := url.Parse(prevArtist.ImageURL)
		if err != nil {
			return fmt.Errorf("Could not parse artist image URL: %v", err)
		}

		if err := s.objectStore.Delete(ctx, url.Path); err != nil {
			return fmt.Errorf("Could not delete previous artist image: %v", err)
		}

		// Set the new artist image url.
		if err := a.WithCfgs(artist.WithImageURL(load.ImageURL)); err != nil {
			return fmt.Errorf("Could not use artist image URL: %v", err)
		}
	}

	err = s.artistRepo.Update(ctx, artistID, a)
	if err != nil {
		return fmt.Errorf("Could not update artist: %v", err)
	}

	return nil
}

func (s ArtistService) Delete(ctx context.Context, artistID int64) error {
	artistEventsResult, err := s.ArtistEvents(ctx, artistID)
	if err != nil {
		return fmt.Errorf("Could not get artist events: %v", err)
	}

	if len(artistEventsResult.Records) > 0 {
		return ErrArtistInEvent
	}

	a, err := s.artistRepo.ByID(ctx, artistID)
	if err != nil {
		return fmt.Errorf("Could not find artist %d: %v", artistID, err)
	}

	url, err := url.Parse(a.ImageURL)
	if err != nil {
		return fmt.Errorf("Could not parse artist image URL: %v", err)
	}

	err = s.objectStore.Delete(ctx, url.Path)
	if err != nil {
		return fmt.Errorf("Could not delete artist image: %v", err)
	}

	err = s.artistRepo.Delete(ctx, artistID)
	if err != nil {
		return fmt.Errorf("Could not delete artist from repository: %v", err)
	}

	return nil
}

func (s ArtistService) ListGenres(ctx context.Context, q query.ListQuery) (query.ListResult[artist.Genre], error) {
	result, err := s.artistRepo.ListGenres(ctx, q)
	if err != nil {
		return query.ListResult[artist.Genre]{}, fmt.Errorf("Could not list genres: %v", err)
	}

	return result, nil
}

func (s ArtistService) CreateGenre(ctx context.Context, name string) (int64, error) {
	genreID, err := s.artistRepo.InsertGenre(ctx, name)
	if err != nil {
		return 0, fmt.Errorf("Could not insert gerne into repository: %v", err)
	}

	return genreID, nil
}

func (s ArtistService) UploadImage(ctx context.Context, r io.Reader) (string, error) {
	img, _, err := image.Decode(r)
	if err != nil {
		return "", fmt.Errorf("Could not decode artist image: %v", err)
	}

	fileName := createRandomImageFileName("jpeg")

	// Resize if too high resolution.
	if img.Bounds().Max.X > ARTIST_IMAGE_WIDTH_PX {
		img = resize.Resize(ARTIST_IMAGE_WIDTH_PX, 0, img, resize.Lanczos2)
	}

	formattedImg, err := formatJPEG(img)
	if err != nil {
		return "", fmt.Errorf("Could not format artist image: %v", err)
	}

	url, err := s.objectStore.Upload(ctx, path.Join("/artists", fileName), formattedImg)
	if err != nil {
		return "", fmt.Errorf("Could not upload artist image: %v", err)
	}

	return url, nil
}

func (s ArtistService) ArtistEvents(ctx context.Context, artistID int64) (query.ListResult[event.Event], error) {
	artistFilter, err := query.NewFilter(query.Equal, strconv.Itoa(int(artistID)))
	if err != nil {
		return query.ListResult[event.Event]{}, fmt.Errorf("Could not create artist filter: %v", err)
	}

	q, err := query.NewListQuery(query.WithFilters(query.FilterCollection{
		"artist_id": []query.Filter{artistFilter},
	}))

	result, err := s.eventRepo.List(ctx, q)
	if err != nil {
		return query.ListResult[event.Event]{}, fmt.Errorf("Could not list artist events: %v", err)
	}

	return result, nil
}

func (s ArtistService) GenreByID(ctx context.Context, genreID int64) (artist.Genre, error) {
	genre, err := s.artistRepo.GenreByID(ctx, genreID)
	if err != nil {
		return artist.Genre{}, err
	}

	return genre, nil
}
