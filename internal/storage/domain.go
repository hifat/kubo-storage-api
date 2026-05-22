package storagemdl

import (
	"context"
)

type Repository interface {
	Upload(ctx context.Context, req *UploadRequest) (string, error)
	GetPresignedURL(ctx context.Context, objectKey string) (string, error)
	Delete(ctx context.Context, objectKey string) error
}

type Service interface {
	Upload(ctx context.Context, req *UploadRequest) (*UploadResponse, error)
	GetPresignedURL(ctx context.Context, objectKey string) (*PresignedResponse, error)
	Delete(ctx context.Context, objectKey string) error
}
