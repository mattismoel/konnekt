package main

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5"
)

func createSessionTable(ctx context.Context, tx pgx.Tx) error {
	query := `
	CREATE TABLE IF NOT EXISTS session (
		id INTEGER GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
		session_id TEXT UNIQUE NOT NULL,
		member_id INTEGER REFERENCES member(id),
		secret_hash BYTEA NOT NULL,
		created_at TIMESTAMP NOT NULL
	);`

	if _, err := tx.Exec(ctx, query); err != nil {
		return fmt.Errorf("Could not execute query: %v", err)
	}

	return nil
}
