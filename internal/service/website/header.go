package website

import (
	"context"

	"github.com/Revolutionize-org/RevolveCMS-backend/internal/gql/model"
	"github.com/Revolutionize-org/RevolveCMS-backend/internal/userutil"
)

func (w *websiteService) GetHeader(ctx context.Context) (*model.Header, error) {
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
