package middlewaresvc

import (
	"fmt"

	"github.com/hifat/kubo-storage-api/config"
	middlewaremld "github.com/hifat/kubo-storage-api/internal/middleware"
	"github.com/hifat/kubo-storage-api/pkg/logger"
)

type middlewareService struct {
	cfg *config.Config
	log logger.Logger
}

func New(cfg *config.Config, log logger.Logger) middlewaremld.Service {
	return &middlewareService{
		cfg: cfg,
		log: log,
	}
}

func (s *middlewareService) ValidateToken(token string) error {
	if token == s.cfg.App.ApiKey {
		return nil
	}

	return fmt.Errorf("invalid token")
}
