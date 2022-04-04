package validate

import (
	"fmt"
	"net/http"
	"reflect"
	"strings"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/rohanraj7316/middleware/libs/response"
)

var v = validator.New()

// Request fiber.Handler a middleware used to validate
// request body.
// reqStruct pointer to your request structure
// ie: a.Post("/", ValidateRequest(new(CreateServiceRequest)), handler.CreateService)
func Request(reqStruct interface{}) fiber.Handler {
	return func(c *fiber.Ctx) error {
		rBody := reqStruct

		if reflect.ValueOf(reqStruct).Kind() != reflect.Ptr {
			err := fmt.Errorf("[ValidationError] pass pointer as 'reqStruct' type")
			return response.NewBody(c, http.StatusInternalServerError, "[ValidationError] unexpected 'reqStruct' type", nil, err)
		}

		if err := c.BodyParser(&rBody); err != nil {
			return response.NewBody(c, http.StatusBadRequest, "[ValidationError] failed to parse request body", nil, err)
		}

		err := v.Struct(rBody)
		if err != nil {
			var vErrs strings.Builder
			for _, err := range err.(validator.ValidationErrors) {
				vErrs.WriteString(fmt.Sprintf("key: %s | val: %s | tag: %s; ", err.Field(), err.Tag(), err.Value()))
			}

			vErrsMsg := fmt.Errorf("[ValidationError] %s", &vErrs)
			return response.NewBody(c, http.StatusBadRequest, "failed to validated request body", nil, vErrsMsg)
		}
		return c.Next()
	}
}
