package resolver

import (
	"context"
	"fmt"

	"github.com/Revolutionize-org/RevolveCMS-backend/internal/gql"
	"github.com/Revolutionize-org/RevolveCMS-backend/internal/gql/model"
)

func (r *queryResolver) Me(ctx context.Context) (*model.User, error) {
	panic(fmt.Errorf("not implemented: Me - me"))
}

// Website is the resolver for the website field.
func (r *queryResolver) Website(ctx context.Context) (*model.Website, error) {
	panic(fmt.Errorf("not implemented: Website - website"))
}

// Header is the resolver for the header field.
func (r *queryResolver) Header(ctx context.Context, websiteID string) (*model.Header, error) {
	panic(fmt.Errorf("not implemented: Header - header"))
}

// Page is the resolver for the page field.
func (r *queryResolver) Page(ctx context.Context, websiteID string) ([]*model.Page, error) {
	panic(fmt.Errorf("not implemented: Page - page"))
}

// Footer is the resolver for the footer field.
func (r *queryResolver) Footer(ctx context.Context, websiteID string) (*model.Footer, error) {
	panic(fmt.Errorf("not implemented: Footer - footer"))
}

// Mutation returns gql.MutationResolver implementation.

// Query returns gql.QueryResolver implementation.
func (r *Resolver) Query() gql.QueryResolver { return &queryResolver{r} }

type queryResolver struct{ *Resolver }
