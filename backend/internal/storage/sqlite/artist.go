package sqlite

import (
	"context"
	"database/sql"

	"github.com/mattismoel/konnekt/internal/domain/artist"
	"github.com/mattismoel/konnekt/internal/query"
)

type Artist struct {
	ID          int64
	Name        string
	Description string
	PreviewURL  string
	ImageURL    string
}

func (a Artist) ToInternal(genres []artist.Genre, socials []artist.Social) artist.Artist {
	return artist.Artist{
		ID:          a.ID,
		Name:        a.Name,
		Description: a.Description,
		PreviewURL:  a.PreviewURL,
		ImageURL:    a.ImageURL,
		Genres:      genres,
		Socials:     socials,
	}
}

var _ artist.Repository = (*ArtistRepository)(nil)

type ArtistRepository struct {
	db *sql.DB
}

func NewArtistRepository(db *sql.DB) (*ArtistRepository, error) {
	return &ArtistRepository{
		db: db,
	}, nil
}

func (repo ArtistRepository) List(ctx context.Context, q artist.Query) (query.ListResult[artist.Artist], error) {
	tx, err := repo.db.BeginTx(ctx, nil)
	if err != nil {
		return query.ListResult[artist.Artist]{}, err
	}

	defer tx.Rollback()

	artists := make([]artist.Artist, 0)

	dbArtists, err := listArtists(ctx, tx, QueryParams{
		Offset:  q.Offset(),
		Limit:   q.Limit,
		OrderBy: q.OrderBy,
		Filters: q.Filters,
	})
	if err != nil {
		return query.ListResult[artist.Artist]{}, err
	}

	for _, dbArtist := range dbArtists {
		dbGenres, err := artistGenres(ctx, tx, dbArtist.ID)
		if err != nil {
			return query.ListResult[artist.Artist]{}, err
		}

		genres := make([]artist.Genre, 0)
		for _, dbGenre := range dbGenres {
			genres = append(genres, artist.Genre{
				ID:   dbGenre.ID,
				Name: dbGenre.Name,
			})
		}

		dbSocials, err := artistSocials(ctx, tx, dbArtist.ID)
		if err != nil {
			return query.ListResult[artist.Artist]{}, err
		}

		socials := make([]artist.Social, 0)
		for _, dbSocial := range dbSocials {
			socials = append(socials, artist.Social(dbSocial.URL))
		}

		artists = append(artists, dbArtist.ToInternal(genres, socials))
	}

	totalCount, err := artistCount(ctx, tx)
	if err != nil {
		return query.ListResult[artist.Artist]{}, err
	}

	if err := tx.Commit(); err != nil {
		return query.ListResult[artist.Artist]{}, err
	}

	return query.ListResult[artist.Artist]{
		Page:       q.Page,
		PerPage:    q.PerPage,
		TotalCount: totalCount,
		PageCount:  q.PageCount(totalCount),
		Records:    artists,
	}, nil
}

func (repo ArtistRepository) Insert(ctx context.Context, a artist.Artist) (int64, error) {
	tx, err := repo.db.BeginTx(ctx, nil)
	if err != nil {
		return 0, err
	}

	artistID, err := insertArtist(ctx, tx, Artist{
		Name:        a.Name,
		Description: a.Description,
		PreviewURL:  a.PreviewURL,
		ImageURL:    a.ImageURL,
	})

	if err != nil {
		return 0, err
	}

	for _, genre := range a.Genres {
		genreID, err := insertGenre(ctx, tx, genre.Name)
		if err != nil {
			return 0, err
		}

		err = associateArtistWithGenre(ctx, tx, artistID, genreID)
		if err != nil {
			return 0, err
		}
	}

	for _, social := range a.Socials {
		socialID, err := insertSocial(ctx, tx, string(social))
		if err != nil {
			return 0, err
		}

		err = associateArtistWithSocial(ctx, tx, artistID, socialID)
		if err != nil {
			return 0, err
		}
	}

	defer tx.Rollback()

	if err := tx.Commit(); err != nil {
		return 0, err
	}

	return artistID, nil
}

func (repo ArtistRepository) Update(ctx context.Context, artistID int64, a artist.Artist) error {
	tx, err := repo.db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}

	defer tx.Rollback()

	err = updateArtist(ctx, tx, artistID, Artist{
		Name:        a.Name,
		Description: a.Description,
		PreviewURL:  a.PreviewURL,
		ImageURL:    a.ImageURL,
	})

	if err != nil {
		return err
	}

	socials := make([]Social, 0)
	for _, social := range a.Socials {
		socials = append(socials, Social{URL: string(social)})
	}

	err = setArtistSocials(ctx, tx, artistID, socials...)
	if err != nil {
		return err
	}

	genres := make([]Genre, 0)
	for _, genre := range a.Genres {
		genres = append(genres, Genre{ID: genre.ID, Name: genre.Name})
	}

	err = setArtistGenres(ctx, tx, artistID, genres...)
	if err != nil {
		return err
	}

	if err := tx.Commit(); err != nil {
		return err
	}

	return nil
}

func (repo ArtistRepository) SetImageURL(ctx context.Context, artistID int64, url string) error {
	tx, err := repo.db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}

	defer tx.Rollback()

	err = setArtistImageURL(ctx, tx, artistID, url)
	if err != nil {
		return err
	}

	if err := tx.Commit(); err != nil {
		return err
	}

	return nil
}

func (repo ArtistRepository) ByID(ctx context.Context, artistID int64) (artist.Artist, error) {
	tx, err := repo.db.BeginTx(ctx, nil)
	if err != nil {
		return artist.Artist{}, err
	}

	defer tx.Rollback()

	dbArtist, err := artistByID(ctx, tx, artistID)
	if err != nil {
		return artist.Artist{}, err
	}

	dbGenres, err := artistGenres(ctx, tx, artistID)
	if err != nil {
		return artist.Artist{}, err
	}

	dbSocials, err := artistSocials(ctx, tx, artistID)
	if err != nil {
		return artist.Artist{}, err
	}

	if err := tx.Commit(); err != nil {
		return artist.Artist{}, err
	}

	genres := make([]artist.Genre, 0)
	for _, dbGenre := range dbGenres {
		genres = append(genres, artist.Genre{
			ID:   dbGenre.ID,
			Name: dbGenre.Name,
		})
	}

	socials := make([]artist.Social, 0)
	for _, s := range dbSocials {
		socials = append(socials, artist.Social(s.URL))
	}

	return dbArtist.ToInternal(genres, socials), nil
}

func (repo ArtistRepository) Delete(ctx context.Context, artistID int64) error {
	tx, err := repo.db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}

	defer tx.Rollback()

	artist, err := artistByID(ctx, tx, artistID)
	if err != nil {
		return err
	}

	socials, err := artistSocials(ctx, tx, artist.ID)
	if err != nil {
		return err
	}

	for _, social := range socials {
		err := deleteSocial(ctx, tx, social.ID)
		if err != nil {
			return err
		}
	}

	err = deleteArtist(ctx, tx, artist.ID)
	if err != nil {
		return err
	}

	if err := tx.Commit(); err != nil {
		return err
	}

	return nil
}

func listArtists(ctx context.Context, tx *sql.Tx, params QueryParams) ([]Artist, error) {
	q, err := NewQuery(`SELECT id, name, description, preview_url, image_url FROM artist a`)
	if err != nil {
		return nil, err
	}

	if err := q.WithOrdering(params.OrderBy); err != nil {
		return nil, err
	}

	if err := q.WithOffset(params.Offset); err != nil {
		return nil, err
	}

	if err := q.WithLimit(params.Limit); err != nil {
		return nil, err
	}

	queryStr, args := q.Build()

	rows, err := tx.QueryContext(ctx, queryStr, args...)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	artists := make([]Artist, 0)

	for rows.Next() {
		var id int64
		var name, description, previewUrl, imageURL string

		if err := rows.Scan(&id, &name, &description, &previewUrl, &imageURL); err != nil {
			return nil, err
		}

		artists = append(artists, Artist{
			ID:          id,
			Name:        name,
			Description: description,
			PreviewURL:  previewUrl,
			ImageURL:    imageURL,
		})
	}

	return artists, nil
}

func artistCount(ctx context.Context, tx *sql.Tx) (int, error) {
	var count int

	err := tx.QueryRowContext(ctx, "SELECT COUNT(*) FROM artist").Scan(&count)
	if err != nil {
		return 0, err
	}

	return count, nil

}

func insertArtist(ctx context.Context, tx *sql.Tx, a Artist) (int64, error) {
	query := `
	INSERT INTO artist (name, description, preview_url, image_url)
	VALUES (@name, @description, @preview_url, @image_url)`

	res, err := tx.ExecContext(ctx, query,
		sql.Named("name", a.Name),
		sql.Named("description", a.Description),
		sql.Named("preview_url", a.PreviewURL),
		sql.Named("image_url", a.ImageURL),
	)

	if err != nil {
		return 0, err
	}

	artistID, err := res.LastInsertId()
	if err != nil {
		return 0, err
	}

	return artistID, nil
}

func artistByID(ctx context.Context, tx *sql.Tx, artistID int64) (Artist, error) {
	query := `
	SELECT name, description, preview_url, image_url
	FROM artist where id = @id`

	var name, description, previewURL, imageURL string
	err := tx.QueryRowContext(ctx, query, sql.Named("id", artistID)).Scan(
		&name, &description, &previewURL, &imageURL,
	)

	if err != nil {
		return Artist{}, err
	}

	return Artist{
		ID:          artistID,
		Name:        name,
		Description: description,
		PreviewURL:  previewURL,
		ImageURL:    imageURL,
	}, nil
}

func deleteArtist(ctx context.Context, tx *sql.Tx, artistID int64) error {
	query := `DELETE FROM artist WHERE id = @artist_id`

	_, err := tx.ExecContext(ctx, query, sql.Named("artist_id", artistID))
	if err != nil {
		return err
	}

	query = `DELETE FROM artists_socials WHERE artist_id = @artist_id`
	_, err = tx.ExecContext(ctx, query, sql.Named("artist_id", artistID))
	if err != nil {
		return err
	}

	query = `DELETE FROM artists_genres WHERE artist_id = @artist_id`
	_, err = tx.ExecContext(ctx, query, sql.Named("artist_id", artistID))
	if err != nil {
		return err
	}

	return nil
}

func setArtistImageURL(ctx context.Context, tx *sql.Tx, artistID int64, url string) error {
	query := `UPDATE artist SET image_url = @image_url WHERE id = @artist_id`

	_, err := tx.ExecContext(ctx, query,
		sql.Named("image_url", url),
		sql.Named("artist_id", artistID),
	)

	if err != nil {
		return err
	}

	return nil
}

func updateArtist(ctx context.Context, tx *sql.Tx, artistID int64, a Artist) error {
	query := `
	UPDATE artist SET 
		name = CASE 
			WHEN @name = '' THEN name
			ELSE @name
		END,
		description = CASE
			WHEN @description = '' THEN description
			ELSE @description
		END,
		preview_url = CASE
			WHEN @preview_url = '' THEN preview_url
			ELSE @preview_url
		END,
		image_url = CASE
			WHEN @image_url = '' THEN image_url
			ELSE @image_url
		END
	WHERE id = @artist_id`

	res, err := tx.ExecContext(ctx, query,
		sql.Named("name", a.Name),
		sql.Named("description", a.Description),
		sql.Named("preview_url", a.PreviewURL),
		sql.Named("image_url", a.ImageURL),
		sql.Named("artist_id", artistID),
	)

	if err != nil {
		return err
	}

	affected, err := res.RowsAffected()
	if err != nil {
		return err
	}

	if affected <= 0 {
		return ErrNotFound
	}

	return nil
}
