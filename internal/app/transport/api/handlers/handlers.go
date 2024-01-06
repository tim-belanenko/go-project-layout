package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/swagger"
)

func RegisterHandlers(app *fiber.App) {
	app.Get("/healthcheck", healthCheckHandler())

	app.Get("/swagger/*", swagger.HandlerDefault) // default
}

// Healthcheck api status
//
//	@Summary      API healthcheck
//	@Description  healthcheck
//	@Tags         sys
//	@Accept       json
//	@Produce      json
//	@Success      200
//	@Router       /healthcheck [get]
func healthCheckHandler() func(c *fiber.Ctx) error {
	return func(c *fiber.Ctx) error {
		return c.SendStatus(fiber.StatusOK)
	}
}
