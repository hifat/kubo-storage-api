package storagerepo

import (
	"bytes"
	"context"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/hifat/kubo-storage-api/config"
	storagemdl "github.com/hifat/kubo-storage-api/internal/storage"
)

type r2Repository struct {
	cfg           *config.Config
	client        *s3.Client
	presignClient *s3.PresignClient
}

func NewR2(cfg *config.Config, s3Client *s3.Client, presignClient *s3.PresignClient) storagemdl.Repository {
	return &r2Repository{
		cfg:           cfg,
		client:        s3Client,
		presignClient: presignClient,
	}
}

func (r *r2Repository) Upload(ctx context.Context, req *storagemdl.UploadRequest) (string, error) {
	_, err := r.client.PutObject(ctx, &s3.PutObjectInput{
		Bucket:      aws.String(r.cfg.Storage.BucketName),
		Key:         aws.String(req.ObjectKey),
		ContentType: aws.String(req.ContentType),
		Body:        bytes.NewReader(req.Body),
	})
	if err != nil {
		return "", err
	}

	presigned, err := r.presignClient.PresignGetObject(ctx, &s3.GetObjectInput{
		Bucket: aws.String(r.cfg.Storage.BucketName),
		Key:    aws.String(req.ObjectKey),
	}, func(opts *s3.PresignOptions) {
		opts.Expires = time.Duration(60 * int64(time.Second))
	})

	return presigned.URL, err
}

func (r *r2Repository) GetPresignedURL(ctx context.Context, objectKey string) (string, error) {
	presigned, err := r.presignClient.PresignGetObject(ctx, &s3.GetObjectInput{
		Bucket: aws.String(r.cfg.Storage.BucketName),
		Key:    aws.String(objectKey),
	}, func(opts *s3.PresignOptions) {
		opts.Expires = time.Duration(60 * int64(time.Second))
	})

	return presigned.URL, err
}

func (r *r2Repository) Delete(ctx context.Context, objectKey string) error {
	_, err := r.client.DeleteObject(ctx, &s3.DeleteObjectInput{
		Bucket: aws.String(r.cfg.Storage.BucketName),
		Key:    aws.String(objectKey),
	})

	return err
}
