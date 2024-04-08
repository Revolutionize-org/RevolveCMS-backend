package postgres

import (
	"errors"

	"github.com/go-pg/pg/v10"
)

func CheckErrNoRows(err error, message string) error {
	if err != nil {
		if err.Error() == pg.ErrNoRows.Error() {
			return errors.New(message)
		}
		return err
	}
	return nil
}
