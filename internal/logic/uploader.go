package logic

import (
	"context"
	"github.com/PlanVX/aweme/internal/config"
	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
	"io"
)

//go:generate mockery --name Uploader --inpackage --filename=mock_uploader_test.go --structname=MockUploader
type (
	// Uploader is the interface for uploading file
	Uploader interface {
		// Upload uploads a file, returns the URL of the uploaded file
		Upload(ctx context.Context, input UploadInput) (url string, err error)
	}

	// UploadInput is the input parameter for Uploader.Upload
	UploadInput struct {
		Key   string
		Value io.Reader
	}
	// Minio is the implementation of Uploader
	Minio struct {
		client     *minio.Client
		bucket     string
		domainName string
	}
)

// NewMinio returns a new *Minio
func NewMinio(c *config.Config) (*Minio, error) {
	client, err := minio.New(c.S3.Endpoint, &minio.Options{
		Creds:  credentials.NewStaticV4(c.S3.AccessKey, c.S3.SecretKey, ""),
		Secure: true,
	})
	return &Minio{
		client:     client,
		bucket:     c.S3.Bucket,
		domainName: c.S3.Endpoint,
	}, err
}

// Upload uploads a file, returns the URL of the uploaded file
func (m *Minio) Upload(ctx context.Context, input UploadInput) (string, error) {
	resp, err := m.client.PutObject(ctx, m.bucket, input.Key, input.Value, -1, minio.PutObjectOptions{})
	if err != nil {
		return "", err
	}
	return resp.Location, nil
}
