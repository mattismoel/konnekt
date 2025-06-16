package service

import (
	"context"
	"fmt"
	"image"
	"io"
	"path"

	"github.com/mattismoel/konnekt/internal/domain/content"
	"github.com/mattismoel/konnekt/internal/object"
	"github.com/nfnt/resize"
)

const LANDING_IMAGE_WIDTH_PX = 2048

type ContentService struct {
	store       object.Store
	contentRepo content.Repository
}

func NewContentService(store object.Store, contentRepo content.Repository) *ContentService {
	return &ContentService{
		store:       store,
		contentRepo: contentRepo,
	}
}

func (s ContentService) LandingImages(ctx context.Context) ([]content.LandingImage, error) {
	images, err := s.contentRepo.LandingImages(ctx)
	if err != nil {
		return nil, fmt.Errorf("Could not get landing images: %v", err)
	}

	return images, nil
}

func (s ContentService) UploadLandingImage(ctx context.Context, r io.Reader) (int64, error) {
	img, _, err := image.Decode(r)
	if err != nil {
		return 0, fmt.Errorf("Could not decode landing image: %v", err)
	}

	if img.Bounds().Max.X > LANDING_IMAGE_WIDTH_PX {
		img = resize.Resize(LANDING_IMAGE_WIDTH_PX, 0, img, resize.Lanczos2)
	}

	formatedImg, err := formatJPEG(img)
	if err != nil {
		return 0, fmt.Errorf("Could not format landing image: %v", err)
	}

	fileName := createRandomImageFileName("jpeg")

	url, err := s.store.Upload(ctx, path.Join("/landing_images", fileName), formatedImg)
	if err != nil {
		return 0, fmt.Errorf("Could not upload landing image: %v", err)
	}

	id, err := s.contentRepo.InsertLandingImage(ctx, url)
	if err != nil {
		return 0, fmt.Errorf("Could not insert landing image into repository: %v", err)
	}

	return id, nil
}

func (s ContentService) LandingImageByID(ctx context.Context, id int64) (content.LandingImage, error) {
	img, err := s.contentRepo.LandingImageByID(ctx, id)
	if err != nil {
		return content.LandingImage{}, fmt.Errorf("Could not get landing image %d: %v", id, err)
	}

	return img, nil
}

func (s ContentService) DeleteLandingImage(ctx context.Context, id int64) error {
	img, err := s.contentRepo.LandingImageByID(ctx, id)
	if err != nil {
		return fmt.Errorf("Could not get landing image %d: %v", id, err)
	}

	if err := s.store.Delete(ctx, img.URL); err != nil {
		return fmt.Errorf("Could not delete image from object store: %v", err)
	}

	if err := s.contentRepo.DeleteLandingImage(ctx, int64(id)); err != nil {
		return fmt.Errorf("Could not delete landing image from repository: %v", err)
	}

	return nil
}
