package httpclient

import (
	"context"
	"io"

	"github.com/gofiber/fiber/v2"
)

// Option used for Request func
type Option struct {
	Ctx         *fiber.Ctx
	Method      string
	Url         string
	Header      map[string]string
	RequestBody io.Reader
}

// OptionSDK used for RequestSDK func
type OptionSDK struct {
	Ctx         context.Context
	Method      string
	Url         string
	Header      map[string]string
	RequestBody io.Reader
}
