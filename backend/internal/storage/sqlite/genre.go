package sqlite

import (
	"context"
	"database/sql"

	"github.com/mattismoel/konnekt/internal/domain/artist"
)

type GenreQueryParams struct {
	QueryParams
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

func genreCount(ctx context.Context, tx *sql.Tx) (int, error) {
	query := `SELECT COUNT(*) FROM genre`

	var totalCount int

	err := tx.QueryRowContext(ctx, query).Scan(&totalCount)
	if err != nil {
		return 0, err
	}

	return totalCount, nil
}
