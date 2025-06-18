package s3_test

import (
	"context"
	"errors"
	"testing"
	"time"

	"github.com/mattismoel/konnekt/internal/object/s3"
)

func TestNewStore(t *testing.T) {
	type test struct {
		region string
		bucket string

		wantRegion string
		wantBucket string

		wantErr error
	}

	tests := map[string]test{
		"Valid input": {
			region:     "eu-north-1",
			wantRegion: "eu-north-1",
			bucket:     "konnekt-bucket",
			wantBucket: "konnekt-bucket",
			wantErr:    nil,
		},
		"Empty region": {
			region:     "",
			wantRegion: "",
			bucket:     "konnekt-bucket",
			wantBucket: "konnekt-bucket",
			wantErr:    s3.ErrInavlidRegion,
		},
		"Empty bucket": {
			region:     "eu-north-1",
			wantRegion: "eu-north-1",
			bucket:     "",
			wantBucket: "",
			wantErr:    s3.ErrInvalidBucket,
		},
		"Inaccessible bucket": {
			region:     "eu-north-404",
			wantRegion: "",
			bucket:     "non-existant-bucket",
			wantBucket: "",
			wantErr:    s3.ErrInaccessibleBucket,
		},
	}

	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
			defer cancel()

			s, err := s3.NewS3ObjectStore(ctx, tt.region, tt.bucket)
			if !errors.Is(err, tt.wantErr) {
				t.Fatalf("got %v, want %v", err, tt.wantErr)
			}

			if err == nil && s.Bucket != tt.wantBucket {
				t.Fatalf("got bucket %q, want %q", s.Bucket, tt.wantBucket)
			}

			if err == nil && s.Region != tt.wantRegion {
				t.Fatalf("got region %q, want %q", s.Region, tt.wantRegion)
			}
		})
	}
}
