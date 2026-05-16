package middlewarehdl

import (
	"context"

	middlewaremld "github.com/hifat/kubo-storage-api/internal/middleware"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
)

type GRPC struct {
	middlewareService middlewaremld.Service
}

func NewGRPC(middlewareService middlewaremld.Service) *GRPC {
	return &GRPC{
		middlewareService: middlewareService,
	}
}

func (h *GRPC) UnaryAuthInterceptor() grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
		md, ok := metadata.FromIncomingContext(ctx)
		if !ok {
			return nil, status.Errorf(codes.Unauthenticated, "missing metadata")
		}

		token := md["x-api-key"]
		if len(token) == 0 {
			return nil, status.Errorf(codes.Unauthenticated, "missing token")
		}

		if err := h.middlewareService.ValidateToken(token[0]); err != nil {
			return nil, status.Errorf(codes.Unauthenticated, "invalid token: %v", err)
		}

		return handler(ctx, req)
	}
}
