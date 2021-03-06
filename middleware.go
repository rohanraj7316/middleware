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

		// setting up headers
		reqHeaders := c.GetReqHeaders()
		for key, val := range reqHeaders {
			if _, ok := cfg.passOnHeader[key]; ok {
				rCtx := context.WithValue(c.UserContext(), key, val)
				c.SetUserContext(rCtx)
			}

			if _, ok := cfg.relaybackHeader[key]; ok {
				c.Set(key, val)
			}
		}

		return c.Next()
	}
}
