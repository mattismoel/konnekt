package psql

import (
	"context"
	"fmt"

	sq "github.com/Masterminds/squirrel"
	"github.com/jackc/pgx/v5"
	konnekt "github.com/mattismoel/konnekt/backend"
)

type Genre struct {
	ID   int64  `db:"id"`
	Name string `db:"name"`
	Timestamps
}

// GenreByID implements konnekt.ArtistRepo.
func (a ArtistRepo) GenreByID(ctx context.Context, genreID int64) (konnekt.Genre, error) {
	var genre konnekt.Genre
	err := pgx.BeginFunc(ctx, a.Pool, func(tx pgx.Tx) error {
		dbGenre, err := genreByID(ctx, tx, int64(genreID))
		if err != nil {
			return fmt.Errorf("Could not get genre with ID %d: %v", genreID, err)
		}

		genre = dbGenre.ToDomain()
		return nil
	})

	if err != nil {
		return "", err
	}

	return genre, nil
}

// InsertGenre implements konnekt.ArtistRepo.
func (a ArtistRepo) InsertGenre(ctx context.Context, cg konnekt.CreateGenre) (int64, error) {
	var genreID int64
	err := pgx.BeginFunc(ctx, a.Pool, func(tx pgx.Tx) error {
		id, err := insertGenre(ctx, tx, Genre{
			Name: cg.Name,
		})

		if err != nil {
			return fmt.Errorf("Could not insert genre: %v", err)
		}

		genreID = id
		return nil
	})

	if err != nil {
		return 0, err
	}

	return genreID, nil
}

func genreByID(ctx context.Context, tx pgx.Tx, genreID int64) (Genre, error) {
	query, args, err := psql.
		Select("*").
		From("genre").
		Where(sq.Eq{"id": genreID}).
		ToSql()

	if err != nil {
		return Genre{}, NewQueryBuildError("genre by id", err)
	}

	rows, err := tx.Query(ctx, query, args...)
	if err != nil {
		return Genre{}, err
	}

	genre, err := pgx.CollectOneRow(rows, pgx.RowToStructByName[Genre])
	if err != nil {
		return Genre{}, err
	}

	return genre, nil
}

func insertGenre(ctx context.Context, tx pgx.Tx, g Genre) (int64, error) {
	query, args, err := psql.
		Insert("genre").
		Columns("name").
		Values(g.Name).
		Suffix("ON CONFLICT (name) DO UPDATE SET name = EXCLUDED.name").
		Suffix("RETURNING id").
		ToSql()

	if err != nil {
		return 0, NewQueryBuildError("insert genre", err)
	}

	var genreID int64
	if err := tx.QueryRow(ctx, query, args...).Scan(&genreID); err != nil {
		return 0, fmt.Errorf("Could not insert genre: %v", err)
	}

	return genreID, nil
}

func artistGenres(ctx context.Context, tx pgx.Tx, artistID int64) (Collection[Genre, konnekt.Genre], error) {
	query, args, err := psql.
		Select("genre.*").
		From("genre").
		Join("artists_genres ag on ag.genre_id = genre.id").
		Where(sq.Eq{"ag.artist_id": artistID}).
		ToSql()

	if err != nil {
		return nil, NewQueryBuildError("artist genres", err)
	}

	rows, err := tx.Query(ctx, query, args...)
	if err != nil {
		return nil, fmt.Errorf("Could not query for artist genres: %v", err)
	}

	genres, err := pgx.CollectRows(rows, pgx.RowToStructByName[Genre])
	if err != nil {
		return nil, err
	}

	return genres, nil
}

func clearArtistGenres(ctx context.Context, tx pgx.Tx, artistID int64) error {
	query, args, err := psql.Delete("artists_genres").Where(sq.Eq{"artist_id": artistID}).ToSql()
	if err != nil {
		return NewQueryBuildError("clear artist genres", err)
	}

	if _, err := tx.Exec(ctx, query, args...); err != nil {
		return err
	}

	return nil
}

func setArtistGenres(ctx context.Context, tx pgx.Tx, artistID int64, genreIDs ...int64) error {
	if err := clearArtistGenres(ctx, tx, artistID); err != nil {
		return fmt.Errorf("Could not delete artist genres: %v", err)
	}

	builder := psql.Insert("artists_genres").Columns("artist_id", "genre_id")

	for _, genreID := range genreIDs {
		builder = builder.Values(artistID, genreID)
	}

	query, args, err := builder.ToSql()
	if err != nil {
		return NewQueryBuildError("insert artist genre", err)
	}

	if _, err := tx.Exec(ctx, query, args...); err != nil {
		return fmt.Errorf("Could not insert artist genres: %v", err)
	}

	return nil
}

func (g Genre) ToDomain() konnekt.Genre {
	return konnekt.Genre(g.Name)
}
