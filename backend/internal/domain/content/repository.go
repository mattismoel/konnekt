package content

import "context"

type Repository interface {
	LandingImages(ctx context.Context) ([]LandingImage, error)
	LandingImageByID(ctx context.Context, id int64) (LandingImage, error)
	InsertLandingImage(ctx context.Context, url string) (int64, error)
	DeleteLandingImage(ctx context.Context, id int64) error
}
