package postgres

import (
	"github.com/go-pg/pg/v10"
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

func (tr *TokenRepo) Get(uuid string) (*Token, error) {
	token := &Token{
		ID: uuid,
	}
	err := tr.DB.Model(token).Where("id = ?", uuid).Select()
	return token, err
}

func (tr *TokenRepo) Add(uuid string, t string) error {
	argon := argon2.DefaultConfig()

	tokenHash, err := argon.HashEncoded([]byte(t))

	if err != nil {
		return err
	}

	token := &Token{
		ID:    uuid,
		Token: string(tokenHash),
	}

	_, err = tr.DB.Model(token).Insert()

	return err
}

func (tr *TokenRepo) Delete(id string) (bool, error) {
	token := &Token{
		ID: id,
	}
	result, err := tr.DB.Model(token).Where("id = ?", token.ID).Delete()

	if err != nil {
		return false, err
	}
	return result.RowsAffected() > 0, nil
}
