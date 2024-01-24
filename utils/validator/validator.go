package validator

import "github.com/go-playground/validator/v10"

// GetValidatorController return validator controller
func GetValidatorController() *validator.Validate {
	return validator.New()
}
