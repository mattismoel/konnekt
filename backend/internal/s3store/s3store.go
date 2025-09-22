package s3store

import (
	"context"
	"fmt"
	"io"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/feature/s3/manager"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/mattismoel/konnekt/backend/internal/server"
)

var _ server.ObjectStore = ObjectStore{}

type ObjectStore struct {
	Bucket string
	Region string

	s3Client *s3.Client
	uploader *manager.Uploader
}

func (o *ObjectStore) Initialise(ctx context.Context) error {
	cfg, err := config.LoadDefaultConfig(ctx)
	if err != nil {
		return fmt.Errorf("Could not load default config: %v", err)
	}

	s3Client := s3.NewFromConfig(cfg)
	uploader := manager.NewUploader(s3Client)

	_, err = s3Client.HeadBucket(ctx, &s3.HeadBucketInput{
		Bucket: aws.String(o.Bucket),
	})

	if err != nil {
		return fmt.Errorf("Could not HEAD bucket: %v", err)
	}

	o.s3Client = s3Client
	o.uploader = uploader

	return nil
}

// Insert implements server.ObjectStore.
func (o ObjectStore) Insert(ctx context.Context, key string, r io.Reader) (string, error) {
	output, err := o.uploader.Upload(ctx, &s3.PutObjectInput{
		Bucket: aws.String(o.Bucket),
		Key:    aws.String(key),
		Body:   r,
	})

	if err != nil {
		return "", fmt.Errorf("Could not upload file to %q: %v", key, err)
	}

	return output.Location, nil
}

// Delete implements server.ObjectStore.
func (o ObjectStore) Delete(ctx context.Context, key string) error {
	_, err := o.s3Client.DeleteObject(ctx, &s3.DeleteObjectInput{
		Bucket: aws.String(o.Bucket),
		Key:    aws.String(key),
	})

	if err != nil {
		return fmt.Errorf("Could not delete file %q: %v", key, err)
	}

	return nil
}
