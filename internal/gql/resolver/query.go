package resolver

import (
	"context"
	"errors"

	"github.com/Revolutionize-org/RevolveCMS-backend/internal/gql"
	"github.com/Revolutionize-org/RevolveCMS-backend/internal/gql/model"
	"github.com/Revolutionize-org/RevolveCMS-backend/internal/middleware"
	"github.com/Revolutionize-org/RevolveCMS-backend/internal/postgres/repository"
)

func (r *queryResolver) Me(ctx context.Context) (*model.User, error) {
	userID, ok := ctx.Value(middleware.UserKey{}).(string)
	if !ok {
		return nil, errors.New("could not get user from context")
	}

	user, err := r.UserRepo.GetByID(userID)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (r *queryResolver) Themes(ctx context.Context) ([]*model.Theme, error) {
	return r.WebsiteRepo.GetAllThemes()
}

func (r *queryResolver) Website(ctx context.Context) (*model.Website, error) {
	user, err := retrieveUser(ctx, r.UserRepo)
	if err != nil {
		return nil, err
	}
	return r.WebsiteRepo.GetWebsiteByID(user.WebsiteID)
}

func retrieveUser(ctx context.Context, userRepo *repository.UserRepo) (*model.User, error) {
	userID, ok := ctx.Value(middleware.UserKey{}).(string)
	if !ok {
		return nil, errors.New("could not get user from context")
	}

	return userRepo.GetByID(userID)
}

func (r *queryResolver) Header(ctx context.Context) (*model.Header, error) {
	user, err := retrieveUser(ctx, r.UserRepo)
	if err != nil {
		return nil, err
	}

	webiste, err := r.WebsiteRepo.GetWebsiteByID(user.WebsiteID)
	if err != nil {
		return nil, err
	}
	return r.WebsiteRepo.GetHeaderByWebsiteID(webiste.ID)
}

func (r *queryResolver) Pages(ctx context.Context) ([]*model.Page, error) {
	user, err := retrieveUser(ctx, r.UserRepo)
	if err != nil {
		return nil, err
	}

	webiste, err := r.WebsiteRepo.GetWebsiteByID(user.WebsiteID)
	if err != nil {
		return nil, err
	}
	return r.WebsiteRepo.GetPagesByWebsiteID(webiste.ID)
}

func (r *queryResolver) Footer(ctx context.Context) (*model.Footer, error) {
	user, err := retrieveUser(ctx, r.UserRepo)
	if err != nil {
		return nil, err
	}

	webiste, err := r.WebsiteRepo.GetWebsiteByID(user.WebsiteID)
	if err != nil {
		return nil, err
	}
	return r.WebsiteRepo.GetFooterByWebsiteID(webiste.ID)
}

func (r *Resolver) Query() gql.QueryResolver { return &queryResolver{r} }

type queryResolver struct{ *Resolver }
