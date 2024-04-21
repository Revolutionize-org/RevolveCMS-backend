package errorutil

import (
	"context"
	"errors"

	"github.com/99designs/gqlgen/graphql"
	"github.com/Revolutionize-org/RevolveCMS-backend/internal/config"
	"github.com/go-pg/pg/v10"
	"github.com/vektah/gqlparser/v2/gqlerror"
)

func HandleErrorDependingEnv(err error) error {
	if config.Config.Api.Env == "dev" {
		return err
	}
	return errors.New("internal server error")
}

func CheckErrNoRows(err error, message string) error {
	if err != nil {
		if err.Error() == pg.ErrNoRows.Error() {
			return errors.New(message)
		}
		return err
	}
	return nil
}

func HandleErrorOrNoRows(err error, message string) error {
	if err := CheckErrNoRows(err, message); err != nil {
		return err
	}
	return HandleErrorDependingEnv(err)
}

func AddGraphQLErrors(ctx context.Context, errors []*gqlerror.Error) {
	for _, err := range errors {
		graphql.AddError(ctx, err)
	}
}
