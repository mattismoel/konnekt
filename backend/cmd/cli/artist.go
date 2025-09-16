package main

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5"
	konnekt "github.com/mattismoel/konnekt/backend"
)

var genres = []konnekt.Genre{"Rock", "Punk", "Pop", "Hip-Hop", "Rap", "Folk",
	"Indie", "R&B", "Elektronisk", "Country", "Jazz", "Blues", "Funk", "Latin"}

func createArtistTables(ctx context.Context, tx pgx.Tx) error {
	query := `
	CREATE TABLE IF NOT EXISTS artist (
		id INTEGER GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
		name TEXT UNIQUE NOT NULL,
		description TEXT NOT NULL,
		image_url TEXT NOT NULL,
		preview_url TEXT,

		created_at TIMESTAMP NOT NULL DEFAULT NOW(),
		updated_at TIMESTAMP NOT NULL DEFAULT NOW(),
		created_by INTEGER NOT NULL REFERENCES member(id),
		updated_by INTEGER NOT NULL REFERENCES member(id)
	);


	CREATE TABLE IF NOT EXISTS social (
		id INTEGER GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
		url TEXT NOT NULL,
		artist_id INTEGER NOT NULL REFERENCES artist(id)
	);

	CREATE TABLE IF NOT EXISTS genre (
		id INTEGER GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
		name TEXT UNIQUE NOT NULL,

		created_at TIMESTAMP NOT NULL DEFAULT NOW(),
		updated_at TIMESTAMP NOT NULL DEFAULT NOW()
	);

	CREATE TABLE IF NOT EXISTS artists_genres (
		artist_id INTEGER NOT NULL REFERENCES ARTIST(ID),
		genre_id INTEGER NOT NULL REFERENCES GENRE(ID),

		PRIMARY KEY (artist_id, genre_id)
	);

	CREATE TRIGGER update_artist_timestamps BEFORE INSERT OR UPDATE ON artist
	FOR EACH ROW EXECUTE FUNCTION update_timestamps();

	CREATE TRIGGER update_genre_timestamps BEFORE INSERT OR UPDATE ON genre 
	FOR EACH ROW EXECUTE FUNCTION update_timestamps();`

	if _, err := tx.Exec(ctx, query); err != nil {
		return fmt.Errorf("Could not execute query: %v", err)
	}

	return nil
}

func seedGenres(ctx context.Context, tx pgx.Tx) error {
	builder := psql.Insert("genre").Columns("name")

	for _, genreName := range genres {
		builder = builder.Values(genreName)
	}

	query, args, err := builder.ToSql()
	if err != nil {
		return fmt.Errorf("Could not create seed genres query: %v", err)
	}

	if _, err := tx.Exec(ctx, query, args...); err != nil {
		return fmt.Errorf("Could not insert genres: %v", err)
	}

	return nil
}
