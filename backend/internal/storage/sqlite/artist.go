package sqlite

import (
	"context"
	"database/sql"

	"github.com/mattismoel/konnekt/internal/domain/artist"
)

type Artist struct {
	ID          int64
	Name        string
	Description string
	ImageURL    string
}

type Genre struct {
	ID   int64
	Name string
}

type Social struct {
	ID  int64
	URL string
}

var _ artist.Repository = (*ArtistRepository)(nil)

type ArtistRepository struct {
	db *sql.DB
}

func NewArtistRepository(db *sql.DB) (*ArtistRepository, error) {
	return &ArtistRepository{
		db: db,
	}, nil
}

func (repo ArtistRepository) List(ctx context.Context, offset, limit int) ([]artist.Artist, int, error) {
	tx, err := repo.db.BeginTx(ctx, nil)
	if err != nil {
		return nil, 0, err
	}

	defer tx.Rollback()

	artists := make([]artist.Artist, 0)

	dbArtists, err := listArtists(ctx, tx)
	if err != nil {
		return nil, 0, err
	}

	for _, dbArtist := range dbArtists {
		dbGenres, err := artistGenres(ctx, tx, dbArtist.ID)
		if err != nil {
			return nil, 0, err
		}

		genres := make([]artist.Genre, 0)
		for _, dbGenre := range dbGenres {
			genres = append(genres, artist.Genre(dbGenre.Name))
		}

		dbSocials, err := artistSocials(ctx, tx, dbArtist.ID)
		if err != nil {
			return nil, 0, err
		}

		socials := make([]artist.Social, 0)
		for _, dbSocial := range dbSocials {
			socials = append(socials, artist.Social(dbSocial.URL))
		}

		artists = append(artists, dbArtist.ToInternal(genres, socials))
	}

	totalCount, err := artistCount(ctx, tx)
	if err != nil {
		return nil, 0, err
	}

	if err := tx.Commit(); err != nil {
		return nil, 0, err
	}

	return artists, totalCount, nil
}

func (repo ArtistRepository) Insert(ctx context.Context, a artist.Artist) (int64, error) {
	tx, err := repo.db.BeginTx(ctx, nil)
	if err != nil {
		return 0, err
	}

	artistID, err := insertArtist(ctx, tx, Artist{
		Name:        a.Name,
		Description: a.Description,
		ImageURL:    a.ImageURL,
	})

	if err != nil {
		return 0, err
	}

	for _, genre := range a.Genres {
		genreID, err := insertGenre(ctx, tx, string(genre))
		if err != nil {
			return 0, err
		}

		err = associateArtistWithGenre(ctx, tx, artistID, genreID)
		if err != nil {
			return 0, err
		}
	}

	for _, social := range a.Socials {
		socialID, err := insertSocial(ctx, tx, string(social))
		if err != nil {
			return 0, err
		}

		err = associateArtistWithSocial(ctx, tx, artistID, socialID)
		if err != nil {
			return 0, err
		}
	}

	defer tx.Rollback()

	if err := tx.Commit(); err != nil {
		return 0, err
	}

	return artistID, nil
}
func (repo ArtistRepository) ByID(ctx context.Context, artistID int64) (artist.Artist, error) {
	tx, err := repo.db.BeginTx(ctx, nil)
	if err != nil {
		return artist.Artist{}, err
	}

	defer tx.Rollback()

	dbArtist, err := artistByID(ctx, tx, artistID)
	if err != nil {
		return artist.Artist{}, err
	}

	dbGenres, err := artistGenres(ctx, tx, artistID)
	if err != nil {
		return artist.Artist{}, err
	}

	dbSocials, err := artistSocials(ctx, tx, artistID)
	if err != nil {
		return artist.Artist{}, err
	}

	if err := tx.Commit(); err != nil {
		return artist.Artist{}, err
	}

	genres := make([]artist.Genre, 0)
	for _, g := range dbGenres {
		genres = append(genres, artist.Genre(g.Name))
	}

	socials := make([]artist.Social, 0)
	for _, s := range dbSocials {
		socials = append(socials, artist.Social(s.URL))
	}

	return dbArtist.ToInternal(genres, socials), nil
}

func (repo ArtistRepository) GenreByID(ctx context.Context, genreID int64) (artist.Genre, error) {
	tx, err := repo.db.BeginTx(ctx, nil)
	if err != nil {
		return "", err
	}

	defer tx.Rollback()

	dbGenre, err := genreByID(ctx, tx, genreID)
	if err != nil {
		return "", err
	}

	if err := tx.Commit(); err != nil {
		return "", err
	}

	return artist.Genre(dbGenre.Name), nil
}

func listArtists(ctx context.Context, tx *sql.Tx) ([]Artist, error) {
	query := `SELECT id, name, description, image_url FROM artist a`

	rows, err := tx.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	artists := make([]Artist, 0)

	for rows.Next() {
		var id int64
		var name, description, imageURL string

		if err := rows.Scan(&id, &name, &description, &imageURL); err != nil {
			return nil, err
		}

		artists = append(artists, Artist{
			ID:          id,
			Name:        name,
			Description: description,
			ImageURL:    imageURL,
		})
	}

	return artists, nil
}

func artistCount(ctx context.Context, tx *sql.Tx) (int, error) {
	var count int

	err := tx.QueryRowContext(ctx, "SELECT COUNT(*) FROM artist").Scan(&count)
	if err != nil {
		return 0, err
	}

	return count, nil

}

func insertArtist(ctx context.Context, tx *sql.Tx, a Artist) (int64, error) {
	query := `
	INSERT INTO artist (name, description, image_url)
	VALUES (@name, @description, @image_url)`

	res, err := tx.ExecContext(ctx, query,
		sql.Named("name", a.Name),
		sql.Named("description", a.Description),
		sql.Named("image_url", a.ImageURL),
	)

	if err != nil {
		return 0, err
	}

	artistID, err := res.LastInsertId()
	if err != nil {
		return 0, err
	}

	return artistID, nil
}

func insertGenre(ctx context.Context, tx *sql.Tx, name string) (int64, error) {
	query := `INSERT OR IGNORE INTO genre (name) VALUES (@name)`

	res, err := tx.ExecContext(ctx, query, sql.Named("name", name))
	if err != nil {
		return 0, err
	}

	genreID, err := res.LastInsertId()
	if err != nil {
		return 0, err
	}

	return genreID, nil
}

func associateArtistWithSocial(ctx context.Context, tx *sql.Tx, artistID int64, socialID int64) error {
	query := `
	INSERT INTO artists_socials (artist_id, social_id) 
	VALUES (@artist_id, @social_id)`

	_, err := tx.ExecContext(ctx, query,
		sql.Named("artist_id", artistID),
		sql.Named("social_id", socialID),
	)

	if err != nil {
		return err
	}

	return nil
}

func associateArtistWithGenre(ctx context.Context, tx *sql.Tx, artistID int64, genreID int64) error {
	query := `
	INSERT INTO artists_genres (artist_id, genre_id) 
	VALUES (@artist_id, @genre_id)`

	_, err := tx.ExecContext(ctx, query,
		sql.Named("artist_id", artistID),
		sql.Named("genre_id", genreID),
	)

	if err != nil {
		return err
	}

	return nil
}

func insertSocial(ctx context.Context, tx *sql.Tx, url string) (int64, error) {
	query := `INSERT INTO social (url) VALUES (@url)`

	res, err := tx.ExecContext(ctx, query, sql.Named("url", url))
	if err != nil {
		return 0, err
	}

	socialID, err := res.LastInsertId()
	if err != nil {
		return 0, err
	}

	return socialID, nil
}

func artistByID(ctx context.Context, tx *sql.Tx, artistID int64) (Artist, error) {
	query := `
	SELECT name, description, image_url 
	FROM artist where id = @id`

	var name, description, imageURL string
	err := tx.QueryRowContext(ctx, query, sql.Named("id", artistID)).Scan(
		&name, &description, &imageURL,
	)

	if err != nil {
		return Artist{}, err
	}

	return Artist{
		ID:          artistID,
		Name:        name,
		Description: description,
		ImageURL:    imageURL,
	}, nil
}

func artistGenres(ctx context.Context, tx *sql.Tx, artistID int64) ([]Genre, error) {
	query := `
	SELECT id, name FROM genre g
	JOIN artists_genres ag ON ag.genre_id = g.id
	WHERE ag.artist_id = @artist_id`

	rows, err := tx.QueryContext(ctx, query, sql.Named("artist_id", artistID))
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	genres := make([]Genre, 0)

	for rows.Next() {
		var id int64
		var name string

		if err := rows.Scan(&id, &name); err != nil {
			return nil, err
		}

		genres = append(genres, Genre{ID: id, Name: name})
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return genres, nil
}

func artistSocials(ctx context.Context, tx *sql.Tx, artistID int64) ([]Social, error) {
	query := `
	SELECT id, url FROM social
	JOIN artists_socials ON artists_socials.social_id = social.id
	WHERE artists_socials.artist_id = @artist_id`

	rows, err := tx.QueryContext(ctx, query, sql.Named("artist_id", artistID))
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	socials := make([]Social, 0)

	for rows.Next() {
		var id int64
		var url string

		if err := rows.Scan(&id, &url); err != nil {
			return nil, err
		}

		socials = append(socials, Social{ID: id, URL: url})
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return socials, nil
}

func genreByID(ctx context.Context, tx *sql.Tx, genreID int64) (Genre, error) {
	query := `SELECT name FROM genre WHERE id = @genre_id`

	var name string
	err := tx.QueryRowContext(ctx, query,
		sql.Named("genre_id", genreID),
	).Scan(&name)

	if err != nil {
		return Genre{}, err
	}

	return Genre{
		ID:   genreID,
		Name: name,
	}, nil
}

func (a Artist) ToInternal(genres []artist.Genre, socials []artist.Social) artist.Artist {
	return artist.Artist{
		ID:          a.ID,
		Name:        a.Name,
		Description: a.Description,
		ImageURL:    a.ImageURL,
		Genres:      genres,
		Socials:     socials,
	}
}
