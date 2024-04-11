package auth

import (
	"context"
	"errors"
	"time"

	"github.com/Revolutionize-org/RevolveCMS-backend/internal/cookie"
	"github.com/Revolutionize-org/RevolveCMS-backend/internal/gql/model"
	"github.com/Revolutionize-org/RevolveCMS-backend/internal/hashing"
	"github.com/Revolutionize-org/RevolveCMS-backend/internal/jwt"
	"github.com/Revolutionize-org/RevolveCMS-backend/internal/postgres"
)

func (a *auth) Login(ctx context.Context, userInfo model.UserInfo) (*model.AuthToken, error) {
	user, err := a.userRepo.GetByEmail(userInfo.Email)
	if err := postgres.CheckErrNoRows(err, "invalid email or password"); err != nil {
		return nil, err
	}

	if err := hashing.CompareHashAndSecret(user.PasswordHash, userInfo.Password); err != nil {
		return nil, errors.New("invalid email or password")
	}

	accessToken, err := jwt.NewAccessToken(user.ID)
	if err != nil {
		return nil, err
	}

	uuid, refreshToken, err := jwt.CreateRefreshToken(user.ID)
	if err != nil {
		return nil, err
	}

	if err := a.tokenRepo.Add(uuid, refreshToken); err != nil {
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
