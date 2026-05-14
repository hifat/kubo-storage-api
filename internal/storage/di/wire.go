//go:build wireinject
// +build wireinject

package storagedi

import (
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/google/wire"
	"github.com/hifat/kubo-storage-api/config"
	storagehdl "github.com/hifat/kubo-storage-api/internal/storage/handler"
	storagerepo "github.com/hifat/kubo-storage-api/internal/storage/repository"
	storagesvc "github.com/hifat/kubo-storage-api/internal/storage/service"
	"github.com/hifat/kubo-storage-api/pkg/logger"
)

func Init(cfg *config.Config, log logger.Logger, client *s3.Client, presignClient *s3.PresignClient) *storagehdl.Handler {
	wire.Build(
		// Repository
		storagerepo.NewR2,

		// Service
		storagesvc.New,

		// Handler
		storagehdl.New,
		storagehdl.NewGRPC,
	)

	return &storagehdl.Handler{}
}
