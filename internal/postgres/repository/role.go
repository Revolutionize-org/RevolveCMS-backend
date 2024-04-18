package repository

import (
	"github.com/Revolutionize-org/RevolveCMS-backend/internal/gql/model"
	"github.com/go-pg/pg/v10"
)

type RoleRepo interface {
	GetByID(id string) (*model.Role, error)
}

type SqlRoleRepo struct {
	DB *pg.DB
}

func NewRoleRepo(DB *pg.DB) SqlRoleRepo {
	return SqlRoleRepo{DB: DB}
}

func (r SqlRoleRepo) GetByID(id string) (*model.Role, error) {
	var role model.Role
	err := r.DB.Model(&role).Where("id = ?", id).First()
	return &role, err
}
