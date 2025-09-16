package main

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5"
)

func createMemberTables(ctx context.Context, tx pgx.Tx) error {
	query := `
	CREATE TABLE IF NOT EXISTS member (
		id INTEGER GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
		email TEXT UNIQUE NOT NULL,
		first_name TEXT NOT NULL,
		last_name TEXT NOT NULL,
		password_hash CHAR(60) NOT NULL,
		avatar_url TEXT NOT NULL,
		special_role TEXT,
		approved BOOLEAN NOT NULL DEFAULT FALSE,

		created_at TIMESTAMP NOT NULL DEFAULT NOW(),
		updated_at TIMESTAMP NOT NULL DEFAULT NOW()
	);

	CREATE TRIGGER update_member_timestamps BEFORE INSERT OR UPDATE ON member
	FOR EACH ROW EXECUTE FUNCTION update_timestamps();`

	if _, err := tx.Exec(ctx, query); err != nil {
		return fmt.Errorf("Could not execute query: %v", err)
	}

	return nil
}
