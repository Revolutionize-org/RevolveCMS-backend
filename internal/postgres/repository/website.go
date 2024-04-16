package repository

import (
	"github.com/Revolutionize-org/RevolveCMS-backend/internal/gql/model"
	"github.com/go-pg/pg/v10"
)

type WebsiteRepo struct {
	DB *pg.DB
}

func NewWebsiteRepo(db *pg.DB) *WebsiteRepo {
	return &WebsiteRepo{DB: db}
}

func (w *WebsiteRepo) GetByID(id string) (*model.Website, error) {
	var website model.Website
	err := w.DB.Model(&website).Where("id = ?", id).First()
	return &website, err
}
