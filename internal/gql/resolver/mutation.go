package resolver

import (
	"context"
	"errors"
	"os"

	"github.com/Revolutionize-org/RevolveCMS-backend/internal/gql"
	"github.com/Revolutionize-org/RevolveCMS-backend/internal/gql/model"
	"github.com/Revolutionize-org/RevolveCMS-backend/internal/hashing"
	"github.com/Revolutionize-org/RevolveCMS-backend/internal/jwt"
	"github.com/Revolutionize-org/RevolveCMS-backend/internal/postgres"
	"github.com/Revolutionize-org/RevolveCMS-backend/internal/request"
	"github.com/Revolutionize-org/RevolveCMS-backend/internal/validation"
)

type mutationResolver struct{ *Resolver }

func (r *Resolver) Mutation() gql.MutationResolver { return &mutationResolver{r} }

func (r *mutationResolver) Login(ctx context.Context, userInfo model.UserInfo) (*model.AuthToken, error) {
	if err := validation.ValidateInput(ctx, userInfo); err != nil {
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

	if err := r.TokenRepo.Add(uuid, refreshToken); err != nil {
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

func (r *mutationResolver) Logout(ctx context.Context, refreshToken string) (bool, error) {
	claims, err := jwt.Parse(refreshToken, os.Getenv("REFRESH_TOKEN_SECRET"))
	if err != nil {
		return false, err
	}

	tokenId, err := claims.GetSubject()
	if err != nil {
		return false, err
	}

	hashedToken, err := r.TokenRepo.Get(tokenId)
	if err := postgres.CheckErrNoRows(err, "invalid token"); err != nil {
		return false, err
	}

	if err := hashing.CompareHashAndSecret(hashedToken.Token, refreshToken); err != nil {
		return false, errors.New("invalid token")
	}

	deleted, err := r.TokenRepo.Delete(tokenId)
	if err != nil {
		return false, err
	}

	if !deleted {
		return false, errors.New("could not delete token")
	}
	return true, nil
}
