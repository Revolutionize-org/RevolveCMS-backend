package validation

import (
	"github.com/go-playground/validator/v10"
)

var validate *validator.Validate

func init() {
	validate = validator.New()
}

func ValidateStruct(data interface{}) validator.ValidationErrors {
	err := validate.Struct(data)

	if err != nil {
		return err.(validator.ValidationErrors)
	}
	return nil
}
