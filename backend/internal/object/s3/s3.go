package s3

import (
	"context"
	"errors"
	"fmt"
	"io"
	"path"
	"strings"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
	"github.com/mattismoel/konnekt/internal/object"
)

var (
	ErrInavlidRegion      = errors.New("S3 bucket region must be valid")
	ErrInvalidBucket      = errors.New("S3 bucket name must be valid")
	ErrInaccessibleBucket = errors.New("S3 bucket is inaccessible")
)

var DEFAULT_CACHE_CONTROL_MS = 2 * time.Hour.Milliseconds()

var _ object.Store = (*S3ObjectStore)(nil)

type S3ObjectStore struct {
	Bucket string
	Region string

	uploader   *s3manager.Uploader
	downloader *s3manager.Downloader
	client     *s3.S3
}

func NewS3ObjectStore(ctx context.Context, region string, bucket string) (*S3ObjectStore, error) {
	if strings.TrimSpace(region) == "" {
		return nil, ErrInavlidRegion
	}

	if strings.TrimSpace(bucket) == "" {
		return nil, ErrInvalidBucket
	}

	config := aws.
		NewConfig().
		WithRegion(region).
		WithCredentialsChainVerboseErrors(true)

	sess, err := session.NewSession(config)
	if err != nil {
		return nil, fmt.Errorf("Could not create AWS session: %v", err)
	}

	_, err = sess.Config.Credentials.Get()
	if err != nil {
		return nil, fmt.Errorf("Could not get AWS credentials: %v", err)
	}

	uploader := s3manager.NewUploader(sess)
	downloader := s3manager.NewDownloader(sess)
	client := s3.New(sess)

	_, err = client.HeadBucketWithContext(ctx, &s3.HeadBucketInput{
		Bucket: aws.String(bucket),
	})

	if err != nil {
		return nil, ErrInaccessibleBucket
	}

	return &S3ObjectStore{
		Region: region,
		Bucket: bucket,

		client:     client,
		uploader:   uploader,
		downloader: downloader,
	}, nil
}

func (s S3ObjectStore) Upload(ctx context.Context, key string, body io.Reader) (string, error) {
	output, err := s.uploader.UploadWithContext(ctx, &s3manager.UploadInput{
		Key:          aws.String(key),
		Bucket:       aws.String(s.Bucket),
		Body:         body,
		CacheControl: aws.String(fmt.Sprintf("Max-Age=%d", DEFAULT_CACHE_CONTROL_MS)),
	})

	if err != nil {
		return "", fmt.Errorf("Could not upload object %q: %v", key, err)
	}

	return output.Location, nil
}

func (s S3ObjectStore) Get(ctx context.Context, key string) (io.ReadCloser, error) {
	output, err := s.client.GetObjectWithContext(ctx, &s3.GetObjectInput{
		Key:    aws.String(key),
		Bucket: aws.String(s.Bucket),
	})

	if err != nil {
		return nil, fmt.Errorf("Could not get object %q: %v", key, err)
	}

	return output.Body, nil
}

func (s S3ObjectStore) Delete(ctx context.Context, key string) error {
	_, err := s.client.DeleteObjectWithContext(ctx, &s3.DeleteObjectInput{
		Key:    aws.String(key),
		Bucket: aws.String(s.Bucket),
	})

	if err != nil {
		return fmt.Errorf("Could not delete object %q: %v", key, err)
	}

	return nil
}

func (s S3ObjectStore) ObjectPath(key string) string {
	key = path.Clean(key)

	return fmt.Sprintf("https://%s.s3.%s.amazonaws.com/%s", s.Bucket, s.Region, key)
}
