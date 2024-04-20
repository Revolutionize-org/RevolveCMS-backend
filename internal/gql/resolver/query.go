package resolver

import (
	"context"
	"errors"

	"github.com/Revolutionize-org/RevolveCMS-backend/internal/errorutil"
	"github.com/Revolutionize-org/RevolveCMS-backend/internal/gql"
	"github.com/Revolutionize-org/RevolveCMS-backend/internal/gql/model"
	"github.com/Revolutionize-org/RevolveCMS-backend/internal/middleware"
)

func (r *queryResolver) Me(ctx context.Context) (*model.User, error) {
	userID, ok := ctx.Value(middleware.UserKey{}).(string)
	if !ok {
		return nil, errors.New("no user provided")
	}

	user, err := r.UserRepo.GetByID(userID)
	if err != nil {
		return nil, errorutil.HandleErrorOrNoRows(err, "user not found")
	}

	return user, nil
}

func (r *queryResolver) Themes(ctx context.Context) ([]*model.Theme, error) {
	return r.WebsiteService.GetService().WebsiteRepo.GetAllThemes()
}

func (r *queryResolver) Website(ctx context.Context) (*model.Website, error) {
	return r.WebsiteService.GetWebsite(ctx)
}

func (r *queryResolver) Header(ctx context.Context) (*model.Header, error) {
	return r.WebsiteService.GetHeader(ctx)
}

func (r *queryResolver) Pages(ctx context.Context) ([]*model.Page, error) {
	return r.WebsiteService.GetPages(ctx)
}

func (r *queryResolver) Footer(ctx context.Context) (*model.Footer, error) {
	return r.WebsiteService.GetFooter(ctx)
}

func (r *Resolver) Query() gql.QueryResolver { return &queryResolver{r} }

type queryResolver struct{ *Resolver }
