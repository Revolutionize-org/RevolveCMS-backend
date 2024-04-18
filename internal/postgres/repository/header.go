package repository

import (
	"github.com/Revolutionize-org/RevolveCMS-backend/internal/gql/model"
	"github.com/go-pg/pg/v10"
)

type HeaderRepo struct {
	DB *pg.DB
}

func NewHeaderRepo(db *pg.DB) *HeaderRepo {
	return &HeaderRepo{DB: db}
}

func (h *HeaderRepo) GetByWebsiteID(websiteID string) (*model.Header, error) {
	var header model.Header
	err := h.DB.Model(&header).Where("website_id = ?", websiteID).First()
	return &header, err
}
