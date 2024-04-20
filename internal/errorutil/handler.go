package errorutil

import (
	"errors"

	"github.com/Revolutionize-org/RevolveCMS-backend/internal/config"
	"github.com/go-pg/pg/v10"
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

func HandleError(err error, message string) error {
	if err := CheckErrNoRows(err, message); err != nil {
		return err
	}
	return HandleErrorDependingEnv(err)
}
