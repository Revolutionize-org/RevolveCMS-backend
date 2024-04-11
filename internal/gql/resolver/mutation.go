package resolver

import (
	"context"
	"errors"

	"github.com/Revolutionize-org/RevolveCMS-backend/internal/cookie"
	"github.com/Revolutionize-org/RevolveCMS-backend/internal/gql"
	"github.com/Revolutionize-org/RevolveCMS-backend/internal/gql/model"
	"github.com/Revolutionize-org/RevolveCMS-backend/internal/validation"
)

type mutationResolver struct{ *Resolver }

func (r *Resolver) Mutation() gql.MutationResolver { return &mutationResolver{r} }

func (r *mutationResolver) Login(ctx context.Context, userInfo model.UserInfo) (*model.AuthToken, error) {
	_, err := cookie.GetFromContext(ctx, "refresh_token")
	if err == nil {
		return nil, errors.New("already logged in")
	}

	if err := validation.ValidateInput[model.UserInfo](ctx, userInfo); err != nil {
		return nil, err
	}

	return r.AuthService.Login(ctx, userInfo)
}

func (r *mutationResolver) Logout(ctx context.Context) (bool, error) {
	return r.AuthService.Logout(ctx)
}

func (r *mutationResolver) RefreshToken(ctx context.Context) (string, error) {
	return r.AuthService.RefreshToken(ctx)
}
