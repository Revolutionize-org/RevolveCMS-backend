package postgres

import (
	"github.com/go-pg/pg/v10"
	"github.com/google/uuid"
	"github.com/matthewhartstonge/argon2"
)

type Token struct {
	tableName struct{} `pg:"token"`
	ID        string   `json:"id"`
	Token     string   `json:"token"`
}

type TokenRepo struct {
	DB *pg.DB
}

func (tr *TokenRepo) Add(t string) error {
	uuid, err := uuid.NewRandom()
	if err != nil {
		return err
	}

	argon := argon2.DefaultConfig()

	tokenHash, err := argon.HashEncoded([]byte(t))

	if err != nil {
		return err
	}

	token := &Token{
		ID:    uuid.String(),
		Token: string(tokenHash),
	}

	_, err = tr.DB.Model(token).Insert()

	return err
}
