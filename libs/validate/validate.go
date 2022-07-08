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

	v.RegisterTagNameFunc(func(fld reflect.StructField) string {
		name := strings.SplitN(fld.Tag.Get("json"), ",", 2)[0]
		if name == "-" {
			return ""
		}
		return name
	})

	return func(c *fiber.Ctx) error {
		rBody := reqStruct

		if reflect.ValueOf(reqStruct).Kind() != reflect.Ptr {
			err := fmt.Errorf("pass 'reqStruct' as pointer")
			return response.NewBody(c, http.StatusInternalServerError, "Invalid Request Body", nil, err)
		}

		if err := c.BodyParser(&rBody); err != nil {
			return response.NewBody(c, http.StatusBadRequest, "Invalid Request Body", nil, err)
		}

		err := v.Struct(rBody)
		if err != nil {
			var vErrs strings.Builder
			for _, err := range err.(validator.ValidationErrors) {
				if vErrs.Len() == 0 {
					vErrs.WriteString(fmt.Sprintf("%s: %s;", err.Field(), err.Tag()))
				} else {
					vErrs.WriteString(fmt.Sprintf(" %s: %s;", err.Field(), err.Tag()))
				}
			}

			vErrsMsg := fmt.Errorf("%s", &vErrs)
			return response.NewBody(c, http.StatusBadRequest, "Invalid Request Body", nil, vErrsMsg)
		}
		return c.Next()
	}
}
