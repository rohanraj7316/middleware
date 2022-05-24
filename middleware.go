package middleware

import (
	"context"
	"sync"

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
	var (
		cfg  Config
		once sync.Once
	)

	// Set up middleware layer once
	// multiple declaration won't update these changes
	once.Do(func() {
		cfg = configDefault(config...)

		// tagging all the request
		app.Use(requestid.New(cfg.requestIdConfig))

		// adding request and response loggers
		app.Use(logger.New(cfg.reqResLogger))
	})

	return func(c *fiber.Ctx) error {
		// setting up relayback headers.
		for i := 0; i < len(cfg.relaybackHeader); i++ {
			c.Vary(cfg.relaybackHeader[i])
		}

		for i := 0; i < len(cfg.relaybackHeader); i++ {
			if key, val := cfg.relaybackHeader[i], c.Get(cfg.relaybackHeader[i]); val != "" {
				rCtx := context.WithValue(c.UserContext(), key, val)
				c.SetUserContext(rCtx)
			}
		}

		return c.Next()
	}
}
