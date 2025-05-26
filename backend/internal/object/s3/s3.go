package s3

import (
	"context"
	"fmt"
	"io"
	"path"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
	"github.com/mattismoel/konnekt/internal/object"
)

var DEFAULT_CACHE_CONTROL_MS = 2 * time.Hour.Milliseconds()

var _ object.Store = (*S3ObjectStore)(nil)

type S3ObjectStore struct {
	bucket     string
	region     string
	uploader   *s3manager.Uploader
	downloader *s3manager.Downloader
	client     *s3.S3
}

func NewS3ObjectStore(region string, bucket string) (*S3ObjectStore, error) {
	config := aws.NewConfig().
		WithRegion(region)

	sess, err := session.NewSession(config)
	if err != nil {
		return nil, err
	}

	return &S3ObjectStore{
		region: region,
		bucket: bucket,

		client:     s3.New(sess),
		uploader:   s3manager.NewUploader(sess),
		downloader: s3manager.NewDownloader(sess),
	}, nil
}

func (s S3ObjectStore) Upload(ctx context.Context, key string, body io.Reader) (string, error) {
	output, err := s.uploader.UploadWithContext(ctx, &s3manager.UploadInput{
		Key:          aws.String(key),
		Bucket:       aws.String(s.bucket),
		Body:         body,
		CacheControl: aws.String(fmt.Sprintf("Max-Age=%d", DEFAULT_CACHE_CONTROL_MS)),
	})

	if err != nil {
		return "", err
	}

	return output.Location, nil
}

func (s S3ObjectStore) Get(ctx context.Context, key string) (io.ReadCloser, error) {
	output, err := s.client.GetObjectWithContext(ctx, &s3.GetObjectInput{
		Key:    aws.String(key),
		Bucket: aws.String(s.bucket),
	})

	if err != nil {
		return nil, err
	}

	return output.Body, nil
}

func (s S3ObjectStore) Delete(ctx context.Context, key string) error {
	_, err := s.client.DeleteObjectWithContext(ctx, &s3.DeleteObjectInput{
		Key:    aws.String(key),
		Bucket: aws.String(s.bucket),
	})

	if err != nil {
		return err
	}

	return nil
}

func (s S3ObjectStore) ObjectPath(key string) string {
	key = path.Clean(key)

	return fmt.Sprintf("https://%s.s3.%s.amazonaws.com/%s", s.bucket, s.region, key)
}
