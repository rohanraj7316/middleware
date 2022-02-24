package helpers

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

type ResponseBodyStruct struct {
	StatusCode int         `json:"statusCode,omitempty"`
	Status     string      `json:"status,omitempty"`
	Message    interface{} `json:"message,omitempty"`
	Err        ErrorStruct `json:"error,omitempty"`
	Data       interface{} `json:"data,omitempty"`
}

// ResponseBody should be
func ResponseBody(c *fiber.Ctx, statusCode int, message string, data interface{}, err interface{}) error {
	rBody := ResponseBodyStruct{
		StatusCode: statusCode,
		Status:     http.StatusText(statusCode),
		Message:    message,
		Data:       data,
	}

	switch t := err.(type) {
	default:
		rBody.Err.Message = t.(string)
	case error:
		rBody.Err.Message = t.Error()
	case ErrorStruct:
		rBody.Err.Code = t.Code
		rBody.Err.Message = t.Message
		rBody.Err.Stack = t.Stack
	case string:
		rBody.Err.Message = t
	}

	iErr := c.Status(statusCode).JSON(rBody)
	if err != nil {
		return iErr
	}
	return nil
}
