package helpers

import (
	"github.com/go-playground/validator/v10"
)

var validate *validator.Validate

func init() {
	validate = validator.New()
}

func ValidateStruct(s interface{}) map[string]string {
	err := validate.Struct(s)
	if err == nil {
		return nil
	}

	errors := make(map[string]string)
	for _, err := range err.(validator.ValidationErrors) {
		errors[err.Field()] = getErrorMessage(err)
	}

	return errors
}

func getErrorMessage(err validator.FieldError) string {
	switch err.Tag() {
	case "required":
		return "This field is required"
	case "email":
		return "Invalid email format"
	case "min":
		return "Value is too short"
	case "max":
		return "Value is too long"
	case "gte":
		return "Value must be greater than or equal to " + err.Param()
	case "lte":
		return "Value must be less than or equal to " + err.Param()
	case "eqfield":
		return "Field must match " + err.Param()
	default:
		return "Invalid value"
	}
}

func GetValidator() *validator.Validate {
	return validate
}
