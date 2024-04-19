package website

import (
	"context"

	"github.com/Revolutionize-org/RevolveCMS-backend/internal/errorutil"
	"github.com/Revolutionize-org/RevolveCMS-backend/internal/gql/model"
	"github.com/Revolutionize-org/RevolveCMS-backend/internal/postgres"
	"github.com/Revolutionize-org/RevolveCMS-backend/internal/postgres/repository"
	"github.com/Revolutionize-org/RevolveCMS-backend/internal/userutil"
)

type Service interface {
	GetWebsite(ctx context.Context) (*model.Website, error)
	GetHeader(ctx context.Context) (*model.Header, error)
	GetPages(ctx context.Context) ([]*model.Page, error)
	GetFooter(ctx context.Context) (*model.Footer, error)
	GetService() *websiteService
}

type websiteService struct {
	WebsiteRepo repository.WebsiteRepo
	UserRepo    repository.UserRepo
}

func New(websiteRepo repository.WebsiteRepo, UserRepo repository.UserRepo) Service {
	return &websiteService{
		WebsiteRepo: websiteRepo,
		UserRepo:    UserRepo,
	}
}

func (w *websiteService) GetWebsite(ctx context.Context) (*model.Website, error) {
	_, website, err := w.retrieveUserAndWebsite(ctx)
	if err != nil {
		return nil, err
	}
	return website, nil
}

func (w *websiteService) GetHeader(ctx context.Context) (*model.Header, error) {
	_, website, err := w.retrieveUserAndWebsite(ctx)
	if err != nil {
		return nil, err
	}

	header, err := w.WebsiteRepo.GetHeaderByWebsiteID(website.ID)
	if err != nil {
		return nil, handleRepoError(err, "header not found")
	}
	return header, nil
}

func (w *websiteService) GetPages(ctx context.Context) ([]*model.Page, error) {
	_, website, err := w.retrieveUserAndWebsite(ctx)
	if err != nil {
		return nil, err
	}

	pages, err := w.WebsiteRepo.GetPagesByWebsiteID(website.ID)
	if err != nil {
		return nil, handleRepoError(err, "pages not found")
	}
	return pages, nil
}

func (w *websiteService) GetFooter(ctx context.Context) (*model.Footer, error) {
	_, website, err := w.retrieveUserAndWebsite(ctx)
	if err != nil {
		return nil, err
	}

	footer, err := w.WebsiteRepo.GetFooterByWebsiteID(website.ID)
	if err != nil {
		return nil, handleRepoError(err, "footer not found")
	}
	return footer, nil
}

func (w *websiteService) GetService() *websiteService {
	return w
}

func (w *websiteService) retrieveUserAndWebsite(ctx context.Context) (*model.User, *model.Website, error) {
	user, err := userutil.RetrieveUser(ctx, w.UserRepo)
	if err != nil {
		return nil, nil, handleRepoError(err, "user not found")
	}

	website, err := w.WebsiteRepo.GetWebsiteByID(user.WebsiteID)
	if err != nil {
		return nil, nil, handleRepoError(err, "website not found")
	}
	return user, website, nil
}

func handleRepoError(err error, notFoundMessage string) error {
	if err := postgres.CheckErrNoRows(err, notFoundMessage); err != nil {
		return err
	}
	return errorutil.HandleError(err)
}
