package repository

import (
	"github.com/Revolutionize-org/RevolveCMS-backend/internal/gql/model"
	"github.com/go-pg/pg/v10"
)

type WebsiteRepo interface {
	GetWebsiteByID(id string) (*model.Website, error)
	GetAllThemes() ([]*model.Theme, error)
	GetThemeByID(id string) (*model.Theme, error)
	GetFooterByWebsiteID(id string) (*model.Footer, error)
	GetPagesByWebsiteID(id string) ([]*model.Page, error)
	GetHeaderByWebsiteID(id string) (*model.Header, error)
	CreateHeader(header *model.Header) error
	DeleteHeader(id string) error
	ModifyHeader(header *model.Header) error
}

type SqlWebsiteRepo struct {
	DB *pg.DB
}

func NewWebsiteRepo(db *pg.DB) SqlWebsiteRepo {
	return SqlWebsiteRepo{DB: db}
}

func (w SqlWebsiteRepo) GetWebsiteByID(id string) (*model.Website, error) {
	var website model.Website
	err := w.DB.Model(&website).Where("id = ?", id).First()
	return &website, err
}

func (w SqlWebsiteRepo) GetAllThemes() ([]*model.Theme, error) {
	var themes []*model.Theme
	err := w.DB.Model(&themes).Select()
	return themes, err
}

func (w SqlWebsiteRepo) GetThemeByID(id string) (*model.Theme, error) {
	var theme model.Theme
	err := w.DB.Model(&theme).Where("id = ?", id).First()
	return &theme, err
}

func (w SqlWebsiteRepo) GetHeaderByWebsiteID(id string) (*model.Header, error) {
	var header model.Header
	err := w.DB.Model(&header).Where("website_id = ?", id).First()
	return &header, err
}

func (w SqlWebsiteRepo) CreateHeader(header *model.Header) error {
	_, err := w.DB.Model(header).Insert()
	return err
}

func (w SqlWebsiteRepo) DeleteHeader(id string) error {
	header := &model.Header{
		ID: id,
	}

	_, err := w.DB.Model(header).WherePK().Delete()
	return err
}

func (w SqlWebsiteRepo) ModifyHeader(header *model.Header) error {
	_, err := w.DB.Model(header).WherePK().Update()
	return err
}

func (w SqlWebsiteRepo) GetFooterByWebsiteID(id string) (*model.Footer, error) {
	var footer model.Footer
	err := w.DB.Model(&footer).Where("website_id = ?", id).First()
	return &footer, err
}

func (w SqlWebsiteRepo) GetPagesByWebsiteID(id string) ([]*model.Page, error) {
	var pages []*model.Page
	err := w.DB.Model(&pages).Where("website_id = ?", id).Select()
	return pages, err
}
