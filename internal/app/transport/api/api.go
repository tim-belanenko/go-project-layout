package api

import (
	"fmt"

	"github.com/gofiber/fiber/v2"

	"layout/api"
	"layout/internal/app/config"
	"layout/internal/app/transport/api/handlers"
	"layout/internal/app/transport/api/middleware"
)

// @title Fiber Example API
// @version 1.0
// @description This is a sample swagger for Fiber
// @termsOfService http://swagger.io/terms/
// @contact.name API Support
// @contact.email fiber@swagger.io
// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html
// @host localhost:8080
// @BasePath /
type APIServer struct {
	app *fiber.App
}

func New(cfg *config.HTTP, hc IHealthcheck) *APIServer {
	app := fiber.New(
		fiber.Config{
			ErrorHandler: handlers.NewErrorHandler(),
		},
	)
	api.SwaggerInfo.Host = fmt.Sprintf("%s:%d", cfg.ListenHost, cfg.ListenPort)

	middleware.RegisterMiddlewares(app)

	r := NewRouter(hc)
	r.RegisterHandlers(app)

	return &APIServer{
		app: app,
	}
}

func (s *APIServer) Listen(addr string) error {
	return s.app.Listen(addr)
}
