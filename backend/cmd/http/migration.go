package main

import (
	"context"
	"database/sql"
	"fmt"
	"log/slog"
	"os"
	"path"
	"slices"
	"strings"
	"time"
)

type migrationFile struct {
	entry     os.DirEntry
	timestamp time.Time
}

var migrations []migrationFile

func parseMigrationFilename(fileName string) (time.Time, string, error) {
	parts := strings.Split(fileName, "_")

	if len(parts) != 2 {
		return time.Time{}, "", fmt.Errorf("Filename must consist of two parts separated by '_'")
	}

	t, err := time.Parse("020106-15-04", parts[0])
	if err != nil {
		return time.Time{}, "", fmt.Errorf("Could not parse timestamp for file %q: %w", fileName, err)
	}

	return t, parts[1], nil
}

func ensureMigrationsTable(ctx context.Context, tx *sql.Tx) error {
	query := `
	CREATE TABLE IF NOT EXISTS schema_migrations (
		version TEXT PRIMARY KEY,
		applied_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP
)`

	_, err := tx.ExecContext(ctx, query)
	if err != nil {
		return fmt.Errorf("Could not execute schema_migrations query: %v", err)
	}

	return nil
}

func appliedMigrations(ctx context.Context, tx *sql.Tx) (map[string]bool, error) {
	rows, err := tx.QueryContext(ctx, "SELECT version FROM schema_migrations")
	if err != nil {
		return nil, fmt.Errorf("Could not query 'version' from 'schema_migrations': %w", err)
	}

	defer rows.Close()

	applied := make(map[string]bool)

	for rows.Next() {
		var version string

		if err := rows.Scan(&version); err != nil {
			return nil, fmt.Errorf("Could not scan version: %v", err)
		}

		applied[version] = true
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("Something went wrong executing query: %v", err)
	}

	return applied, nil
}

func applyMigrations(ctx context.Context, db *sql.DB, directory string) error {
	tx, err := db.BeginTx(ctx, nil)
	if err != nil {
		return fmt.Errorf("Could not begin transaction: %v", err)
	}

	defer tx.Rollback()

	if err := ensureMigrationsTable(ctx, tx); err != nil {
		return fmt.Errorf("Could not ensure migration table: %v", err)
	}

	files, err := os.ReadDir(directory)
	if err != nil {
		return fmt.Errorf("Could not read directory: %v", err)
	}

	for _, entry := range files {
		ts, _, err := parseMigrationFilename(entry.Name())
		if err != nil {
			return fmt.Errorf("Could not parse migration filename for file %q: %w", entry.Name(), err)
		}

		migrations = append(migrations, migrationFile{
			entry:     entry,
			timestamp: ts,
		})
	}

	slices.SortFunc(migrations, func(a, b migrationFile) int {
		return a.timestamp.Compare(b.timestamp)
	})

	applied, err := appliedMigrations(ctx, tx)
	if err != nil {
		return fmt.Errorf("Could not get applied migrations: %w", err)
	}

	for _, migration := range migrations {
		version := migration.entry.Name()

		if applied[version] {
			slog.Info("[SKIP] migration", "version", version)
			continue
		}

		slog.Info("[APPLY] migration", "version", version)
		filePath := path.Join(directory, migration.entry.Name())

		contents, err := os.ReadFile(filePath)
		if err != nil {
			return fmt.Errorf("Could not read file %q: %v", filePath, err)
		}

		if _, err := tx.ExecContext(ctx, string(contents)); err != nil {
			return fmt.Errorf("Could not execute migration: %v", err)
		}

		query := `
			INSERT INTO schema_migrations (version) 
			VALUES (?)
		`

		if _, err := tx.ExecContext(ctx, query, version); err != nil {
			return fmt.Errorf("Could not record schema migration %q: %w", version, err)
		}
	}

	if err := tx.Commit(); err != nil {
		return fmt.Errorf("Could not commit transaction: %v", err)
	}

	return nil
}
