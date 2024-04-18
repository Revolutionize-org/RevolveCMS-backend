package website

import (
	"context"

	"github.com/Revolutionize-org/RevolveCMS-backend/internal/gql/model"
	"github.com/Revolutionize-org/RevolveCMS-backend/internal/postgres/repository"
	"github.com/Revolutionize-org/RevolveCMS-backend/internal/userutil"
)

type Service interface {
	GetWebsite(ctx context.Context) (*model.Website, error)
	GetHeader(ctx context.Context) (*model.Header, error)
	GetPages(ctx context.Context) ([]*model.Page, error)
	GetFooter(ctx context.Context) (*model.Footer, error)
	GetService() websiteService
}

type websiteService struct {
	WebsiteRepo repository.WebsiteRepo
	UserRepo    repository.UserRepo
}

func New(webisteRepo repository.WebsiteRepo, userRepo repository.UserRepo) Service {
	return websiteService{
		WebsiteRepo: webisteRepo,
		UserRepo:    userRepo,
	}
}

func (w websiteService) GetWebsite(ctx context.Context) (*model.Website, error) {
	user, err := userutil.RetrieveUser(ctx, w.UserRepo)
	if err != nil {
		return nil, err
	}
	return w.WebsiteRepo.GetWebsiteByID(user.WebsiteID)
}

func (w websiteService) GetHeader(ctx context.Context) (*model.Header, error) {
	user, err := userutil.RetrieveUser(ctx, w.UserRepo)
	if err != nil {
		return nil, err
	}

	webiste, err := w.WebsiteRepo.GetWebsiteByID(user.ID)
	if err != nil {
		return nil, err
	}
	return w.WebsiteRepo.GetHeaderByWebsiteID(webiste.ID)
}

func (w websiteService) GetPages(ctx context.Context) ([]*model.Page, error) {
	user, err := userutil.RetrieveUser(ctx, w.UserRepo)
	if err != nil {
		return nil, err
	}

	webiste, err := w.WebsiteRepo.GetWebsiteByID(user.WebsiteID)
	if err != nil {
		return nil, err
	}
	return w.WebsiteRepo.GetPagesByWebsiteID(webiste.ID)
}

func (w websiteService) GetFooter(ctx context.Context) (*model.Footer, error) {
	user, err := userutil.RetrieveUser(ctx, w.UserRepo)
	if err != nil {
		return nil, err
	}

	webiste, err := w.WebsiteRepo.GetWebsiteByID(user.WebsiteID)
	if err != nil {
		return nil, err
	}
	return w.WebsiteRepo.GetFooterByWebsiteID(webiste.ID)
}

func (w websiteService) GetService() websiteService {
	return w
}
