package sqlite

import (
	"context"
	"database/sql"

	sq "github.com/Masterminds/squirrel"
	"github.com/mattismoel/konnekt/internal/domain/artist"
	"github.com/mattismoel/konnekt/internal/query"
)

type Genre struct {
	ID   int64
	Name string
}

type GenreQueryParams struct {
	QueryParams
}

func (repo ArtistRepository) InsertGenre(ctx context.Context, name string) (int64, error) {
	tx, err := repo.db.BeginTx(ctx, nil)
	if err != nil {
		return 0, err
	}

	defer tx.Rollback()

	genreID, err := insertGenre(ctx, tx, name)
	if err != nil {
		return 0, err
	}

	if err := tx.Commit(); err != nil {
		return 0, err
	}

	return genreID, nil
}

func (repo ArtistRepository) ListGenres(ctx context.Context, q artist.GenreQuery) (query.ListResult[artist.Genre], error) {
	tx, err := repo.db.BeginTx(ctx, nil)
	if err != nil {
		return query.ListResult[artist.Genre]{}, err
	}

	defer tx.Rollback()

	dbGenres, err := listGenres(ctx, tx, GenreQueryParams{
		QueryParams: QueryParams{
			Offset: q.Offset(), Limit: q.Limit,
		},
	})

	if err != nil {
		return query.ListResult[artist.Genre]{}, err
	}

	genres := make([]artist.Genre, 0)
	for _, dbGenre := range dbGenres {
		genres = append(genres, artist.Genre{
			ID:   dbGenre.ID,
			Name: dbGenre.Name,
		})
	}

	totalCount, err := count(ctx, tx, "genre")
	if err != nil {
		return query.ListResult[artist.Genre]{}, err
	}

	if err := tx.Commit(); err != nil {
		return query.ListResult[artist.Genre]{}, err
	}

	return query.ListResult[artist.Genre]{
		Page:       q.Page,
		PerPage:    q.PerPage,
		TotalCount: totalCount,
		PageCount:  q.PageCount(totalCount),
		Records:    genres,
	}, nil
}

func (repo ArtistRepository) GenreByID(ctx context.Context, genreID int64) (artist.Genre, error) {
	tx, err := repo.db.BeginTx(ctx, nil)
	if err != nil {
		return artist.Genre{}, err
	}

	defer tx.Rollback()

	dbGenre, err := genreByID(ctx, tx, genreID)
	if err != nil {
		return artist.Genre{}, err
	}

	if err := tx.Commit(); err != nil {
		return artist.Genre{}, err
	}

	return artist.Genre{
		ID:   dbGenre.ID,
		Name: dbGenre.Name,
	}, nil
}

var genreBuilder = sq.Select("id", "name").From("genre")

func scanGenre(s Scanner, dst *Genre) error {
	err := s.Scan(&dst.ID, &dst.Name)
	if err != nil {
		return err
	}

	return nil
}

// Lists genres based on the input {GenreQueryParams}.
func listGenres(ctx context.Context, tx *sql.Tx, params GenreQueryParams) ([]Genre, error) {
	builder := genreBuilder

	if params.Limit > 0 {
		builder = builder.Limit(uint64(params.Limit))
	}

	if params.Offset > 0 {
		builder = builder.Offset(uint64(params.Offset))
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

	genres := make([]Genre, 0)

	for rows.Next() {
		var g Genre
		if err := scanGenre(rows, &g); err != nil {
			return nil, err
		}

		genres = append(genres, g)
	}

	return genres, nil
}

// Inserts the genre with the given name.
//
// If the attemptedly inserted genre already exists, the previous genre is
// returned.
func insertGenre(ctx context.Context, tx *sql.Tx, name string) (int64, error) {
	existing, args, err := genreBuilder.
		Where(sq.Eq{"name": name}).
		ToSql()

	if err != nil {
		return 0, err
	}

	var existingGenre Genre
	row := tx.QueryRowContext(ctx, existing, args...)
	if err := scanGenre(row, &existingGenre); err == nil {
		return existingGenre.ID, err
	}

	query, args, err := sq.
		Insert("genre").
		Columns("name").
		Values(name).
		ToSql()

	if err != nil {
		return 0, err
	}

	res, err := tx.ExecContext(ctx, query, args...)
	if err != nil {
		return 0, err
	}

	genreID, err := res.LastInsertId()
	if err != nil {
		return 0, err
	}

	return genreID, nil
}

// Sets the genres of an artist to the specified input genres.
//
// All previous genres of the artist will be overridden with this function.
func setArtistGenres(ctx context.Context, tx *sql.Tx, artistID int64, genres ...Genre) error {
	err := deleteArtistGenres(ctx, tx, artistID)
	if err != nil {
		return err
	}

	for _, genre := range genres {
		genreID, err := insertGenre(ctx, tx, genre.Name)
		if err != nil {
			return err
		}

		err = associateArtistWithGenre(ctx, tx, artistID, genreID)
		if err != nil {
			return err
		}
	}

	return nil
}

// Deletes all artist-to-genre relationships, clearing all associated genres for
// the given artist.
//
// This does not delete the genre entity itself.
func deleteArtistGenres(ctx context.Context, tx *sql.Tx, artistID int64) error {
	genres, err := artistGenres(ctx, tx, artistID)
	if err != nil {
		return err
	}

	for _, genre := range genres {
		err = dissasociateArtistFromGenre(ctx, tx, artistID, genre.ID)
		if err != nil {
			return err
		}
	}

	return nil
}

// Associates the given genre with the given artist.
func associateArtistWithGenre(ctx context.Context, tx *sql.Tx, artistID int64, genreID int64) error {
	query, args, err := sq.
		Insert("artists_genres").
		Columns("artist_id", "genre_id").
		Values(artistID, genreID).
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

// Dissasociates the given genre from the given artist.
func dissasociateArtistFromGenre(ctx context.Context, tx *sql.Tx, artistID int64, genreID int64) error {
	query, args, err := sq.
		Delete("artists_genres").
		Where(sq.Eq{"artist_id": artistID}).
		Where(sq.Eq{"genre_id": genreID}).
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

// Lists all genres associated with an artist.
func artistGenres(ctx context.Context, tx *sql.Tx, artistID int64) ([]Genre, error) {
	query, args, err := genreBuilder.
		Join("artists_genres ag on ag.genre_id = genre.id").
		Where(sq.Eq{"ag.artist_id": artistID}).
		ToSql()

	if err != nil {
		return nil, err
	}

	rows, err := tx.QueryContext(ctx, query, args...)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	genres := make([]Genre, 0)

	for rows.Next() {
		var g Genre
		if err := scanGenre(rows, &g); err != nil {
			return nil, err
		}

		genres = append(genres, g)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return genres, nil
}

// Gets a genre by its ID.
func genreByID(ctx context.Context, tx *sql.Tx, genreID int64) (Genre, error) {
	query, args, err := genreBuilder.
		Where(sq.Eq{"id": genreID}).
		ToSql()

	if err != nil {
		return Genre{}, err
	}

	var g Genre
	row := tx.QueryRowContext(ctx, query, args...)
	if err := scanGenre(row, &g); err != nil {
		return Genre{}, err
	}

	return g, nil
}
