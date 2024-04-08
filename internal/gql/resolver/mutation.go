package resolver

import (
	"context"
	"errors"
	"os"

	"github.com/Revolutionize-org/RevolveCMS-backend/internal/gql"
	"github.com/Revolutionize-org/RevolveCMS-backend/internal/gql/model"
	"github.com/Revolutionize-org/RevolveCMS-backend/internal/hashing"
	"github.com/Revolutionize-org/RevolveCMS-backend/internal/jwt"
	"github.com/Revolutionize-org/RevolveCMS-backend/internal/request"
	"github.com/Revolutionize-org/RevolveCMS-backend/internal/validation"
	"github.com/go-pg/pg"
)

type mutationResolver struct{ *Resolver }

func (r *Resolver) Mutation() gql.MutationResolver { return &mutationResolver{r} }

func (r *mutationResolver) Login(ctx context.Context, userInfo model.UserInfo) (*model.AuthToken, error) {
	if err := validation.ValidateInput(ctx, userInfo); err != nil {
		return nil, err
	}

	user, err := r.UserRepo.GetByEmail(userInfo.Email)
	if err != nil {
		if err.Error() == pg.ErrNoRows.Error() {
			return nil, errors.New("invalid email or password")
		}
		return nil, err
	}

	if err := hashing.CompareHashAndPassword(user.PasswordHash, userInfo.Password); err != nil {
		return nil, errors.New("invalid email or password")

	}

	accessToken, err := jwt.New(user, os.Getenv("ACCESS_TOKEN_SECRET"))
	if err != nil {
		return nil, err
	}

	refreshToken, err := jwt.New(user, os.Getenv("REFRESH_TOKEN_SECRET"))
	if err != nil {
		return nil, err
	}

	if err := r.TokenRepo.Add(refreshToken); err != nil {
		return nil, err
	}

	if err = request.AddCookieToContext(ctx, "refresh_token", refreshToken); err != nil {
		return nil, err

	}

	return &model.AuthToken{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}, nil
}
