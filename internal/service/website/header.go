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

func (w *websiteService) GetHeader(ctx context.Context) (*model.Header, error) {
	_, website, err := w.retrieveWebsiteViaCtx(ctx)
	if err != nil {
		return nil, err
	}

	header, err := w.WebsiteRepo.GetHeaderByWebsiteID(website.ID)
	if err != nil {
		return nil, errorutil.CheckErrNoRows(err, "header not found")
	}
	return header, nil
}

func (w *websiteService) CreateHeader(ctx context.Context, h model.HeaderInput) (*model.Header, error) {
	user, err := userutil.RetrieveUser(ctx, w.UserRepo)
	if err != nil {
		return nil, errorutil.HandleErrorOrNoRows(err, "user not found")
	}

	uuid, err := uuid.NewRandom()
	if err != nil {
		return nil, errorutil.HandleErrorDependingEnv(err)
	}

	header := &model.Header{
		ID:        uuid.String(),
		Name:      h.Name,
		Data:      h.Data,
		WebsiteID: user.WebsiteID,
	}

	if err := w.WebsiteRepo.CreateHeader(header); err != nil {
		if strings.Contains(err.Error(), "#23505") {
			return nil, errors.New("you already have an header")
		}
		return nil, errorutil.HandleErrorDependingEnv(err)
	}
	return header, nil
}

func (w *websiteService) ModifyHeader(ctx context.Context, h model.HeaderInput) (*model.Header, error) {
	timestampz := time.Now().Format(time.RFC3339)

	header := &model.Header{
		ID:        *h.ID,
		Name:      h.Name,
		Data:      h.Data,
		UpdatedAt: timestampz,
	}

	didUpdate, err := w.WebsiteRepo.ModifyHeader(header)
	if err != nil {
		return nil, errorutil.HandleErrorDependingEnv(err)
	}

	if !didUpdate {
		return nil, errors.New("header not found")
	}
	return header, nil
}

func (w *websiteService) DeleteHeader(ctx context.Context, header *model.Header) (bool, error) {
	isDeleted, err := w.WebsiteRepo.DeleteHeader(header)

	if err != nil {
		return false, errorutil.HandleErrorDependingEnv(err)
	}

	if !isDeleted {
		return false, errors.New("header not found")
	}
	return true, nil
}
