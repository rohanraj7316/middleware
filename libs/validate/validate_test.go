package validate_test

import (
	"testing"

	"github.com/rohanraj7316/middleware/libs/validate"
)

type Student struct {
	Name        string `json:"policyName" validate:"required"`
	Description string `json:"description" validate:"required,startsnotwith=" ""`
}

type StudentData struct {
	Students Student `json:"students" validate:"required"`
	Type     int     `json:"type" validate:"required"`
}

func TestValidateString(t *testing.T) {
	studentData := StudentData{
		Students: Student{
			Name:        "Rohan Raj",
			Description: " ",
		},
	}

	err := validate.ValidateStruct(studentData)
	if err != nil {
		t.Logf("%+v", err)
		t.Errorf("failed to parse struct")
	}
}
