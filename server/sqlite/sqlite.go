package sqlite

import (
	"context"
	"database/sql"
	"database/sql/driver"
	"embed"
	"fmt"
	"io/fs"
	"sort"
	"strings"
	"time"

	_ "github.com/mattn/go-sqlite3"
)

//go:embed migration/*.sql
var migrationFS embed.FS

type Repository struct {
	db  *sql.DB
	dsn string
}

func New(dsn string) *Repository {
	r := &Repository{
		dsn: strings.TrimSpace(dsn),
	}

	return r
}

func (repo *Repository) Open() error {
	if strings.TrimSpace(repo.dsn) == "" {
		return fmt.Errorf("No DSN specified")
	}

	db, err := sql.Open("sqlite3", repo.dsn)
	if err != nil {
		return fmt.Errorf("Could not open sqlite database: %v", err)
	}

	err = db.Ping()
	if err != nil {
		return fmt.Errorf("Could not ping database: %v", err)
	}

	repo.db = db

	err = repo.migrate()
	if err != nil {
		return fmt.Errorf("Could not migrate: %v", err)
	}

	return nil
}

func (repo Repository) migrate() error {
	query := `
	CREATE TABLE IF NOT EXISTS migrations (
		name TEXT PRIMARY KEY
	)`

	if _, err := repo.db.Exec(query); err != nil {
		return fmt.Errorf("Could not create migrations table: %v", err)
	}

	names, err := fs.Glob(migrationFS, "migration/*.sql")
	if err != nil {
		return err
	}

	sort.Strings(names)

	for _, name := range names {
		if err = repo.migrateFile(name); err != nil {
			return fmt.Errorf("Could not migrate file %q: %v", name, err)
		}
	}

	return nil
}

func (repo Repository) migrateFile(fileName string) error {
	err := repo.withTx(context.Background(), nil, func(tx *sql.Tx) error {

		var n int
		query := "SELECT COUNT(*) FROM migrations WHERE name = ?"

		err := tx.QueryRow(query, fileName).Scan(&n)
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

		return nil
	})

	if err != nil {
		return err
	}

	return nil
}

func (r *Repository) Close() error {
	if r.db == nil {
		return nil
	}

	return r.db.Close()
}

func (r Repository) withTx(ctx context.Context, opts *sql.TxOptions, fn func(tx *sql.Tx) error) error {
	tx, err := r.db.BeginTx(ctx, opts)
	if err != nil {
		return fmt.Errorf("Could not begin transaction: %v", err)
	}

	defer tx.Rollback()

	if err = fn(tx); err != nil {
		return err
	}

	if err = tx.Commit(); err != nil {
		return fmt.Errorf("Could not commit transaction: %v", err)
	}

	return nil
}

type NullTime time.Time

func (n *NullTime) Scan(value any) error {
	if value == nil {
		*(*time.Time)(n) = time.Time{}
		return nil
	}

	if value, ok := value.(string); ok {
		*(*time.Time)(n), _ = time.Parse(time.RFC3339, value)
		return nil
	}

	return fmt.Errorf("Cannot value of type %t scan to time.Time", value)
}

func (n *NullTime) Value() (driver.Value, error) {
	if n == nil || (*time.Time)(n).IsZero() {
		return nil, nil
	}

	return (*time.Time)(n).UTC().Format(time.RFC3339), nil
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
