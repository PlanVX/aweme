package logic

import (
	conf "github.com/PlanVX/aweme/internal/config"
	"github.com/aws/aws-sdk-go/aws/session"
	"io"

	// import the S3 client package
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
)

//go:generate mockery --name Uploader --inpackage --filename=mock_uploader_test.go --structname=MockUploader
type (
	// Uploader is the interface for uploading file
	Uploader interface {
		// Upload uploads a file, returns the URL of the uploaded file
		Upload(input UploadInput) (url string, err error)
	}

	// UploadInput is the input parameter for Uploader.Upload
	UploadInput struct {
		Key   string
		Value io.Reader
	}
	// S3 is the implementation of Uploader
	// based on AWS S3
	S3 struct {
		config *aws.Config
		bucket string
	}
)

// NewS3 creates a new S3 based Uploader
func NewS3(c *conf.Config) *S3 {
	return &S3{
		config: &aws.Config{
			Credentials:      credentials.NewStaticCredentials(c.S3.AccessKey, c.S3.SecretKey, ""),
			Endpoint:         aws.String(c.S3.Endpoint),
			Region:           aws.String(c.S3.Region),
			S3ForcePathStyle: aws.Bool(false),
		},
		bucket: c.S3.Bucket,
	}
}

// newUploader creates a new S3 Uploader
func (s3 *S3) newUploader() (*s3manager.Uploader, error) {
	newSession, err := session.NewSession(s3.config)
	if err != nil {
		return nil, err
	}
	return s3manager.NewUploader(newSession), nil
}

// Upload uploads a file to S3
// Returns the URL of the uploaded file
func (s3 *S3) Upload(input UploadInput) (string, error) {
	uploader, err := s3.newUploader()
	if err != nil {
		return "", err
	}
	r, err := uploader.Upload(&s3manager.UploadInput{
		Bucket: aws.String(s3.bucket),
		Key:    aws.String(input.Key),
		Body:   input.Value,
	})
	if err != nil {
		return "", err
	}
	return r.Location, nil
}
