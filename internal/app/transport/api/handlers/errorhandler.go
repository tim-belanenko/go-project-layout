package handlers

import (
	"errors"

	"github.com/gofiber/fiber/v2"

	"layout/internal/pkg/exterrors"
)

func NewErrorHandler() func(ctx *fiber.Ctx, err error) error {
	return func(ctx *fiber.Ctx, err error) error {
		if errors.Is(err, exterrors.ErrFailedGetLoggerFormContext) {
			return ctx.Status(fiber.StatusInternalServerError).SendString("Internal Server Error")
		}

		var e *fiber.Error
		if errors.As(err, &e) {
			return ctx.Status(e.Code).SendString(e.Message)
		}
		return nil
	}
}
