package main

import (
	"log"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/hifat/kubo-storage-api/config"
	storagedi "github.com/hifat/kubo-storage-api/internal/storage/di"
	"github.com/hifat/kubo-storage-api/pkg/logger"
	storageproto "github.com/hifat/kubo-storage-api/proto/storage"
	"gofr.dev/pkg/gofr"
)

func main() {
	cfg, err := config.LoadConfig("./env")
	if err != nil {
		log.Fatalf("can't load config .env: %v", err)
	}

	app := gofr.New()

	log := app.Logger()
	gfrLog := logger.NewGofrLogger(log)

	awsCfg := aws.Config{}
	s3Client := s3.NewFromConfig(awsCfg)

	storageDi := storagedi.Init(cfg, gfrLog, s3Client, s3.NewPresignClient(s3Client))

	storageproto.RegisterStorageServer(app, storageDi.GRPC)

	app.Run()
}
