package sqlite_test

import (
	"context"
	"testing"

	"github.com/mattismoel/konnekt"
	"github.com/mattismoel/konnekt/sqlite"
)

var baseGenre = konnekt.Genre{
	ID:   1,
	Name: "Rock",
}

func TestCreateGenre(t *testing.T) {
	type test struct {
		genre    konnekt.Genre
		wantCode string
		wantID   int64
	}

	tests := map[string]test{
		"Valid Genre": {
			genre:    baseGenre,
			wantCode: "",
			wantID:   1,
		},
		"Empty Name": {
			genre: konnekt.Genre{
				Name: " ",
			},
			wantCode: konnekt.ERRINVALID,
			wantID:   0,
		},
	}

	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			repo, dsn := MustOpenRepo(t)
			defer MustCloseRepo(t, repo, dsn)

			service := sqlite.NewGenreService(repo)

			genre, err := service.CreateGenre(context.Background(), tt.genre)
			code := konnekt.ErrorCode(err)

			if code != tt.wantCode {
				t.Fatalf("got code %q, want code %q, error: %v", code, tt.wantCode, err)
			}

			if genre.ID != tt.wantID {
				t.Fatalf("got id %d, want id %d", genre.ID, tt.wantID)
			}
		})
	}
}

func TestFindGenreByID(t *testing.T) {
	type test struct {
		id        int64
		wantGenre konnekt.Genre
		wantCode  string
	}

	tests := map[string]test{
		"Valid ID": {
			id:        1,
			wantGenre: baseGenres[0],
			wantCode:  "",
		},
		"Invalid ID": {
			id:        999,
			wantGenre: konnekt.Genre{},
			wantCode:  konnekt.ERRNOTFOUND,
		},
	}

	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			repo, dsn := MustOpenRepo(t)
			defer MustCloseRepo(t, repo, dsn)

			service := sqlite.NewGenreService(repo)
			MustCreateGenre(t, repo, context.Background(), baseGenres[0])

			genre, err := service.GenreByID(context.Background(), tt.id)

			code := konnekt.ErrorCode(err)

			if code != tt.wantCode {
				t.Fatalf("got code %q, want code %q, error: %v", code, tt.wantCode, err)
			}

			if !genre.Equals(tt.wantGenre) {
				t.Fatalf("got %+v, want %+v", genre, tt.wantGenre)
			}
		})
	}
}

func TestDeleteGenre(t *testing.T) {
	type test struct {
		id       int64
		wantCode string
	}

	tests := map[string]test{
		"Valid ID": {
			id:       1,
			wantCode: "",
		},
		"Invalid ID": {
			id:       999,
			wantCode: "", // No return code as we dont care if the genre exists.
		},
	}

	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			repo, dsn := MustOpenRepo(t)
			defer MustCloseRepo(t, repo, dsn)

			service := sqlite.NewGenreService(repo)
			MustCreateGenre(t, repo, context.Background(), baseGenres[0])

			err := service.DeleteGenre(context.Background(), tt.id)

			code := konnekt.ErrorCode(err)
			if code != tt.wantCode {
				t.Fatalf("got code %q, want code %q", code, tt.wantCode)
			}
		})
	}
}

func MustCreateGenre(t testing.TB, repo *sqlite.Repository, ctx context.Context, genre konnekt.Genre) {
	t.Helper()

	_, err := sqlite.NewGenreService(repo).CreateGenre(ctx, genre)
	if err != nil {
		t.Fatal(err)
	}
}
