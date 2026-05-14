package storagesvc

import (
	"context"

	storagemdl "github.com/hifat/kubo-storage-api/internal/storage"
	"github.com/hifat/kubo-storage-api/pkg/logger"
)

type service struct {
	log         logger.Logger
	storageRepo storagemdl.Repository
}

func New(log logger.Logger, storageRepo storagemdl.Repository) storagemdl.Service {
	return &service{
		log:         log,
		storageRepo: storageRepo,
	}
}

func (s *service) Upload(ctx context.Context, req *storagemdl.UploadRequest) (string, error) {
	url, err := s.storageRepo.Upload(ctx, req)
	if err != nil {
		s.log.Error(ctx, err)
		return "", err
	}

	return url, nil
}

func (s *service) GetPresignedURL(ctx context.Context, objectKey string) (*storagemdl.PresignedResponse, error) {
	url, err := s.storageRepo.GetPresignedURL(ctx, objectKey)
	if err != nil {
		s.log.Error(ctx, err)
		return nil, err
	}

	return &storagemdl.PresignedResponse{
		URL: url,
	}, nil
}

func (s *service) Delete(ctx context.Context, objectKey string) error {
	err := s.storageRepo.Delete(ctx, objectKey)
	if err != nil {
		s.log.Error(ctx, err)
		return err
	}

	return nil
}
