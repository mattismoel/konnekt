package sqlite

import (
	"context"
	"database/sql"
	"fmt"

	sq "github.com/Masterminds/squirrel"
	"github.com/mattismoel/konnekt/internal/domain/content"
)

var _ content.Repository = (*ContentRepository)(nil)

type Image struct {
	ID  int64
	URL string

	CreatedAt UnixTime
	CreatedBy int64
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
func (r *ContentRepository) InsertLandingImage(ctx context.Context, url string, createdByID int64) (int64, error) {
	tx, err := r.db.BeginTx(ctx, nil)
	if err != nil {
		return 0, err
	}

	defer tx.Rollback()

	id, err := insertLandingImage(ctx, tx, url, createdByID)
	if err != nil {
		return 0, fmt.Errorf("Could not insert landing image: %v", err)
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
		return content.LandingImage{}, fmt.Errorf("Could not find landing image with id %d: %v", id, err)
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
		return nil, fmt.Errorf("Could not list landing images: %v", err)
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
		return fmt.Errorf("Could not delete landing image with id %d: %v", id, err)
	}

	if err := tx.Commit(); err != nil {
		return err
	}

	return nil
}

func insertLandingImage(ctx context.Context, tx *sql.Tx, url string, createdByID int64) (int64, error) {
	query, args, err := sq.
		Insert("landing_image").
		Columns("url", "created_by").
		Values(url, createdByID).
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
		Select("id", "url", "created_at", "created_by").
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
		var createdAt UnixTime
		var createdByID int64

		if err := rows.Scan(&id, &url, &createdAt, &createdByID); err != nil {
			return nil, err
		}

		images = append(images, Image{
			ID:        id,
			URL:       url,
			CreatedAt: createdAt,
			CreatedBy: createdByID,
		})
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return images, nil
}

func landingImageByID(ctx context.Context, tx *sql.Tx, id int64) (Image, error) {
	query, args, err := sq.
		Select("url", "createdBy", "createdAt").
		From("landing_image").
		Where(sq.Eq{"id": id}).
		ToSql()

	if err != nil {
		return Image{}, err
	}

	var url string
	var createdByID int64
	var createdAtUnix UnixTime

	err = tx.
		QueryRowContext(ctx, query, args...).
		Scan(&url, &createdByID, &createdAtUnix)

	if err != nil {
		return Image{}, err
	}

	fmt.Printf("createdBY: %d, createdAt: %v\n", createdByID, createdAtUnix)

	return Image{
		ID:        id,
		URL:       url,
		CreatedBy: createdByID,
		CreatedAt: createdAtUnix,
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
		ID:        img.ID,
		URL:       img.URL,
		CreatedAt: img.CreatedAt.Time(),
		CreatedBy: img.CreatedBy,
	}
}
