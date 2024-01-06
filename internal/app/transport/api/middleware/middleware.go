package middleware

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/gofiber/fiber/v2/middleware/requestid"

	"layout/internal/app/transport/api/middleware/logger"
	"layout/internal/pkg/constants"
)

func RegisterMiddlewares(app *fiber.App) {
	app.Use(recover.New())
	app.Use(requestid.New(requestid.Config{
		ContextKey: constants.RequestIDCtxKey,
	}))

	app.Use(logger.New())
	app.Use(logger.NewHandleRequestTimeMiddleware())

	app.Use(cors.New(cors.Config{
		AllowOrigins: "*",
		AllowHeaders: "Origin, Content-Type, Accept",
	}),
	)
}
