package resolver

import (
	"context"
	"fmt"

	"github.com/Revolutionize-org/RevolveCMS-backend/internal/gql"
	"github.com/Revolutionize-org/RevolveCMS-backend/internal/gql/model"
)

type websiteResolver struct{ *Resolver }

func (r *Resolver) Website() gql.WebsiteResolver { return &websiteResolver{r} }

func (r *websiteResolver) Theme(ctx context.Context, website *model.Website) (*model.Theme, error) {
	return r.ThemeRepo.GetByID(website.ThemeID)
}

func (r *websiteResolver) Header(ctx context.Context, website *model.Website) (*model.Header, error) {
	return r.HeaderRepo.GetByWebsiteID(website.ID)
}

// Pages is the resolver for the pages field.
func (r *websiteResolver) Pages(ctx context.Context, obj *model.Website) ([]*model.Page, error) {
	panic(fmt.Errorf("not implemented: Pages - pages"))
}

// Footer is the resolver for the footer field.
func (r *websiteResolver) Footer(ctx context.Context, obj *model.Website) (*model.Footer, error) {
	panic(fmt.Errorf("not implemented: Footer - footer"))
}
