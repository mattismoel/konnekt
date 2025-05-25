package sqlite

import (
	"context"
	"database/sql"

	sq "github.com/Masterminds/squirrel"
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
	builder := sq.
		Select("id", "name", "description", "preview_url", "image_url").
		From("artist")

	if order, ok := params.OrderBy["name"]; ok {
		builder = builder.OrderBy("name " + string(order))
	}

	if params.Offset > 0 {
		builder = builder.Offset(uint64(params.Offset))
	}

	if params.Limit > 0 {
		builder = builder.Limit(uint64(params.Limit))
	}

	if filters, ok := params.Filters["artist_id"]; ok {
		for _, f := range filters {
			builder = builder.Where(sq.Eq{"artist_id": f.Value})
		}
	}

	query, args, err := builder.ToSql()
	if err != nil {
		return nil, err
	}

	rows, err := tx.QueryContext(ctx, query, args...)
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

	query, args, err := sq.Select("COUNT(*)").From("artist").ToSql()
	if err != nil {
		return 0, err
	}

	err = tx.QueryRowContext(ctx, query, args...).Scan(&count)
	if err != nil {
		return 0, err
	}

	return count, nil
}

func insertArtist(ctx context.Context, tx *sql.Tx, a Artist) (int64, error) {
	query, args, err := sq.
		Insert("artist").
		Columns("name", "description", "preview_url", "image_url").
		Values(a.Name, a.Description, a.PreviewURL, a.ImageURL).
		ToSql()

	res, err := tx.ExecContext(ctx, query, args...)
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
	query, args, err := sq.Select("name", "description", "preview_url", "image_url").
		From("artist").
		Where(sq.Eq{"id": artistID}).
		ToSql()

	if err != nil {
		return Artist{}, err
	}

	var name, description, previewURL, imageURL string

	err = tx.
		QueryRowContext(ctx, query, args...).
		Scan(&name, &description, &previewURL, &imageURL)

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
	artist, args, err := sq.Delete("artist").Where(sq.Eq{"id": artistID}).ToSql()
	if err != nil {
		return err
	}

	_, err = tx.ExecContext(ctx, artist, args...)
	if err != nil {
		return err
	}

	socials, args, err := sq.Delete("artists_socials").Where(sq.Eq{"artist_id": artistID}).ToSql()
	if err != nil {
		return err
	}

	_, err = tx.ExecContext(ctx, socials, args...)
	if err != nil {
		return err
	}

	genres, args, err := sq.Delete("artists_genres").Where(sq.Eq{"artist_id": artistID}).ToSql()
	_, err = tx.ExecContext(ctx, genres, args...)
	if err != nil {
		return err
	}

	return nil
}

func setArtistImageURL(ctx context.Context, tx *sql.Tx, artistID int64, url string) error {
	query, args, err := sq.
		Update("artist").
		Where(sq.Eq{"id": artistID}).
		Set("image_url", url).
		ToSql()

	if err != nil {
		return err
	}

	_, err = tx.ExecContext(ctx, query, args...)
	if err != nil {
		return err
	}

	return nil
}

func updateArtist(ctx context.Context, tx *sql.Tx, artistID int64, a Artist) error {
	builder := sq.Update("artist").Where(sq.Eq{"id": artistID})

	if a.Name != "" {
		builder = builder.Set("name", a.Name)
	}

	if a.Description != "" {
		builder = builder.Set("description", a.Description)
	}

	if a.PreviewURL != "" {
		builder = builder.Set("preview_url", a.PreviewURL)
	}

	if a.ImageURL != "" {
		builder = builder.Set("image_url", a.ImageURL)
	}

	query, args, err := builder.ToSql()
	if err != nil {
		return err
	}

	res, err := tx.ExecContext(ctx, query, args...)
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
