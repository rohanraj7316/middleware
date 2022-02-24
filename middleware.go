package middleware

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/requestid"
)

// adding pre requisite middlewares in the existing
// router stack.
// Default: tagging all the request.
// 			request and response logging.
// 			request timeout.
func New(app *fiber.App, config ...Config) fiber.Handler {
	cfg := configDefault(config...)

	// tagging all the request
	app.Use(requestid.New(cfg.requestIdConfig))
	// adding request and response loggers
	app.Use(logger.New(cfg.reqResLogger))

	return func(c *fiber.Ctx) error {
		return fiber.ErrBadGateway
		// return c.Next()
	}
}
