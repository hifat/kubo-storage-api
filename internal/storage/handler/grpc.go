package storagehdl

import (
	"context"

	"github.com/google/uuid"
	storagemdl "github.com/hifat/kubo-storage-api/internal/storage"
	storageproto "github.com/hifat/kubo-storage-api/proto/storage"
)

type GRPC struct {
	storageproto.UnimplementedStorageServer
	service storagemdl.Service
}

func NewGRPC(service storagemdl.Service) *GRPC {
	return &GRPC{
		service: service,
	}
}

func (h *GRPC) Upload(ctx context.Context, req *storageproto.UploadRequest) (*storageproto.UploadResponse, error) {
	uploadReq := &storagemdl.UploadRequest{
		ObjectKey:   uuid.New().String(),
		Body:        req.File,
		Filename:    req.Filename,
		ContentType: req.ContentType,
		Path:        req.Path,
	}

	res, err := h.service.Upload(ctx, uploadReq)
	if err != nil {
		return &storageproto.UploadResponse{
			Error: err.Error(),
		}, nil
	}

	return &storageproto.UploadResponse{
		Filename:  uploadReq.Filename,
		ObjectKey: res.ObjectKey,
		Url:       res.URL,
	}, nil
}

func (h *GRPC) GetPresignedURL(ctx context.Context, req *storageproto.GetPresignedURLRequest) (*storageproto.GetPresignedURLResponse, error) {
	res, err := h.service.GetPresignedURL(ctx, req.ObjectKey)
	if err != nil {
		return &storageproto.GetPresignedURLResponse{
			Error: err.Error(),
		}, nil
	}

	return &storageproto.GetPresignedURLResponse{
		Url: res.URL,
	}, nil
}

func (h *GRPC) Delete(ctx context.Context, req *storageproto.DeleteRequest) (*storageproto.DeleteResponse, error) {
	err := h.service.Delete(ctx, req.ObjectKey)
	if err != nil {
		return &storageproto.DeleteResponse{
			Error: err.Error(),
		}, nil
	}

	return &storageproto.DeleteResponse{
		Message: "File deleted successfully",
	}, nil
}
