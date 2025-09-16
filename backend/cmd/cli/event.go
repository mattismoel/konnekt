package main

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5"
)

func createEventTables(ctx context.Context, tx pgx.Tx) error {
	query := `
	CREATE TABLE IF NOT EXISTS event (
		id INTEGER GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
		title TEXT NOT NULL,
		description TEXT NOT NULL,
		ticket_url TEXT NOT NULL,
		image_url TEXT NOT NULL,
		venue_id INTEGER NOT NULL REFERENCES venue(id)
	);

	CREATE TABLE IF NOT EXISTS concert (
		id INTEGER GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
		from_date TIMESTAMP NOT NULL,
		to_date TIMESTAMP NOT NULL,
		event_id INTEGER NOT NULL REFERENCES event(id),
		artist_id INTEGER NOT NULL REFERENCES artist(id)
	);`

	if _, err := tx.Exec(ctx, query); err != nil {
		return fmt.Errorf("Could not execute query: %v", err)
	}

	return nil
}
