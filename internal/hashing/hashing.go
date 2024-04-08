package hashing

import (
	"errors"

	"github.com/matthewhartstonge/argon2"
)

var Argon argon2.Config

func init() {
	Argon = argon2.DefaultConfig()
}

func CompareHashAndPassword(hashed, password string) error {
	ok, err := argon2.VerifyEncoded([]byte(password), []byte(hashed))

	if err != nil {
		return err
	}

	if !ok {
		return errors.New("invalid password")
	}
	return nil
}
