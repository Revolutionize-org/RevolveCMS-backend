package resolver

import (
	"context"

	"github.com/Revolutionize-org/RevolveCMS-backend/internal/gql"
	"github.com/Revolutionize-org/RevolveCMS-backend/internal/gql/model"
)

type websiteResolver struct{ *Resolver }

func (r *Resolver) Website() gql.WebsiteResolver { return &websiteResolver{r} }

func (r *websiteResolver) Theme(ctx context.Context, website *model.Website) (*model.Theme, error) {
	return r.WebsiteRepo.GetThemeByID(website.ThemeID)
}

func (r *websiteResolver) Header(ctx context.Context, website *model.Website) (*model.Header, error) {
	return r.WebsiteRepo.GetHeaderByWebsiteID(website.ID)
}

func (r *websiteResolver) Pages(ctx context.Context, website *model.Website) ([]*model.Page, error) {
	return r.WebsiteRepo.GetPagesByWebsiteID(website.ID)
}

func (r *websiteResolver) Footer(ctx context.Context, website *model.Website) (*model.Footer, error) {
	return r.WebsiteRepo.GetFooterByWebsiteID(website.ID)
}
