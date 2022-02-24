package response

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
)

// TODO: can add multiple fields
type ErrorStruct struct {
	Code    int    `json:"code,omitempty"`
	Message string `json:"message,omitempty"`
	Stack   string `json:"stack,omitempty"`
}

type BodyStruct struct {
	StatusCode int         `json:"statusCode,omitempty"`
	Status     string      `json:"status,omitempty"`
	Message    interface{} `json:"message,omitempty"`
	Err        ErrorStruct `json:"error,omitempty"`
	Data       interface{} `json:"data,omitempty"`
}

// ResponseBody should be
func NewBody(c *fiber.Ctx, statusCode int, message string, data interface{}, err error) error {
	rBody := BodyStruct{
		StatusCode: statusCode,
		Status:     http.StatusText(statusCode),
		Message:    message,
		Data:       data,
		Err: ErrorStruct{
			Message: err.Error(),
		},
	}

	errHandler := c.App().ErrorHandler

	iErr := c.Status(statusCode).JSON(rBody)
	if err != nil {
		_ = errHandler(c, iErr)
	}
	return nil
}
