package validators

import (
	"fmt"

	"github.com/go-playground/validator/v10"
)

var validate = validator.New()

type ValidationError struct {
	Field   string
	Message string
}

func (e *ValidationError) Error() string {
	return e.Message
}

func formatError(err validator.FieldError) string {
	switch err.Tag() {
	case "required":
		return fmt.Sprintf("%s is required", err.Field())
	case "min":
		return fmt.Sprintf("%s must be at least %s characters", err.Field(), err.Param())
	case "max":
		return fmt.Sprintf("%s cannot exceed %s characters", err.Field(), err.Param())
	case "len":
		return fmt.Sprintf("%s must be exactly %s characters", err.Field(), err.Param())
	case "numeric":
		return fmt.Sprintf("%s must contain only numbers", err.Field())
	default:
		return fmt.Sprintf("%s is invalid", err.Field())
	}
}

func ValidateStruct(s any) *ValidationError {
	err := validate.Struct(s)
	if err != nil {
		for _, e := range err.(validator.ValidationErrors) {
			return &ValidationError{
				Field:   e.Field(),
				Message: formatError(e),
			}
		}
	}
	return nil
}

func ValidatePagination(page int, limit int) *ValidationError {
	if page < 1 {
		return &ValidationError{
			Field:   "page",
			Message: "page must be greater than 0",
		}
	}

	if limit < 1 || limit > 100 {
		return &ValidationError{
			Field:   "limit",
			Message: "limit must be between 1 and 100",
		}
	}

	return nil
}
