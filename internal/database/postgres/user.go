package postgres

import (
	"github.com/Revolutionize-org/RevolveCMS-backend/internal/gql/model"
	"github.com/go-pg/pg/v10"
)

type UserRepo struct {
	DB *pg.DB
}

func (u *UserRepo) GetByEmail(email string) (*model.User, error) {
	var user model.User
	err := u.DB.Model(&user).Where("email = ?", email).First()
	return &user, err
}
