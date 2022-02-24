package configs

import (
	"freecharge/rsrc-bp/api/middleware/configs/helpers"
	"net/http"

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
}
