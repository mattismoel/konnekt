package sqlite_test

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/mattismoel/konnekt/sqlite"
)

func MustOpenRepo(tb testing.TB) (*sqlite.Repository, string) {
	tb.Helper()

	dir, err := os.MkdirTemp("", "")
	if err != nil {
		tb.Fatal(err)
	}

	dsn := filepath.Join(dir, "db")

	repo := sqlite.New(dsn)

	if err := repo.Open(); err != nil {
		tb.Fatal(err)
	}

	return repo, dsn
}

func MustCloseRepo(tb testing.TB, repo *sqlite.Repository, dsn string) {
	tb.Helper()

	if err := repo.Close(); err != nil {
		tb.Fatal(err)
	}

	err := os.RemoveAll(dsn)
	if err != nil {
		tb.Fatal(err)
	}
}

func testError(tb testing.TB, repo *sqlite.Repository, dsn string, err error) {
	tb.Helper()

	MustCloseRepo(tb, repo, dsn)
	tb.Fatal(err.Error())
}

func strPtr(s string) *string { return &s }
