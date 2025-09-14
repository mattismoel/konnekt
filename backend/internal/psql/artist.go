package psql

import (
	"context"
	"errors"
	"fmt"

	sq "github.com/Masterminds/squirrel"
	"github.com/jackc/pgx/v5"
	konnekt "github.com/mattismoel/konnekt/backend"
	"github.com/mattismoel/konnekt/backend/api"
	"github.com/mattismoel/konnekt/backend/mask"
	"github.com/mattismoel/konnekt/backend/order"
)

var artistSelectBuilder = psql.
	Select(
		"id",
		"name",
		"description",
		"image_url",
		"preview_url",
		"created_by",
		"updated_by",
		"created_at",
		"updated_at",
	).
	From("artist")

var _ konnekt.ArtistRepo = ArtistRepo{}

type Artist struct {
	ID          int64
	Name        string
	Description string
	ImageURL    string
	PreviewURL  string

	AuditFields
	Timestamps
}

type ArtistRepo struct {
	Conn *pgx.Conn
}

// ArtistByID implements konnekt.ArtistRepo.
func (a ArtistRepo) ArtistByID(ctx context.Context, artistID konnekt.ID) (konnekt.Artist, error) {
	var artist konnekt.Artist
	err := pgx.BeginFunc(ctx, a.Conn, func(tx pgx.Tx) error {
		dbArtist, err := artistByID(ctx, tx, int64(artistID))
		if err != nil {
			if errors.Is(err, pgx.ErrNoRows) {
				return konnekt.ErrResourceNotFound
			}
			return fmt.Errorf("Could not get artist by ID: %v", err)
		}

		dbGenres, err := artistGenres(ctx, tx, int64(artistID))
		if err != nil {
			if errors.Is(err, pgx.ErrNoRows) {
				return konnekt.ErrResourceNotFound
			}

			return fmt.Errorf("Could not get artist genres: %v", err)
		}

		dbSocials, err := artistSocials(ctx, tx, int64(artistID))
		if err != nil {
			if errors.Is(err, pgx.ErrNoRows) {
				return konnekt.ErrResourceNotFound
			}

			return fmt.Errorf("Could not get artist socials: %v", err)
		}

		fmt.Println("SOCIALS", dbSocials)
		artist = dbArtist.ToDomain()
		artist.Genres = dbGenres.ToDomain()
		artist.Socials = dbSocials.ToDomain()

		return nil
	})

	if err != nil {
		return konnekt.Artist{}, err
	}

	return artist, nil
}

// DeleteArtist implements konnekt.ArtistRepo.
func (a ArtistRepo) DeleteArtist(ctx context.Context, artistID konnekt.ID) error {
	err := pgx.BeginFunc(ctx, a.Conn, func(tx pgx.Tx) error {
		if err := deleteArtistSocials(ctx, tx, int64(artistID)); err != nil {
			return err
		}

		if err := clearArtistGenres(ctx, tx, int64(artistID)); err != nil {
			return err
		}

		if err := deleteArtist(ctx, tx, int64(artistID)); err != nil {
			return err
		}

		return nil
	})

	if err != nil {
		return err
	}

	return nil
}

// InsertArtist implements konnekt.ArtistRepo.
func (a ArtistRepo) InsertArtist(ctx context.Context, ca konnekt.CreateArtist) (konnekt.ID, error) {
	var insertedID konnekt.ID
	err := pgx.BeginFunc(ctx, a.Conn, func(tx pgx.Tx) error {
		artistID, err := insertArtist(ctx, tx, Artist{
			Name:        ca.Name,
			Description: ca.Description,
			ImageURL:    ca.ImageURL,
			PreviewURL:  ca.PreviewURL,
			AuditFields: AuditFields{
				CreatedBy: ca.CreatedBy,
				UpdatedBy: ca.CreatedBy,
			},
		})

		// Ensure that artist genres and socials are cleared, in case of conflict update.
		if err := clearArtistGenres(ctx, tx, artistID); err != nil {
			return err
		}

		if err := deleteArtistSocials(ctx, tx, artistID); err != nil {
			return err
		}

		for _, genreID := range ca.GenreIDs {
			if err := associateArtistWithGenre(ctx, tx, artistID, int64(genreID)); err != nil {
				return fmt.Errorf("Could not associate artist with genre: %v", err)
			}
		}

		for _, url := range ca.Socials {
			_, err := insertSocial(ctx, tx, Social{URL: string(url), ArtistID: artistID})
			if err != nil {
				return fmt.Errorf("Could not insert artist social: %v", err)
			}
		}

		if err != nil {
			return fmt.Errorf("Could not insert artist: %v", err)
		}

		insertedID = konnekt.ID(artistID)
		return nil
	})

	if err != nil {
		return 0, err
	}

	return insertedID, nil
}

// ListArtists implements konnekt.ArtistRepo.
func (a ArtistRepo) ListArtists(ctx context.Context, lr api.ListRequest) (api.ListResponse[konnekt.Artist], error) {
	artists := make([]konnekt.Artist, 0)
	err := pgx.BeginFunc(ctx, a.Conn, func(tx pgx.Tx) error {
		pagination := paginationFromListRequest(lr)
		dbArtists, err := listArtists(ctx, tx, pagination, lr.OrderMap)
		if err != nil {
			return err
		}

		artists = dbArtists.ToDomain()

		for i := range artists {
			artistGenres, err := artistGenres(ctx, tx, int64(artists[i].ID))
			if err != nil {
				return err
			}

			artistSocials, err := artistSocials(ctx, tx, int64(artists[i].ID))
			if err != nil {
				return err
			}

			artists[i].Socials = artistSocials.ToDomain()
			artists[i].Genres = artistGenres.ToDomain()
		}

		return nil
	})

	if err != nil {
		return api.ListResponse[konnekt.Artist]{}, err
	}

	return api.ListResponse[konnekt.Artist]{
		// TotalSize: totalSize,
		Records: artists,
	}, nil
}

// UpdateArtist implements konnekt.ArtistRepo.
func (a ArtistRepo) UpdateArtist(ctx context.Context, artistID konnekt.ID, ur api.UpdateRequest[konnekt.UpdateArtist]) error {
	err := pgx.BeginFunc(ctx, a.Conn, func(tx pgx.Tx) error {
		updateMap := ur.UpdateMap()

		if err := updateArtist(ctx, tx, int64(artistID), updateMap); err != nil {
			if errors.Is(err, pgx.ErrNoRows) {
				return konnekt.ErrResourceNotFound
			}

			return fmt.Errorf("Could not update artist: %v", err)
		}

		if _, ok := updateMap["socials"]; ok {
			if err := deleteArtistSocials(ctx, tx, int64(artistID)); err != nil {
				return fmt.Errorf("Could not delete artist socials: %v", err)
			}

			for _, s := range ur.Data.Socials {
				_, err := insertSocial(ctx, tx, Social{
					URL:      string(s),
					ArtistID: int64(artistID),
				})

				if err != nil {
					return fmt.Errorf("Could not insert social: %v", err)
				}
			}
		}

		if _, ok := updateMap["genres"]; ok {
			if err := clearArtistGenres(ctx, tx, int64(artistID)); err != nil {
				return fmt.Errorf("Could not delete artist genres: %v", err)
			}

			for _, genreID := range ur.Data.GenreIDs {
				err := associateArtistWithGenre(ctx, tx, int64(artistID), int64(genreID))
				if err != nil {
					return fmt.Errorf("Could not associate artist %d with genre %d: %v", artistID, genreID, err)
				}
			}
		}
		return nil
	})

	if err != nil {
		return err
	}

	return nil
}

func insertArtist(ctx context.Context, tx pgx.Tx, a Artist) (int64, error) {
	existingArtist, err := artistByName(ctx, tx, a.Name)
	if err == nil {
		return existingArtist.ID, nil
	}

	if !errors.Is(err, pgx.ErrNoRows) {
		return 0, fmt.Errorf("Could not check for existing artist: %v", err)
	}

	query, args, err := psql.
		Insert("artist").
		Columns("name", "description", "image_url", "preview_url", "created_by", "updated_by").
		Values(a.Name, a.Description, a.ImageURL, a.PreviewURL, a.CreatedBy, a.CreatedBy).
		Suffix("ON CONFLICT (name) DO UPDATE SET name = EXCLUDED.name").
		Suffix("RETURNING id").
		ToSql()

	if err != nil {
		return 0, NewQueryBuildError("insert artist", err)
	}

	var id int64
	if err := tx.QueryRow(ctx, query, args...).Scan(&id); err != nil {
		return 0, fmt.Errorf("Could not insert artist: %v", err)
	}

	return id, nil
}

func updateArtist(ctx context.Context, tx pgx.Tx, artistID int64, fieldMap mask.FieldMap) error {
	allowed := []string{"name", "description", "image_url", "preview_url"}

	updates := make(map[string]any)
	for _, key := range allowed {
		if v, ok := fieldMap[mask.FieldName(key)]; ok {
			updates[key] = v
		}
	}

	if len(updates) == 0 {
		return nil
	}

	query, args, err := psql.
		Update("artist").
		Where(sq.Eq{"id": artistID}).
		SetMap(updates).
		ToSql()

	if err != nil {
		return NewQueryBuildError("update artist", err)
	}

	cmdTag, err := tx.Exec(ctx, query, args...)
	if err != nil {
		return fmt.Errorf("Could not update artist")
	}

	if cmdTag.RowsAffected() == 0 {
		return pgx.ErrNoRows
	}

	return nil
}

func listArtists(ctx context.Context, tx pgx.Tx, pg Pagination, orderMap order.Map) (Collection[Artist, konnekt.Artist], error) {
	builder := artistSelectBuilder
	builder = applyPagination(builder, pg)
	builder = applyOrdering(builder, orderMap)

	query, args, err := builder.ToSql()
	if err != nil {
		return nil, NewQueryBuildError("list artists", err)
	}

	rows, err := tx.Query(ctx, query, args...)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	artists := make([]Artist, 0)
	for rows.Next() {
		artist, err := scanArtist(rows)
		if err != nil {
			return nil, err
		}
		artists = append(artists, artist)
	}

	return artists, nil
}

func artistByID(ctx context.Context, tx pgx.Tx, artistID int64) (Artist, error) {
	query, args, err := artistSelectBuilder.Where(sq.Eq{"id": artistID}).ToSql()
	if err != nil {
		return Artist{}, NewQueryBuildError("artist by id", err)
	}

	row := tx.QueryRow(ctx, query, args...)
	artist, err := scanArtist(row)
	if err != nil {
		return Artist{}, err
	}

	return artist, nil
}

func artistByName(ctx context.Context, tx pgx.Tx, artistName string) (Artist, error) {
	query, args, err := artistSelectBuilder.Where(sq.Eq{"name": artistName}).ToSql()
	if err != nil {
		return Artist{}, NewQueryBuildError("artist by name", err)
	}

	row := tx.QueryRow(ctx, query, args...)
	artist, err := scanArtist(row)
	if err != nil {
		return Artist{}, err
	}

	return artist, nil
}

func deleteArtist(ctx context.Context, tx pgx.Tx, artistID int64) error {
	query, args, err := psql.Delete("artist").Where(sq.Eq{"id": artistID}).ToSql()
	if err != nil {
		return NewQueryBuildError("delete artist", err)
	}

	if _, err := tx.Exec(ctx, query, args...); err != nil {
		return err
	}

	return nil
}

func scanArtist(s Scanner) (Artist, error) {
	var a Artist
	err := scan(s, &a.ID, &a.Name, &a.Description, &a.ImageURL, &a.PreviewURL,
		&a.CreatedBy, &a.UpdatedBy, &a.CreatedAt, &a.UpdatedAt)

	if err != nil {
		return Artist{}, err
	}

	return a, nil
}

func (a Artist) ToDomain() konnekt.Artist {
	return konnekt.Artist{
		ID:          konnekt.ID(a.ID),
		Name:        a.Name,
		Description: a.Description,
		ImageURL:    a.ImageURL,
		PreviewURL:  a.PreviewURL,
		Socials:     make([]konnekt.Social, 0),
		Genres:      make([]konnekt.Genre, 0),
	}
}

func (a Artist) Fields() mask.FieldMap {
	return mask.FieldMap{
		"id":          a.ID,
		"name":        a.Description,
		"description": a.Description,
		"image_url":   a.ImageURL,
		"preview_url": a.PreviewURL,
	}
}
