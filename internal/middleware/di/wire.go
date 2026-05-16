//go:build wireinject
// +build wireinject

package middlewaredi

import (
	"github.com/google/wire"
	"github.com/hifat/kubo-storage-api/config"
	middlewarehdl "github.com/hifat/kubo-storage-api/internal/middleware/handler"
	middlewaresvc "github.com/hifat/kubo-storage-api/internal/middleware/service"
	"github.com/hifat/kubo-storage-api/pkg/logger"
)

func Init(cfg *config.Config, log logger.Logger) *middlewarehdl.Handler {
	wire.Build(

		// Service
		middlewaresvc.New,

		// Handler
		middlewarehdl.New,
		middlewarehdl.NewGRPC,
	)

	return &middlewarehdl.Handler{}
}
