package hashing

import (
	"errors"

	"github.com/matthewhartstonge/argon2"
)

var Argon argon2.Config

func init() {
	Argon = argon2.DefaultConfig()
}

func CompareHashAndSecret(hashed, secret string) error {
	ok, err := argon2.VerifyEncoded([]byte(secret), []byte(hashed))

	if err != nil {
		return err
	}

	if !ok {
		return errors.New("invalid secret")
	}
	return nil
}
