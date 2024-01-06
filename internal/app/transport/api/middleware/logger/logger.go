package logger

import (
	"log/slog"
	"os"
	"time"

	"github.com/gofiber/fiber/v2"

	"layout/internal/pkg/constants"
	"layout/internal/pkg/ctxlogger"
)

func New() func(c *fiber.Ctx) error {
	return func(c *fiber.Ctx) error {
		ctx := c.Context()
		requestID := ctx.Value(constants.RequestIDCtxKey)

		log := slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{
			Level: slog.LevelDebug,
		}))
		log = log.WithGroup("context")
		log = log.With(slog.Any("request_id", requestID))
		c.Locals(constants.RequestIDCtxKey, log)

		return c.Next()
	}
}

func NewHandleRequestTimeMiddleware() func(c *fiber.Ctx) error {
	return func(c *fiber.Ctx) error {
		ctx := c.Context()
		log, err := ctxlogger.Logger(ctx)
		if err != nil {
			return err
		}

		t1 := time.Now()
		err = c.Next()

		log.Info("request",
			slog.Any("response_time", time.Since(t1).String()),
			slog.Any("request_url", c.OriginalURL()),
		)
		return err
	}
}
