package sharedinfra

import (
	"fmt"
	"reflect"
	"strings"

	"github.com/go-playground/validator/v10"
)

type ValidationError struct {
	Field   string `json:"field"`
	Message string `json:"message"`
}

type Validate struct {
	rules map[string]string
}

func NewValidate() *Validate {
	return &Validate{
		rules: map[string]string{
			"required": "this field is required.",
			"min":      "minimum length is %s characters.",
			"max":      "maximum length is %s characters.",
		},
	}
}

func (v *Validate) Execute(payload interface{}) []ValidationError {
	validate := validator.New(validator.WithRequiredStructEnabled())

	if err := validate.Struct(payload); err != nil {
		validationErrors := err.(validator.ValidationErrors)
		structType := reflect.TypeOf(payload)
		if structType.Kind() == reflect.Ptr {
			structType = structType.Elem()
		}

		var errorsList []ValidationError

		for _, e := range validationErrors {
			jsonField := v.getJSONFieldName(structType, e.Field())

			var message string
			switch e.Tag() {
			case "required":
				message = v.rules["required"]
			case "min":
				message = fmt.Sprintf(v.rules["min"], e.Param())
			case "max":
				message = fmt.Sprintf(v.rules["max"], e.Param())
			default:
				message = "Invalid value."
			}

			errorsList = append(errorsList, ValidationError{
				Field:   jsonField,
				Message: message,
			})
		}

		return errorsList
	}

	return nil
}

func (v *Validate) getJSONFieldName(structType reflect.Type, fieldName string) string {
	if field, ok := structType.FieldByName(fieldName); ok {
		tag := field.Tag.Get("json")
		if tag != "" && tag != "-" {
			return strings.Split(tag, ",")[0]
		}
	}
	return fieldName
}

//TODO: Improv this validation
// validate := validator.New(validator.WithRequiredStructEnabled())

// if err := validate.Struct(req); err != nil {
// 	validationErrors := err.(validator.ValidationErrors)
// 	errorsMap := make(map[string]string)
// 	for _, e := range validationErrors {
// 		field := e.Field()
// 		tag := e.Tag()

// 		switch tag {
// 		case "required":
// 			errorsMap[field] = "This field is required."
// 		case "min":
// 			errorsMap[field] = "Minimum length is " + e.Param() + " characters."
// 		case "max":
// 			errorsMap[field] = "Maximum length is " + e.Param() + " characters."
// 		default:
// 			errorsMap[field] = "Invalid value."
// 		}
// 	}

// 	return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"validation_errors": errorsMap})
// }
