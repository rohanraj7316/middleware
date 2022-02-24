package helpers

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
)

type Response struct {
	StatusCode int         `json:"statusCode,omitempty"`
	Status     string      `json:"status,omitempty"`
	Message    interface{} `json:"message,omitempty"`
	Error      interface{} `json:"error,omitempty"`
	Data       interface{} `json:"data,omitempty"`
}

func ResponseBody(c *fiber.Ctx, statusCode int, message string, data interface{}, err error) error {
	rBody := Response{
		StatusCode: statusCode,
		Status:     http.StatusText(statusCode),
		Message:    message,
		Data:       data,
		Error:      err,
	}

	err = c.Status(statusCode).JSON(rBody)
	if err != nil {
		return err
	}
	return nil
}
