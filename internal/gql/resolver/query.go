package resolver

import (
	"context"
	"errors"
	"fmt"

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

// Themes is the resolver for the themes field.
func (r *queryResolver) Themes(ctx context.Context) ([]*model.Theme, error) {
	return r.ThemeRepo.GetAll()
}

func (r *queryResolver) Website(ctx context.Context) (*model.Website, error) {
	user, err := retrieveUser(ctx, r.UserRepo)
	if err != nil {
		return nil, err
	}
	return r.WebsiteRepo.GetByID(user.WebsiteID)
}

func retrieveUser(ctx context.Context, userRepo *repository.UserRepo) (*model.User, error) {
	userID, ok := ctx.Value(middleware.UserKey{}).(string)
	if !ok {
		return nil, errors.New("could not get user from context")
	}

	return userRepo.GetByID(userID)
}

func (r *queryResolver) Header(ctx context.Context) (*model.Header, error) {
	panic(fmt.Errorf("not implemented: Header - header"))
}

func (r *queryResolver) Pages(ctx context.Context) ([]*model.Page, error) {
	panic(fmt.Errorf("not implemented: Page - page"))
}

func (r *queryResolver) Footer(ctx context.Context) (*model.Footer, error) {
	panic(fmt.Errorf("not implemented: Footer - footer"))
}

func (r *Resolver) Query() gql.QueryResolver { return &queryResolver{r} }

type queryResolver struct{ *Resolver }
