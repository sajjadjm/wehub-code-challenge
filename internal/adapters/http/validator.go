package http

import (
	"fmt"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

// Initialize a new validator instance
var validate = validator.New()

// ValidateStruct Generic function to validate any struct
func ValidateStruct(s interface{}) []*fiber.Error {
	var errors []*fiber.Error
	err := validate.Struct(s)
	if err != nil {
		// Iterate through validation errors
		for _, err := range err.(validator.ValidationErrors) {
			// Create an error for each validation failure
			e := fiber.NewError(fiber.StatusBadRequest, fmt.Sprintf("Field validation for '%s' failed on the '%s' tag", err.Field(), err.Tag()))
			errors = append(errors, e)
		}
	}
	return errors
}
