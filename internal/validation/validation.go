package validation

import (
	"context"
	"errors"
	"fmt"

	"github.com/99designs/gqlgen/graphql"
	"github.com/go-playground/validator/v10"
	"github.com/vektah/gqlparser/v2/gqlerror"
)

var validate *validator.Validate

func init() {
	validate = validator.New()
}

func ValidateStruct(data interface{}) validator.ValidationErrors {
	if err := validate.Struct(data); err != nil {
		return err.(validator.ValidationErrors)
	}
	return nil
}

func ValidateInput(ctx context.Context, data interface{}) error {
	validErr := ValidateStruct(data)

	if len(validErr) > 0 {
		for _, err := range validErr {
			fmt.Print(err)
			graphql.AddError(ctx, &gqlerror.Error{
				Message: err.Error(),
				Extensions: map[string]interface{}{
					"field": err.Field(),
				},
			})
		}
		return errors.New("invalid input received")
	}

	return nil
}
