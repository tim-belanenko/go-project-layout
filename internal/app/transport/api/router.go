package api

import (
	"context"
	"layout/internal/app/usecase"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/swagger"
)

type IHealthcheck interface {
	Status(ctx context.Context) (*usecase.HealthStatus, error)
}

type Router struct {
	hc IHealthcheck
}

func NewRouter(hc IHealthcheck) *Router {
	return &Router{
		hc: hc,
	}
}

func (r *Router) RegisterHandlers(app *fiber.App) {
	app.Get("/healthcheck", r.healthCheckHandler())

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
func (r *Router) healthCheckHandler() func(c *fiber.Ctx) error {
	return func(c *fiber.Ctx) error {
		status, err := r.hc.Status(c.Context())
		if err != nil {
			return err
		}
		return c.JSON(status)
	}
}
