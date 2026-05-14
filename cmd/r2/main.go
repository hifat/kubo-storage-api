package main

import (
	"context"
	"flag"
	"log"

	"github.com/aws/aws-sdk-go-v2/aws"
	awsconfig "github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/credentials"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/hifat/kubo-storage-api/config"
	storagedi "github.com/hifat/kubo-storage-api/internal/storage/di"
	"github.com/hifat/kubo-storage-api/pkg/logger"
	storageproto "github.com/hifat/kubo-storage-api/proto/storage"
	"gofr.dev/pkg/gofr"
)

func main() {
	envPath := flag.String("envPath", "", "env path")
	flag.Parse()

	cfg, err := config.LoadConfig(*envPath)
	if err != nil {
		log.Fatalf("failed to load config: %v", err)
	}

	app := gofr.New()
	log := app.Logger()
	gfrLog := logger.NewGofrLogger(log)

	awsCfg, err := awsconfig.LoadDefaultConfig(context.Background(),
		awsconfig.WithRegion("auto"),
		awsconfig.WithCredentialsProvider(
			credentials.NewStaticCredentialsProvider(
				cfg.Storage.AccessKey,
				cfg.Storage.SecretKey,
				"",
			),
		),
	)
	if err != nil {
		app.Logger().Fatalf("unable to load AWS SDK config: %v", err)
	}

	s3Client := s3.NewFromConfig(awsCfg, func(o *s3.Options) {
		o.BaseEndpoint = aws.String(cfg.Storage.DomainURL)
		o.UsePathStyle = true
	})

	storageDi := storagedi.Init(cfg, gfrLog, s3Client, s3.NewPresignClient(s3Client))

	storageproto.RegisterStorageServer(app, storageDi.GRPC)

	app.Run()
}
