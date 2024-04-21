package repository

import (
	"github.com/Revolutionize-org/RevolveCMS-backend/internal/gql/model"
	"github.com/go-pg/pg/v10"
)

type WebsiteRepo interface {
	GetWebsiteByID(id string) (*model.Website, error)
	ModiftyWebsiteTheme(website *model.Website) (bool, error)

	GetAllThemes() ([]*model.Theme, error)
	GetThemeByID(id string) (*model.Theme, error)

	GetFooterByWebsiteID(id string) (*model.Footer, error)
	CreateFooter(footer *model.Footer) error
	DeleteFooter(f *model.Footer) (bool, error)
	ModifyFooter(footer *model.Footer) (bool, error)

	GetPagesByWebsiteID(id string) ([]*model.Page, error)
	CreatePage(page *model.Page) error
	DeletePage(p *model.Page) (bool, error)
	ModifyPage(page *model.Page) (bool, error)

	GetHeaderByWebsiteID(id string) (*model.Header, error)
	CreateHeader(header *model.Header) error
	DeleteHeader(h *model.Header) (bool, error)
	ModifyHeader(header *model.Header) (bool, error)
}

type SqlWebsiteRepo struct {
	DB *pg.DB
}

func NewWebsiteRepo(db *pg.DB) *SqlWebsiteRepo {
	return &SqlWebsiteRepo{DB: db}
}

func (w *SqlWebsiteRepo) GetWebsiteByID(id string) (*model.Website, error) {
	var website model.Website
	err := w.DB.Model(&website).Where("id = ?", id).First()
	return &website, err
}

func (w *SqlWebsiteRepo) ModiftyWebsiteTheme(website *model.Website) (bool, error) {
	_, err := w.GetThemeByID(website.ThemeID)
	if err != nil {
		return false, err
	}

	res, err := w.DB.Model(website).WherePK().UpdateNotZero()
	return res.RowsAffected() > 0, err
}

func (w *SqlWebsiteRepo) GetAllThemes() ([]*model.Theme, error) {
	var themes []*model.Theme
	err := w.DB.Model(&themes).Select()
	return themes, err
}

func (w *SqlWebsiteRepo) GetThemeByID(id string) (*model.Theme, error) {
	var theme model.Theme
	err := w.DB.Model(&theme).Where("id = ?", id).First()
	return &theme, err
}

func (w *SqlWebsiteRepo) GetHeaderByWebsiteID(id string) (*model.Header, error) {
	var header model.Header
	err := w.DB.Model(&header).Where("website_id = ?", id).First()
	return &header, err
}

func (w *SqlWebsiteRepo) CreateHeader(header *model.Header) error {
	_, err := w.DB.Model(header).Insert()
	return err
}

func (w *SqlWebsiteRepo) DeleteHeader(header *model.Header) (bool, error) {
	res, err := w.DB.Model(header).WherePK().Delete()
	return res.RowsAffected() > 0, err
}

func (w *SqlWebsiteRepo) ModifyHeader(header *model.Header) (bool, error) {
	res, err := w.DB.Model(header).WherePK().UpdateNotZero()
	return res.RowsAffected() > 0, err
}

func (w *SqlWebsiteRepo) GetFooterByWebsiteID(id string) (*model.Footer, error) {
	var footer model.Footer
	err := w.DB.Model(&footer).Where("website_id = ?", id).First()
	return &footer, err
}

func (w *SqlWebsiteRepo) CreateFooter(footer *model.Footer) error {
	_, err := w.DB.Model(footer).Insert()
	return err
}

func (w *SqlWebsiteRepo) DeleteFooter(footer *model.Footer) (bool, error) {
	res, err := w.DB.Model(footer).WherePK().Delete()
	return res.RowsAffected() > 0, err
}

func (w *SqlWebsiteRepo) ModifyFooter(footer *model.Footer) (bool, error) {
	res, err := w.DB.Model(footer).WherePK().UpdateNotZero()
	return res.RowsAffected() > 0, err
}

func (w *SqlWebsiteRepo) GetPagesByWebsiteID(id string) ([]*model.Page, error) {
	var pages []*model.Page
	err := w.DB.Model(&pages).Where("website_id = ?", id).Select()
	return pages, err
}

func (w *SqlWebsiteRepo) DeletePage(page *model.Page) (bool, error) {
	res, err := w.DB.Model(page).WherePK().Delete()
	return res.RowsAffected() > 0, err
}

func (w *SqlWebsiteRepo) CreatePage(page *model.Page) error {
	_, err := w.DB.Model(page).Insert()
	return err
}

func (w *SqlWebsiteRepo) ModifyPage(page *model.Page) (bool, error) {
	res, err := w.DB.Model(page).WherePK().UpdateNotZero()
	return res.RowsAffected() > 0, err
}
