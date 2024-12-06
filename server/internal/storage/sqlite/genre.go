package sqlite

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"strings"

	"github.com/mattismoel/konnekt"
	"github.com/mattismoel/konnekt/internal/service"
	"github.com/mattismoel/konnekt/internal/storage"
)

type genreRepository struct {
	store *Store
}

func NewGenreRepository(store *Store) *genreRepository {
	return &genreRepository{store: store}
}

func (s genreRepository) GenreByID(ctx context.Context, id int64) (storage.Genre, error) {
	tx, err := s.store.BeginTx(ctx, nil)
	if err != nil {
		return storage.Genre{}, err
	}

	defer tx.Rollback()

	genre, err := findGenreByID(ctx, tx, id)
	if err != nil {
		return storage.Genre{}, err
	}

	if err = tx.Commit(); err != nil {
		return storage.Genre{}, err
	}

	return genre, nil
}

func (s genreRepository) FindGenres(ctx context.Context, filter service.GenreFilter) ([]storage.Genre, error) {
	tx, err := s.store.BeginTx(ctx, nil)
	if err != nil {
		return nil, err
	}

	defer tx.Rollback()

	genres, err := findGenres(ctx, tx, filter)
	if err != nil {
		return nil, err
	}

	if err = tx.Commit(); err != nil {
		return nil, err
	}

	return genres, nil
}

func (s genreRepository) InsertGenre(ctx context.Context, name string) (storage.Genre, error) {
	tx, err := s.store.BeginTx(ctx, nil)
	if err != nil {
		return storage.Genre{}, err
	}

	genre, err := insertGenre(ctx, tx, name)
	if err != nil {
		return storage.Genre{}, err
	}

	defer tx.Rollback()

	if err = tx.Commit(); err != nil {
		return storage.Genre{}, err
	}

	return genre, nil
}

func (s genreRepository) UpdateGenre(ctx context.Context, id int64, newName string) (storage.Genre, error) {
	tx, err := s.store.BeginTx(ctx, nil)
	if err != nil {
		return storage.Genre{}, err
	}

	defer tx.Rollback()

	genre, err := updateGenre(ctx, tx, id, newName)
	if err != nil {
		return storage.Genre{}, err
	}

	if err = tx.Commit(); err != nil {
		return storage.Genre{}, err
	}

	return genre, nil
}

func (s genreRepository) DeleteGenre(ctx context.Context, id int64) error {
	tx, err := s.store.BeginTx(ctx, nil)
	if err != nil {
		return err
	}

	defer tx.Rollback()

	if err = deleteGenre(ctx, tx, id); err != nil {
		return err
	}

	if err = tx.Commit(); err != nil {
		return err
	}

	return nil
}

func insertGenres(ctx context.Context, tx *sql.Tx, genres []string) ([]storage.Genre, error) {
	insertedGenres := []storage.Genre{}

	for _, genreName := range genres {
		var err error
		genre, err := insertGenre(ctx, tx, genreName)
		if err != nil {
			return nil, err
		}

		insertedGenres = append(insertedGenres, genre)
	}

	return insertedGenres, nil
}

func findGenreByID(ctx context.Context, tx *sql.Tx, id int64) (storage.Genre, error) {
	genres, err := findGenres(ctx, tx, service.GenreFilter{ID: &id})
	if err != nil {
		return storage.Genre{}, err
	}

	if len(genres) <= 0 {
		return storage.Genre{}, konnekt.Errorf(konnekt.ERRNOTFOUND, "No genres found with id %d", id)
	}

	return genres[0], nil
}

func findGenreByName(ctx context.Context, tx *sql.Tx, name string) (storage.Genre, error) {
	genres, err := findGenres(ctx, tx, service.GenreFilter{Name: &name})
	if err != nil {
		return storage.Genre{}, err
	}

	if len(genres) <= 0 {
		return storage.Genre{}, konnekt.Errorf(konnekt.ERRNOTFOUND, "No genres found with name %q", name)
	}

	return genres[0], nil
}

func updateGenre(ctx context.Context, tx *sql.Tx, id int64, newName string) (storage.Genre, error) {
	genre, err := findGenreByID(ctx, tx, id)
	if err != nil {
		return storage.Genre{}, err
	}

	query := `
	UPDATE genres
	SET
		name = @name
	WHERE id = @id`

	_, err = tx.ExecContext(ctx, query,
		sql.Named("name", newName),
		sql.Named("id", id),
	)

	if err != nil {
		return storage.Genre{}, err
	}

	return genre, nil
}

func deleteGenre(ctx context.Context, tx *sql.Tx, id int64) error {
	query := "DELETE FROM genre WHERE id = ?"

	_, err := tx.ExecContext(ctx, query, id)
	if err != nil {
		return err
	}

	return nil
}

func deleteEventGenres(ctx context.Context, tx *sql.Tx, eventID int64) error {
	query := "DELETE FROM events_genres WHERE event_id = ?"

	_, err := tx.ExecContext(ctx, query, eventID)
	if err != nil {
		return err
	}

	return nil
}

func updateEventGenres(ctx context.Context, tx *sql.Tx, eventID int64, genreNames []string) error {
	err := deleteEventGenres(ctx, tx, eventID)
	if err != nil {
		return fmt.Errorf("Could not delete event genres: %v", err)
	}

	for _, genreName := range genreNames {
		err = relateGenreToEvent(ctx, tx, eventID, genreName)
		if err != nil {
			return fmt.Errorf("Could not relate genre %q to event with name %q: %v", genreName, eventID, err)
		}
	}

	return nil
}

func relateGenreToEvent(ctx context.Context, tx *sql.Tx, eventID int64, name string) error {
	genre, err := findGenreByName(ctx, tx, name)
	if err != nil {
		return err
	}

	query := "INSERT INTO events_genres (event_id, genre_id) VALUES (?, ?)"

	_, err = tx.ExecContext(ctx, query, eventID, genre.ID)
	if err != nil {
		return err
	}

	return nil
}

func findEventGenres(ctx context.Context, tx *sql.Tx, eventId int64) ([]storage.Genre, error) {
	query := `
	SELECT 
		id, 
		name 
	FROM genre 
	JOIN events_genres ON genre.id = events_genres.genre_id
	WHERE events_genres.event_id = ?`

	rows, err := tx.QueryContext(ctx, query, eventId)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	genres := []storage.Genre{}

	for rows.Next() {
		var genre storage.Genre
		if err := rows.Scan(&genre.ID, &genre.Name); err != nil {
			return nil, err
		}

		genres = append(genres, genre)
	}

	return genres, nil
}

func doesGenreExist(ctx context.Context, tx *sql.Tx, name string) (bool, error) {
	var exists bool
	query := "SELECT EXISTS(SELECT 1 FROM genre WHERE name = ?)"
	err := tx.QueryRowContext(ctx, query, name).Scan(&exists)
	if err != nil && !errors.Is(sql.ErrNoRows, err) {
		return false, err
	}

	return exists, nil
}

func genreIDs(genres []storage.Genre) []int64 {
	ids := []int64{}

	for _, genre := range genres {
		ids = append(ids, genre.ID)

	}

	return ids
}

func genreNames(genres []storage.Genre) []string {
	names := []string{}

	for _, genre := range genres {
		names = append(names, genre.Name)

	}

	return names
}

func insertGenre(ctx context.Context, tx *sql.Tx, name string) (storage.Genre, error) {
	exists, err := doesGenreExist(ctx, tx, name)
	if err != nil {
		return storage.Genre{}, err
	}

	if exists {
		genre, err := findGenreByName(ctx, tx, name)
		if err != nil && !errors.Is(sql.ErrNoRows, err) {
			return storage.Genre{}, err
		}

		return genre, nil
	}

	query := "INSERT OR IGNORE INTO genre (name) VALUES (?) ON CONFLICT (name) DO NOTHING"

	_, err = tx.ExecContext(ctx, query, name)
	if err != nil {
		return storage.Genre{}, err
	}

	query = "SELECT id, name FROM genre WHERE name = ?"

	var genre storage.Genre

	err = tx.QueryRowContext(ctx, query, name).Scan(&genre.ID, &genre.Name)
	if err != nil {
		return storage.Genre{}, err
	}

	return genre, nil
}

func findGenres(ctx context.Context, tx *sql.Tx, filter service.GenreFilter) ([]storage.Genre, error) {
	genres := []storage.Genre{}
	var n int

	where, args := []string{"1 = 1\n"}, []any{}

	if v := filter.ID; v != nil {
		where, args = append(where, "id = ?\n"), append(args, *v)
	}

	// TODO: Implement
	// if v := filter.EventID; v != nil {}

	if v := filter.Name; v != nil {
		where, args = append(where, "name = ?"), append(args, *v)
	}

	query := `
	SELECT 
		id,
		name,
		COUNT(*) OVER()
	FROM genre
	WHERE `

	query += strings.Join(where, " AND ")
	query += "ORDER BY name ASC\n"
	query += formatLimitOffset(filter.Limit, filter.Offset)

	rows, err := tx.QueryContext(ctx, query, args...)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	for rows.Next() {
		var genre storage.Genre

		err := rows.Scan(&genre.ID, &genre.Name, &n)
		if err != nil {
			return nil, err
		}

		genres = append(genres, genre)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return genres, nil
}
