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

func (w *websiteService) GetFooter(ctx context.Context) (*model.Footer, error) {
	_, website, err := w.retrieveWebsiteViaCtx(ctx)
	if err != nil {
		return nil, err
	}

	footer, err := w.WebsiteRepo.GetFooterByWebsiteID(website.ID)
	if err != nil {
		return nil, errorutil.HandleError(err, "footer not found")
	}
	return footer, nil
}

func (w *websiteService) CreateFooter(ctx context.Context, f model.FooterInput) (*model.Footer, error) {
	user, err := userutil.RetrieveUser(ctx, w.UserRepo)
	if err != nil {
		return nil, errorutil.HandleError(err, "user not found")
	}

	uuid, err := uuid.NewRandom()
	if err != nil {
		return nil, errorutil.HandleErrorDependingEnv(err)
	}

	footer := &model.Footer{
		ID:        uuid.String(),
		Name:      f.Name,
		Data:      f.Data,
		WebsiteID: user.WebsiteID,
	}

	if err := w.WebsiteRepo.CreateFooter(footer); err != nil {
		if strings.Contains(err.Error(), "#23505") {
			return nil, errors.New("you already have a footer")
		}
		return nil, errorutil.HandleErrorDependingEnv(err)
	}
	return footer, nil
}
func (w *websiteService) DeleteFooter(ctx context.Context, id string) (bool, error) {
	isDeleted, err := w.WebsiteRepo.DeleteFooter(id)

	if err != nil {
		return false, errorutil.HandleErrorDependingEnv(err)
	}

	if !isDeleted {
		return false, errors.New("footer not found")
	}
	return true, nil

}
func (w *websiteService) ModifyFooter(ctx context.Context, f model.FooterInput) (*model.Footer, error) {
	timestampz := time.Now().Format(time.RFC3339)

	footer := &model.Footer{
		ID:        *f.ID,
		Name:      f.Name,
		Data:      f.Data,
		UpdatedAt: timestampz,
	}

	didUpdate, err := w.WebsiteRepo.ModifyFooter(footer)
	if err != nil {
		return nil, errorutil.HandleErrorDependingEnv(err)
	}

	if !didUpdate {
		return nil, errors.New("footer not found")
	}
	return footer, nil
}
