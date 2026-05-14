package logger

import (
	"context"

	"gofr.dev/pkg/gofr/logging"
)

type Logger interface {
	Info(ctx context.Context, args ...any)
	Error(ctx context.Context, args ...any)
	Debug(ctx context.Context, args ...any)
	Warn(ctx context.Context, args ...any)
}

type gofrLogger struct {
	log logging.Logger
}

func NewGofrLogger(log logging.Logger) Logger {
	return &gofrLogger{
		log: log,
	}
}

func (l *gofrLogger) Info(ctx context.Context, args ...any) {
	l.log.Info(ctx, args)
}

func (l *gofrLogger) Error(ctx context.Context, args ...any) {
	l.log.Error(ctx, args)
}

func (l *gofrLogger) Debug(ctx context.Context, args ...any) {
	l.log.Debug(ctx, args)
}

func (l *gofrLogger) Warn(ctx context.Context, args ...any) {
	l.log.Warn(ctx, args)
}
