package main

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5"
)

func createVenueTable(ctx context.Context, tx pgx.Tx) error {
	query := `
	CREATE TABLE IF NOT EXISTS venue (
		id INTEGER GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
		name TEXT UNIQUE NOT NULL,
		city TEXT NOT NULL,
		country TEXT NOT NULL,

		created_at TIMESTAMP NOT NULL DEFAULT NOW(),
		updated_at TIMESTAMP NOT NULL DEFAULT NOW(),
		created_by INTEGER NOT NULL REFERENCES member(id),
		updated_by INTEGER NOT NULL REFERENCES member(id)
	);

	CREATE TRIGGER update_venue_timestamps BEFORE INSERT OR UPDATE ON venue
	FOR EACH ROW EXECUTE FUNCTION update_timestamps();`

	if _, err := tx.Exec(ctx, query); err != nil {
		return fmt.Errorf("Could not execute query: %v", err)
	}

	return nil
}
