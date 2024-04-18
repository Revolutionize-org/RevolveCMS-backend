package repository

import (
	"github.com/Revolutionize-org/RevolveCMS-backend/internal/gql/model"
	"github.com/go-pg/pg/v10"
)

type UserRepo interface {
	GetByID(id string) (*model.User, error)
	GetByEmail(email string) (*model.User, error)
}

type SqlUserRepo struct {
	DB *pg.DB
}

func NewUserRepo(DB *pg.DB) SqlUserRepo {
	return SqlUserRepo{DB: DB}
}

func (u SqlUserRepo) GetByID(id string) (*model.User, error) {
	var user model.User
	err := u.DB.Model(&user).Where("id = ?", id).First()
	return &user, err
}

func (u SqlUserRepo) GetByEmail(email string) (*model.User, error) {
	var user model.User
	err := u.DB.Model(&user).Where("email = ?", email).First()
	return &user, err
}
