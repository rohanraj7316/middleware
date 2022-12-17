package response

import (
	"encoding/json"
	"net/http"

	"github.com/gofiber/fiber/v2"
)

type BodyStruct struct {
	StatusCode int         `json:"statusCode,omitempty"`
	Status     string      `json:"status,omitempty"`
	Message    interface{} `json:"message,omitempty"`
	Err        interface{} `json:"error,omitempty"`
	Data       interface{} `json:"data,omitempty"`
}

// NewBody uses BodyStruct to send back the json response
func NewBody(c *fiber.Ctx, statusCode int, message interface{}, data interface{}, err interface{}) error {
	rBody := BodyStruct{
		StatusCode: statusCode,
		Status:     http.StatusText(statusCode),
		Message:    message,
		Data:       data,
	}

	if err != nil {
		switch t := err.(type) {
		case error:
			rBody.Err = t.Error()
		case string:
			rBody.Err = t
		case []byte:
			rBody.Err = string(t)
		case interface{}:
			rBody.Err = t
		}
	}

	errHandler := c.App().ErrorHandler

	iErr := c.Status(statusCode).JSON(rBody)
	if err != nil {
		_ = errHandler(c, iErr)
	}
	return nil
}

func (b *BodyStruct) ToJson() ([]byte, error) {
	switch t := b.Err.(type) {
	case error:
		b.Err = t.Error()
	case string:
		b.Err = t
	case []byte:
		b.Err = string(t)
	case interface{}:
		b.Err = t
	}

	switch t := b.Message.(type) {
	case error:
		b.Message = t.Error()
	case string:
		b.Message = t
	case []byte:
		b.Message = string(t)
	case interface{}:
		b.Message = t
	}

	by, err := json.Marshal(b)
	if err != nil {
		return nil, err
	}

	return by, nil
}
