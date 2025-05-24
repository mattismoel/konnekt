package sqlite

import (
	"context"
	"database/sql"

	"github.com/mattismoel/konnekt/internal/domain/content"
)

var _ content.Repository = (*ContentRepository)(nil)

type Image struct {
	ID  int64
	URL string
}

type ImageCollection = []Image

type ContentRepository struct {
	db *sql.DB
}

func NewContentRepository(db *sql.DB) (*ContentRepository, error) {
	return &ContentRepository{
		db: db,
	}, nil
}

// InsertLandingImage implements content.Repository.
func (r *ContentRepository) InsertLandingImage(ctx context.Context, url string) (int64, error) {
	tx, err := r.db.BeginTx(ctx, nil)
	if err != nil {
		return 0, err
	}

	defer tx.Rollback()

	id, err := insertLandingImage(ctx, tx, url)
	if err != nil {
		return 0, err
	}

	if err := tx.Commit(); err != nil {
		return 0, err
	}

	return id, nil
}

// LandingImageByID implements content.Repository.
func (r ContentRepository) LandingImageByID(ctx context.Context, id int64) (content.LandingImage, error) {
	tx, err := r.db.BeginTx(ctx, nil)
	if err != nil {
		return content.LandingImage{}, err
	}

	defer tx.Rollback()

	dbImg, err := landingImageByID(ctx, tx, id)
	if err != nil {
		return content.LandingImage{}, err
	}

	if err := tx.Commit(); err != nil {
		return content.LandingImage{}, err
	}

	return dbImg.ToInternal(), nil
}

// LandingImages implements content.Repository.
func (r *ContentRepository) LandingImages(ctx context.Context) ([]content.LandingImage, error) {
	tx, err := r.db.BeginTx(ctx, nil)
	if err != nil {
		return nil, err
	}

	defer tx.Rollback()

	dbImages, err := landingImages(ctx, tx)
	if err != nil {
		return nil, err
	}

	if err := tx.Commit(); err != nil {
		return nil, err
	}

	images := make([]content.LandingImage, 0)

	for _, img := range dbImages {
		images = append(images, img.ToInternal())
	}

	return images, nil
}

// DeleteLandingImage implements content.Repository.
func (r ContentRepository) DeleteLandingImage(ctx context.Context, id int64) error {
	tx, err := r.db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}

	defer tx.Rollback()

	if err := deleteLandingImage(ctx, tx, id); err != nil {
		return err
	}

	if err := tx.Commit(); err != nil {
		return err
	}

	return nil
}

func insertLandingImage(ctx context.Context, tx *sql.Tx, url string) (int64, error) {
	query := "INSERT INTO landing_image (url) VALUES (@url)"

	res, err := tx.ExecContext(ctx, query, sql.Named("url", url))
	if err != nil {
		return 0, err
	}

	id, err := res.LastInsertId()
	if err != nil {
		return 0, err
	}

	return id, nil
}

func landingImages(ctx context.Context, tx *sql.Tx) (ImageCollection, error) {
	query := "SELECT id, url FROM landing_image"

	rows, err := tx.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	images := make(ImageCollection, 0)

	for rows.Next() {
		var id int64
		var url string

		if err := rows.Scan(&id, &url); err != nil {
			return nil, err
		}

		images = append(images, Image{
			ID:  id,
			URL: url,
		})
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return images, nil
}

func landingImageByID(ctx context.Context, tx *sql.Tx, id int64) (Image, error) {
	query := "SELECT url FROM landing_image WHERE id = @id"

	var url string
	err := tx.QueryRowContext(ctx, query, sql.Named("id", id)).Scan(&url)
	if err != nil {
		return Image{}, err
	}

	return Image{
		ID:  id,
		URL: url,
	}, nil
}

func deleteLandingImage(ctx context.Context, tx *sql.Tx, id int64) error {
	query := "DELETE FROM landing_image WHERE id = @id"

	_, err := tx.ExecContext(ctx, query, sql.Named("id", id))
	if err != nil {
		return err
	}

	return nil
}

func (img Image) ToInternal() content.LandingImage {
	return content.LandingImage{
		ID:  img.ID,
		URL: img.URL,
	}
}
