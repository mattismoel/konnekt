package sqlite

import (
	"context"
	"database/sql"

	"github.com/mattismoel/konnekt/internal/domain/artist"
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

func (repo ArtistRepository) ListGenres(ctx context.Context, limit, offset int) ([]artist.Genre, int, error) {
	tx, err := repo.db.BeginTx(ctx, nil)
	if err != nil {
		return nil, 0, err
	}

	defer tx.Rollback()

	dbGenres, err := listGenres(ctx, tx, GenreQueryParams{
		QueryParams: QueryParams{
			Offset: offset, Limit: limit,
		},
	})

	if err != nil {
		return nil, 0, err
	}

	genres := make([]artist.Genre, 0)
	for _, dbGenre := range dbGenres {
		genres = append(genres, artist.Genre{
			ID:   dbGenre.ID,
			Name: dbGenre.Name,
		})
	}

	totalCount, err := genreCount(ctx, tx)
	if err != nil {
		return nil, 0, err
	}

	if err := tx.Commit(); err != nil {
		return nil, 0, err
	}

	return genres, totalCount, nil
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

// Lists genres based on the input {GenreQueryParams}.
func listGenres(ctx context.Context, tx *sql.Tx, params GenreQueryParams) ([]Genre, error) {
	query, err := NewQuery(`SELECT id, name FROM genre`)
	if err != nil {
		return nil, err
	}

	err = query.WithLimit(params.Limit)
	if err != nil {
		return nil, err
	}

	err = query.WithOffset(params.Offset)
	if err != nil {
		return nil, err
	}

	queryString, args := query.Build()

	rows, err := tx.QueryContext(ctx, queryString, args...)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	genres := make([]Genre, 0)

	for rows.Next() {
		var id int64
		var name string

		if err := rows.Scan(&id, &name); err != nil {
			return nil, err
		}

		genres = append(genres, Genre{ID: id, Name: name})
	}

	return genres, nil
}

// Returns the total amount of genres in the database.
func genreCount(ctx context.Context, tx *sql.Tx) (int, error) {
	query := `SELECT COUNT(*) FROM genre`

	var totalCount int

	err := tx.QueryRowContext(ctx, query).Scan(&totalCount)
	if err != nil {
		return 0, err
	}

	return totalCount, nil
}

// Inserts the genre with the given name.
//
// If the attemptedly inserted genre already exists, the previous genre is
// returned.
func insertGenre(ctx context.Context, tx *sql.Tx, name string) (int64, error) {
	// Return exising genre, if exists.
	var id int64
	query := `SELECT id FROM genre WHERE name = @name`

	err := tx.QueryRowContext(ctx, query, sql.Named("name", name)).Scan(&id)
	if err == nil {
		return id, nil
	}

	query = `INSERT INTO genre (name) VALUES (@name) RETURNING id`
	res, err := tx.ExecContext(ctx, query, sql.Named("name", name))
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
	query := `
	INSERT INTO artists_genres (artist_id, genre_id)
	VALUES (@artist_id, @genre_id)`

	_, err := tx.ExecContext(ctx, query,
		sql.Named("artist_id", artistID),
		sql.Named("genre_id", genreID),
	)

	if err != nil {
		return err
	}

	return nil
}

// Dissasociates the given genre from the given artist.
func dissasociateArtistFromGenre(ctx context.Context, tx *sql.Tx, artistID int64, genreID int64) error {
	query, err := NewQuery("DELETE FROM artists_genres")
	if err != nil {
		return err
	}

	err = query.WithFilters(map[string]any{
		"artist_id = ?": artistID,
		"genre_id = ?":  genreID,
	})

	if err != nil {
		return err
	}

	queryStr, args := query.Build()

	_, err = tx.ExecContext(ctx, queryStr, args...)
	if err != nil {
		return err
	}

	return nil
}

// Lists all genres associated with an artist.
func artistGenres(ctx context.Context, tx *sql.Tx, artistID int64) ([]Genre, error) {
	query := `
	SELECT id, name FROM genre g
	JOIN artists_genres ag ON ag.genre_id = g.id
	WHERE ag.artist_id = @artist_id`

	rows, err := tx.QueryContext(ctx, query, sql.Named("artist_id", artistID))
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	genres := make([]Genre, 0)

	for rows.Next() {
		var id int64
		var name string

		if err := rows.Scan(&id, &name); err != nil {
			return nil, err
		}

		genres = append(genres, Genre{ID: id, Name: name})
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return genres, nil
}

// Gets a genre by its ID.
func genreByID(ctx context.Context, tx *sql.Tx, genreID int64) (Genre, error) {
	query := `SELECT name FROM genre WHERE id = @genre_id`

	var name string
	err := tx.QueryRowContext(ctx, query,
		sql.Named("genre_id", genreID),
	).Scan(&name)

	if err != nil {
		return Genre{}, err
	}

	return Genre{
		ID:   genreID,
		Name: name,
	}, nil
}

// Gets a genre by its name.
func genreByName(ctx context.Context, tx *sql.Tx, name string) (Genre, error) {
	query, err := NewQuery("SELECT id FROM genre")
	if err != nil {
		return Genre{}, err
	}

	err = query.AddFilter("name = ?", name)
	if err != nil {
		return Genre{}, err
	}

	queryStr, args := query.Build()

	var id int64

	err = tx.QueryRowContext(ctx, queryStr, args...).Scan(&id)
	if err != nil {
		return Genre{}, err
	}

	return Genre{ID: id, Name: name}, nil
}
