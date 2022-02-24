package configs

import (
	"net/http"

	"github.com/rohanraj7316/middleware/helpers"

	"github.com/gofiber/fiber/v2"
)

var ServerDefault = fiber.Config{
	// ErrorHandler is executed when an error is returned from fiber.Handler.
	//
	// Default: below func will be get executed if you don't override it.
	ErrorHandler: func(c *fiber.Ctx, e error) error {
		if e != nil {
			err := helpers.ResponseBody(c, http.StatusInternalServerError, "unable to process your query", nil, e)
			if err != nil {
				// TODO: add error logging
			}
			// TODO: add error logging
		}
		return nil
	},
	// Max body size that the server accepts.
	// -1 will decline any body size
	//
	// Default: 5 * 1024 * 1024 = 5mb
	BodyLimit: 5 * 1024 * 1024,
}

var ServerStaticDefault = fiber.Static{}
