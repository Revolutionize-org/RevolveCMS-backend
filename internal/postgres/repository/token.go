package repository

import (
	"time"

	"github.com/go-pg/pg/v10"
)

type Token struct {
	ID        string    `json:"id"`
	ExpiresAt time.Time `json:"expires_at"`
}

type TokenRepo interface {
	Get(jti string) (*Token, error)
	Add(jti string, exp time.Time) error
	Delete(id string) (bool, error)
}

type SqlTokenRepo struct {
	DB *pg.DB
}

func NewTokenRepo(db *pg.DB) SqlTokenRepo {
	return SqlTokenRepo{DB: db}
}

func (tr SqlTokenRepo) Get(jti string) (*Token, error) {
	token := &Token{
		ID: jti,
	}
	err := tr.DB.Model(token).Where("id = ?", jti).Select()
	return token, err
}

func (tr SqlTokenRepo) Add(jti string, exp time.Time) error {
	token := &Token{
		ID:        jti,
		ExpiresAt: exp,
	}

	_, err := tr.DB.Model(token).Insert()

	return err
}

func (tr SqlTokenRepo) Delete(id string) (bool, error) {
	token := &Token{
		ID: id,
	}
	result, err := tr.DB.Model(token).Where("id = ?", token.ID).Delete()

	if err != nil {
		return false, err
	}
	return result.RowsAffected() > 0, nil
}
