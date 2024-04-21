package website

import (
	"context"
	"errors"
	"strings"
	"time"

	"github.com/Revolutionize-org/RevolveCMS-backend/internal/errorutil"
	"github.com/Revolutionize-org/RevolveCMS-backend/internal/gql/model"
	"github.com/Revolutionize-org/RevolveCMS-backend/internal/userutil"
	"github.com/google/uuid"
)

func (w *websiteService) GetPages(ctx context.Context) ([]*model.Page, error) {
	_, website, err := w.retrieveWebsiteViaCtx(ctx)
	if err != nil {
		return nil, err
	}

	pages, err := w.WebsiteRepo.GetPagesByWebsiteID(website.ID)
	if err != nil {
		return nil, errorutil.CheckErrNoRows(err, "no pages found")
	}
	return pages, nil
}

func (w *websiteService) CreatePage(ctx context.Context, p model.PageInput) (*model.Page, error) {
	user, err := userutil.RetrieveUser(ctx, w.UserRepo)
	if err != nil {
		return nil, errorutil.HandleErrorOrNoRows(err, "user not found")
	}

	uuid, err := uuid.NewRandom()
	if err != nil {
		return nil, errorutil.HandleErrorDependingEnv(err)
	}

	page := &model.Page{
		ID:        uuid.String(),
		Name:      p.Name,
		Data:      p.Data,
		Slug:      p.Slug,
		WebsiteID: user.WebsiteID,
	}

	if err := w.WebsiteRepo.CreatePage(page); err != nil {
		if strings.Contains(err.Error(), "#23505") {
			return nil, errors.New("page already exist")
		}
		return nil, errorutil.HandleErrorDependingEnv(err)
	}
	return page, nil
}

func (w *websiteService) DeletePage(ctx context.Context, page *model.Page) (bool, error) {
	isDeleted, err := w.WebsiteRepo.DeletePage(page)

	if err != nil {
		return false, errorutil.HandleErrorDependingEnv(err)
	}

	if !isDeleted {
		return false, errors.New("page not found")
	}
	return true, nil
}

func (w *websiteService) ModifyPage(ctx context.Context, p model.PageInput) (*model.Page, error) {
	timestampz := time.Now().Format(time.RFC3339)

	page := &model.Page{
		ID:        *p.ID,
		Name:      p.Name,
		Slug:      p.Slug,
		Data:      p.Data,
		CreatedAt: timestampz,
	}

	didUpdate, err := w.WebsiteRepo.ModifyPage(page)
	if err != nil {
		return nil, errorutil.HandleErrorDependingEnv(err)
	}

	if !didUpdate {
		return nil, errors.New("page not found")
	}
	return page, nil
}
