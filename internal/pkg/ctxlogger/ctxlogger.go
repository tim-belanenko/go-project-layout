package ctxlogger

import (
	"context"
	"log/slog"
	"os"

	"layout/internal/pkg/constants"
	"layout/internal/pkg/exterrors"
)

var defaultLogger *CtxLogger

func init() {
	defaultLogger = &CtxLogger{
		slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{
			Level: slog.LevelDebug,
		}))}

}

type CtxLogger struct {
	*slog.Logger
}

func DefaultLogger() *CtxLogger {
	return defaultLogger
}

func Logger(ctx context.Context) (*CtxLogger, error) {
	logger, ok := ctx.Value(constants.RequestIDCtxKey).(*slog.Logger)
	if !ok {
		return nil, exterrors.ErrFailedGetLoggerFormContext
	}

	return &CtxLogger{logger}, nil
}

func (l *CtxLogger) Fatal(msg string, args ...any) {
	defaultLogger.Error(msg, args...)
	os.Exit(1)
}
