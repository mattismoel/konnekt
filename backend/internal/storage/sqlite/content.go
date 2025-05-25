package sqlite

import (
	"context"
	"database/sql"

	sq "github.com/Masterminds/squirrel"
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
	query, args, err := sq.
		Insert("landing_image").
		Columns("url").
		Values(url).
		ToSql()

	if err != nil {
		return 0, err
	}

	res, err := tx.ExecContext(ctx, query, args...)
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
	query, args, err := sq.
		Select("id", "url").
		From("landing_image").
		ToSql()

	if err != nil {
		return nil, err
	}

	rows, err := tx.QueryContext(ctx, query, args...)
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
	query, args, err := sq.
		Select("url").
		From("landing_image").
		Where(sq.Eq{"id": id}).
		ToSql()

	if err != nil {
		return Image{}, err
	}

	var url string

	err = tx.
		QueryRowContext(ctx, query, args...).
		Scan(&url)

	if err != nil {
		return Image{}, err
	}

	return Image{
		ID:  id,
		URL: url,
	}, nil
}

func deleteLandingImage(ctx context.Context, tx *sql.Tx, id int64) error {
	query, args, err := sq.
		Delete("landing_image").
		Where(sq.Eq{"id": id}).
		ToSql()

	if err != nil {
		return err
	}

	_, err = tx.ExecContext(ctx, query, args...)
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
