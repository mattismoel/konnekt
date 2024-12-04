package sqlite

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"strings"

	"github.com/mattismoel/konnekt"
)

type genreService struct {
	repo *Repository
}

func NewGenreService(repo *Repository) *genreService {
	return &genreService{repo: repo}
}

func (s genreService) GenreByID(ctx context.Context, id int64) (konnekt.Genre, error) {
	tx, err := s.repo.db.BeginTx(ctx, nil)
	if err != nil {
		return konnekt.Genre{}, err
	}

	defer tx.Rollback()

	genre, err := findGenreByID(ctx, tx, id)
	if err != nil {
		return konnekt.Genre{}, err
	}

	if err = tx.Commit(); err != nil {
		return konnekt.Genre{}, err
	}

	return genre, nil
}

func (s genreService) FindGenres(ctx context.Context, filter konnekt.GenreFilter) ([]konnekt.Genre, error) {
	tx, err := s.repo.db.BeginTx(ctx, nil)
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

func (s genreService) CreateGenre(ctx context.Context, genre konnekt.Genre) (konnekt.Genre, error) {
	tx, err := s.repo.db.BeginTx(ctx, nil)
	if err != nil {
		return konnekt.Genre{}, err
	}

	genre, err = insertGenre(ctx, tx, genre)
	if err != nil {
		return konnekt.Genre{}, err
	}

	defer tx.Rollback()

	if err = tx.Commit(); err != nil {
		return konnekt.Genre{}, err
	}

	return genre, nil
}

func (s genreService) UpdateGenre(ctx context.Context, id int64, update konnekt.GenreUpdate) (konnekt.Genre, error) {
	tx, err := s.repo.db.BeginTx(ctx, nil)
	if err != nil {
		return konnekt.Genre{}, err
	}

	defer tx.Rollback()

	genre, err := updateGenre(ctx, tx, id, update)
	if err != nil {
		return konnekt.Genre{}, err
	}

	if err = tx.Commit(); err != nil {
		return konnekt.Genre{}, err
	}

	return genre, nil
}

func (s genreService) DeleteGenre(ctx context.Context, id int64) error {
	tx, err := s.repo.db.BeginTx(ctx, nil)
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

func insertGenres(ctx context.Context, tx *sql.Tx, genres []konnekt.Genre) ([]konnekt.Genre, error) {
	for _, genre := range genres {
		var err error
		genre, err = insertGenre(ctx, tx, genre)
		if err != nil {
			return nil, err
		}
	}

	return genres, nil
}

func findGenreByID(ctx context.Context, tx *sql.Tx, id int64) (konnekt.Genre, error) {
	genres, err := findGenres(ctx, tx, konnekt.GenreFilter{ID: &id})
	if err != nil {
		return konnekt.Genre{}, err
	}

	if len(genres) <= 0 {
		return konnekt.Genre{}, konnekt.Errorf(konnekt.ERRNOTFOUND, "No genres found with id %d", id)
	}

	return genres[0], nil
}

func findGenreByName(ctx context.Context, tx *sql.Tx, name string) (konnekt.Genre, error) {
	genres, err := findGenres(ctx, tx, konnekt.GenreFilter{Name: &name})
	if err != nil {
		return konnekt.Genre{}, err
	}

	if len(genres) <= 0 {
		return konnekt.Genre{}, konnekt.Errorf(konnekt.ERRNOTFOUND, "No genres found with name %q", name)
	}

	return genres[0], nil
}

func updateGenre(ctx context.Context, tx *sql.Tx, id int64, update konnekt.GenreUpdate) (konnekt.Genre, error) {
	genre, err := findGenreByID(ctx, tx, id)
	if err != nil {
		return konnekt.Genre{}, err
	}

	query := `
	UPDATE genres
	SET
		name = ?
	WHERE id = ?`

	if v := update.Name; v != nil {
		genre.Name = *v
	}

	err = genre.Validate()
	if err != nil {
		return konnekt.Genre{}, err
	}

	_, err = tx.ExecContext(ctx, query, genre.Name, id)
	if err != nil {
		return konnekt.Genre{}, err
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

func setEventGenres(ctx context.Context, tx *sql.Tx, eventID int64, genreNames []string) error {
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

func findEventGenres(ctx context.Context, tx *sql.Tx, eventId int64) ([]konnekt.Genre, error) {
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

	genres := []konnekt.Genre{}

	for rows.Next() {
		var genre konnekt.Genre
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

func genreIDs(genres []konnekt.Genre) []int64 {
	ids := []int64{}

	for _, genre := range genres {
		ids = append(ids, genre.ID)

	}

	return ids
}

func genreNames(genres []konnekt.Genre) []string {
	names := []string{}

	for _, genre := range genres {
		names = append(names, genre.Name)

	}

	return names
}

func insertGenre(ctx context.Context, tx *sql.Tx, genre konnekt.Genre) (konnekt.Genre, error) {
	exists, err := doesGenreExist(ctx, tx, genre.Name)
	if err != nil {
		return konnekt.Genre{}, err
	}

	if exists {
		genre, err := findGenreByName(ctx, tx, genre.Name)
		if err != nil && !errors.Is(sql.ErrNoRows, err) {
			return konnekt.Genre{}, err
		}

		return genre, nil
	}

	if err = genre.Validate(); err != nil {
		return konnekt.Genre{}, err
	}

	query := "INSERT OR IGNORE INTO genre (name) VALUES (?) ON CONFLICT (name) DO NOTHING"

	_, err = tx.ExecContext(ctx, query, genre.Name)
	if err != nil {
		return konnekt.Genre{}, err
	}

	query = "SELECT id FROM genre WHERE name = ?"

	err = tx.QueryRowContext(ctx, query, genre.Name).Scan(&genre.ID)
	if err != nil {
		return konnekt.Genre{}, err
	}

	return genre, nil
}

func findGenres(ctx context.Context, tx *sql.Tx, filter konnekt.GenreFilter) ([]konnekt.Genre, error) {
	genres := []konnekt.Genre{}
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
		var genre konnekt.Genre

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
