package validation

import (
	"context"
	"errors"

	"github.com/99designs/gqlgen/graphql"
	"github.com/Revolutionize-org/RevolveCMS-backend/internal/config"

	"github.com/go-playground/validator/v10"
	"github.com/vektah/gqlparser/v2/gqlerror"
)

var validate *validator.Validate

func init() {
	validate = validator.New()
}

func ValidateStruct[T any](data T) validator.ValidationErrors {
	if err := validate.Struct(data); err != nil {
		return err.(validator.ValidationErrors)
	}
	return nil
}

func ValidateInput[T any](ctx context.Context, data T) error {
	validErr := ValidateStruct[T](data)

	if len(validErr) > 0 {
		for _, err := range validErr {
			if config.Config.Api.Env == "dev" {
				graphql.AddError(ctx, &gqlerror.Error{
					Message: err.Error(),
					Extensions: map[string]interface{}{
						"field": err.Field(),
					},
				})
			}
		}
		return errors.New("invalid input received")
	}

	return nil
}
