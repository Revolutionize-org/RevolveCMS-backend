package repository

import (
	"github.com/Revolutionize-org/RevolveCMS-backend/internal/gql/model"
	"github.com/go-pg/pg/v10"
)

type ThemeRepo struct {
	DB *pg.DB
}

func NewThemeRepo(db *pg.DB) *ThemeRepo {
	return &ThemeRepo{DB: db}
}

func (t *ThemeRepo) GetAll() ([]*model.Theme, error) {
	var themes []*model.Theme
	err := t.DB.Model(&themes).Select()
	return themes, err
}

func (t *ThemeRepo) GetByID(ID string) (*model.Theme, error) {
	var theme model.Theme
	err := t.DB.Model(&theme).Where("id = ?", ID).First()
	return &theme, err
}
