package validate

import (
	"fmt"
	"net/http"
	"reflect"
	"strings"
	"sync"

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
			err := fmt.Errorf("pass 'reqStruct' as pointer")
			return response.NewBody(c, http.StatusInternalServerError, "Invalid Request Body", nil, err)
		}

		if err := c.BodyParser(&rBody); err != nil {
			return response.NewBody(c, http.StatusBadRequest, "Invalid Request Body", nil, err)
		}

		err := ValidateStruct(rBody)
		if len(err) > 0 {
			return response.NewBody(c, http.StatusBadRequest, "Invalid Request Body", nil, err)
		}

		return c.Next()
	}
}

func ValidateStruct(rBody interface{}) []map[string]interface{} {
	var (
		once  sync.Once
		vErrs []map[string]interface{}
	)

	once.Do(func() {
		v.RegisterTagNameFunc(func(fld reflect.StructField) string {
			name := strings.SplitN(fld.Tag.Get("json"), ",", 2)[0]
			if name == "-" {
				return ""
			}
			return name
		})
	})

	errs := v.Struct(rBody)
	if errs != nil {
		for _, err := range errs.(validator.ValidationErrors) {
			// NOTE: just for testing out the values
			// fmt.Println(err.Namespace()) // can differ when a custom TagNameFunc is registered or
			// fmt.Println(err.Field())     // by passing alt name to ReportError like below
			// fmt.Println(err.StructNamespace())
			// fmt.Println(err.StructField())
			// fmt.Println(err.Tag())
			// fmt.Println(err.ActualTag())
			// fmt.Println(err.Kind())
			// fmt.Println(err.Type())
			// fmt.Println(err.Value())
			// fmt.Println(err.Param())
			// fmt.Println()

			vErr := map[string]interface{}{
				"field": err.Field(),
				"type":  err.Type(),
				"value": err.Value(),
			}

			if err.Tag() == "required" {
				vErr["msg"] = fmt.Sprintf("%s Is Mandatory", err.Field())
			} else if err.Tag() == "alphanum" {
				vErr["msg"] = fmt.Sprintf("%s Should Be Alphanumeric", err.Field())
			} else if err.Tag() == "max" {
				vErr["msg"] = fmt.Sprintf("%s Can Be Of Maximum %s Character", err.Field(), err.Param())
			} else if err.Tag() == "eqfield" {
				vErr["msg"] = fmt.Sprintf("%s Does Not Match With %s", err.Field(), err.Param())
			} else if err.Tag() == "min" {
				vErr["msg"] = fmt.Sprintf("%s Can Be Of Minimum %s Character", err.Field(), err.Param())
			} else if err.Tag() == "email" {
				vErr["msg"] = fmt.Sprintf("%s Should Be A Valid Email", err.Field())
			} else if err.Tag() == "len" {
				vErr["msg"] = fmt.Sprintf("%s Length Should Be %s", err.Field(), err.Param())
			} else {
				vErr["msg"] = fmt.Sprintf("Failed Validation For %s", err.Field())
			}

			vErrs = append(vErrs, vErr)
		}
	}

	return vErrs
}
