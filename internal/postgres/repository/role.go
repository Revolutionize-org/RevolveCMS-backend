package repository

import (
	"github.com/Revolutionize-org/RevolveCMS-backend/internal/gql/model"
	"github.com/go-pg/pg/v10"
)

type RoleRepo struct {
	DB *pg.DB
}

func NewRoleRepo(DB *pg.DB) RoleRepo {
	return RoleRepo{DB: DB}
}

func (r RoleRepo) GetByID(id string) (*model.Role, error) {
	var role model.Role
	err := r.DB.Model(&role).Where("id = ?", id).First()
	return &role, err
}
