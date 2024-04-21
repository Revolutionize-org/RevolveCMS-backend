package validation

import (
	"context"
	"errors"
	"fmt"
	"strings"

	"github.com/Revolutionize-org/RevolveCMS-backend/internal/gql/model"
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

func ValidateInput[T any](ctx context.Context, data T) []*gqlerror.Error {
	validErr := ValidateStruct[T](data)
	var errors []*gqlerror.Error

	if len(validErr) > 0 {
		for _, err := range validErr {
			errors = append(errors, &gqlerror.Error{
				Message: fmt.Sprintf("invalid %s received", err.Field()),
				Extensions: map[string]interface{}{
					"field": strings.ToLower(err.Field()),
				},
			})

		}
	}
	return errors
}

func ValidatePageInput(page model.PageInput) error {
	if page.Name == "" {
		return errors.New("invalid name received")
	}

	if page.Slug == "" {
		return errors.New("invalid slug received")
	}

	if page.Data == "" {
		return errors.New("invalid data received")
	}
	return nil
}

func ValidateHeaderInput(header model.HeaderInput) error {
	if header.Name == "" {
		return errors.New("invalid name received")
	}

	if header.Data == "" {
		return errors.New("invalid data received")
	}
	return nil
}

func ValidateFooterInput(footer model.FooterInput) error {
	if footer.Name == "" {
		return errors.New("invalid name received")
	}

	if footer.Data == "" {
		return errors.New("invalid data received")
	}
	return nil
}
