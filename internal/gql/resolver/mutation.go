package resolver

import (
	"context"
	"errors"
	"os"
	"time"

	"github.com/Revolutionize-org/RevolveCMS-backend/internal/cookie"
	"github.com/Revolutionize-org/RevolveCMS-backend/internal/gql"
	"github.com/Revolutionize-org/RevolveCMS-backend/internal/gql/model"
	"github.com/Revolutionize-org/RevolveCMS-backend/internal/hashing"
	"github.com/Revolutionize-org/RevolveCMS-backend/internal/jwt"
	"github.com/Revolutionize-org/RevolveCMS-backend/internal/postgres"
	"github.com/Revolutionize-org/RevolveCMS-backend/internal/validation"
)

type mutationResolver struct{ *Resolver }

func (r *Resolver) Mutation() gql.MutationResolver { return &mutationResolver{r} }

func (r *mutationResolver) Login(ctx context.Context, userInfo model.UserInfo) (*model.AuthToken, error) {
	if err := validation.ValidateInput[model.UserInfo](ctx, userInfo); err != nil {
		return nil, err
	}

	user, err := r.UserRepo.GetByEmail(userInfo.Email)
	if err := postgres.CheckErrNoRows(err, "invalid email or password"); err != nil {
		return nil, err
	}

	if err := hashing.CompareHashAndSecret(user.PasswordHash, userInfo.Password); err != nil {
		return nil, errors.New("invalid email or password")
	}

	accessToken, err := jwt.New(user, os.Getenv("ACCESS_TOKEN_SECRET"))
	if err != nil {
		return nil, err
	}

	uuid, refreshToken, err := jwt.CreateRefreshToken()
	if err != nil {
		return nil, err
	}

	if err := r.TokenRepo.Add(uuid, refreshToken); err != nil {
		return nil, err
	}

	if err = cookie.AddToContext(ctx, "refresh_token", refreshToken, time.Now().Add(time.Hour*24*90)); err != nil {
		return nil, err
	}

	return &model.AuthToken{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}, nil
}

func (r *mutationResolver) Logout(ctx context.Context) (bool, error) {
	token, err := cookie.GetFromContext(ctx, "refresh_token")
	if err != nil {
		return false, nil
	}

	id, err := jwt.Validate(token, r.TokenRepo)
	if err != nil {
		return false, err
	}

	deleted, err := r.TokenRepo.Delete(id)
	if err != nil {
		return false, err
	}

	if !deleted {
		return false, errors.New("could not delete token")
	}

	if err := cookie.DeleteFromContext(ctx, "refresh_token"); err != nil {
		return false, nil
	}
	return true, nil
}

func (r *mutationResolver) RefreshToken(ctx context.Context) (string, error) {
	token, err := cookie.GetFromContext(ctx, "refresh_token")
	if err != nil {
		return "", err
	}

	_, err = jwt.Validate(token, r.TokenRepo)
	if err != nil {
		return "", err
	}

	// accessToken, err := jwt.New(id, os.Getenv("ACCESS_TOKEN_SECRET"))
	// if err != nil {
	// 	return "", nil
	// }

	// return accessToken, nil
	return "", nil
}
