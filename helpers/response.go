package helpers

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
)

// TODO: can add multiple fields
type ErrorStruct struct {
	Code    int    `json:"code,omitempty"`
	Message string `json:"message,omitempty"`
}

type ResponseBodyStruct struct {
	StatusCode int         `json:"statusCode,omitempty"`
	Status     string      `json:"status,omitempty"`
	Message    interface{} `json:"message,omitempty"`
	Err        ErrorStruct `json:"error,omitempty"`
	Data       interface{} `json:"data,omitempty"`
}

func ResponseBody(c *fiber.Ctx, statusCode int, message string, data interface{}, err error) error {
	rBody := ResponseBodyStruct{
		StatusCode: statusCode,
		Status:     http.StatusText(statusCode),
		Message:    message,
		Data:       data,
		Err: ErrorStruct{
			Message: err.Error(),
		},
	}

	err = c.Status(statusCode).JSON(rBody)
	if err != nil {
		return err
	}
	return nil
}
