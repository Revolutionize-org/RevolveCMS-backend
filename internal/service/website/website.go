package website

import (
	"context"
	"errors"

	"github.com/Revolutionize-org/RevolveCMS-backend/internal/database/repository"
	"github.com/Revolutionize-org/RevolveCMS-backend/internal/errorutil"
	"github.com/Revolutionize-org/RevolveCMS-backend/internal/gql/model"
	"github.com/Revolutionize-org/RevolveCMS-backend/internal/userutil"
)

type Service interface {
	GetWebsite(ctx context.Context) (*model.Website, error)
	ModifyWebsiteTheme(ctx context.Context, id string, themeID string) (*model.Website, error)

	GetHeader(ctx context.Context) (*model.Header, error)
	CreateHeader(ctx context.Context, h model.HeaderInput) (*model.Header, error)
	DeleteHeader(ctx context.Context, id string) (bool, error)
	ModifyHeader(ctx context.Context, h model.HeaderInput) (*model.Header, error)

	GetPages(ctx context.Context) ([]*model.Page, error)
	CreatePage(ctx context.Context, p model.PageInput) (*model.Page, error)
	DeletePage(ctx context.Context, id string) (bool, error)
	ModifyPage(ctx context.Context, p model.PageInput) (*model.Page, error)

	GetFooter(ctx context.Context) (*model.Footer, error)
	CreateFooter(ctx context.Context, h model.FooterInput) (*model.Footer, error)
	DeleteFooter(ctx context.Context, id string) (bool, error)
	ModifyFooter(ctx context.Context, h model.FooterInput) (*model.Footer, error)

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
	_, website, err := w.retrieveWebsiteViaCtx(ctx)
	if err != nil {
		return nil, err
	}
	return website, nil
}

func (w *websiteService) ModifyWebsiteTheme(ctx context.Context, id string, themeID string) (*model.Website, error) {
	_, website, err := w.retrieveWebsiteViaCtx(ctx)
	if err != nil {
		return nil, err
	}

	website.ThemeID = themeID
	didUpdate, err := w.WebsiteRepo.ModiftyWebsiteTheme(website)
	if err != nil {
		return nil, errorutil.HandleErrorDependingEnv(err)
	}

	if !didUpdate {
		return nil, errors.New("website not found")
	}
	return website, nil
}

func (w *websiteService) GetService() *websiteService {
	return w
}

func (w *websiteService) retrieveWebsiteViaCtx(ctx context.Context) (*model.User, *model.Website, error) {
	user, err := userutil.RetrieveUser(ctx, w.UserRepo)
	if err != nil {
		return nil, nil, errorutil.HandleError(err, "user not found")
	}

	website, err := w.WebsiteRepo.GetWebsiteByID(user.WebsiteID)
	if err != nil {
		return nil, nil, errorutil.HandleError(err, "website not found")
	}
	return user, website, nil
}
