package repository

import (
	"github.com/Revolutionize-org/RevolveCMS-backend/internal/gql/model"
	"github.com/go-pg/pg/v10"
)

type UserRepo struct {
	DB *pg.DB
}

func NewUserRepo(DB *pg.DB) *UserRepo {
	return &UserRepo{DB: DB}
}

func (u *UserRepo) GetByID(id string) (*model.User, error) {
	var user model.User
	err := u.DB.Model(&user).Where("id = ?", id).First()
	return &user, err
}

func (u *UserRepo) GetByEmail(email string) (*model.User, error) {
	var user model.User
	err := u.DB.Model(&user).Where("email = ?", email).First()
	return &user, err
}
