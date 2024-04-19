package errorutil

import (
	"errors"

	"github.com/Revolutionize-org/RevolveCMS-backend/internal/config"
)

func HandleError(err error) error {
	if config.Config.Api.Env == "dev" {
		return err
	}
	return errors.New("internal server error")
}
