package sqlite

import (
	"context"
	"database/sql"
	"embed"
	"fmt"
	"io/fs"
	"sort"
	"strings"

	_ "github.com/mattn/go-sqlite3"
)

//go:embed migration/*.sql
var migrationFS embed.FS

type Store struct {
	*sql.DB
	dsn string
}

func NewStore(dsn string) *Store {
	return &Store{
		dsn: dsn,
	}
}

func (store *Store) Open() error {
	if strings.TrimSpace(store.dsn) == "" {
		return fmt.Errorf("No DSN specified")
	}
	var err error

	store.DB, err = sql.Open("sqlite3", store.dsn)
	if err != nil {
		return fmt.Errorf("Could not open sqlite database: %v", err)
	}

	err = store.Ping()
	if err != nil {
		return fmt.Errorf("Could not ping database: %v", err)
	}

	err = store.migrate()
	if err != nil {
		return fmt.Errorf("Could not migrate: %v", err)
	}

	return nil
}

func (store Store) migrate() error {
	query := `
	CREATE TABLE IF NOT EXISTS migrations (
		name TEXT PRIMARY KEY
	)`

	if _, err := store.Exec(query); err != nil {
		return fmt.Errorf("Could not create migrations table: %v", err)
	}

	names, err := fs.Glob(migrationFS, "migration/*.sql")
	if err != nil {
		return err
	}

	sort.Strings(names)

	for _, name := range names {
		if err = store.migrateFile(name); err != nil {
			return fmt.Errorf("Could not migrate file %q: %v", name, err)
		}
	}

	return nil
}

func (store Store) migrateFile(fileName string) error {
	tx, err := store.BeginTx(context.Background(), nil)
	if err != nil {
		return err
	}

	defer tx.Rollback()

	var n int
	query := "SELECT COUNT(*) FROM migrations WHERE name = ?"

	err = tx.QueryRow(query, fileName).Scan(&n)
	if err != nil {
		return err
	}

	if n != 0 {
		return nil
	}

	buf, err := fs.ReadFile(migrationFS, fileName)
	if err != nil {
		return err
	}

	_, err = tx.Exec(string(buf))
	if err != nil {
		return err
	}

	query = "INSERT INTO migrations (name) values (?)"

	_, err = tx.Exec(query, fileName)
	if err != nil {
		return err
	}

	if err = tx.Commit(); err != nil {
		return err
	}

	return nil
}

func (store Store) Close() error {
	if store.DB == nil {
		return nil
	}

	return store.DB.Close()
}

func formatLimitOffset(limit int, offset int) string {
	if limit > 0 && offset > 0 {
		return fmt.Sprintf("LIMIT %d OFFSET %d", limit, offset)
	}

	if limit > 0 {
		return fmt.Sprintf("LIMIT %d", limit)
	}

	if offset > 0 {
		return fmt.Sprintf("OFFSET %d", offset)
	}

	return ""
}
