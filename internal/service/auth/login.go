package auth

import (
	"context"
	"errors"
	"time"

	"github.com/Revolutionize-org/RevolveCMS-backend/internal/config"
	"github.com/Revolutionize-org/RevolveCMS-backend/internal/cookie"
	"github.com/Revolutionize-org/RevolveCMS-backend/internal/errorutil"

	"github.com/Revolutionize-org/RevolveCMS-backend/internal/gql/model"
	"github.com/Revolutionize-org/RevolveCMS-backend/internal/hashing"
	"github.com/Revolutionize-org/RevolveCMS-backend/internal/jwt"
	"github.com/Revolutionize-org/RevolveCMS-backend/internal/postgres"
)

func (a *auth) Login(ctx context.Context, userInfo model.UserInfo) (*model.AuthToken, error) {
	user, err := a.userRepo.GetByEmail(userInfo.Email)
	if err != nil {
		if err := postgres.CheckErrNoRows(err, "invalid email or password"); err != nil {
			return nil, err
		}
		return nil, errorutil.HandleError(err)
	}

	if err := hashing.CompareHashAndSecret(user.PasswordHash, userInfo.Password); err != nil {
		return nil, errors.New("invalid email or password")
	}

	now := time.Now()
	_, accessToken, err := jwt.New(user.ID, now.Add(time.Hour*1), config.Config.Secret.AccessToken)
	if err != nil {
		return nil, errorutil.HandleError(err)
	}

	rtExp := time.Now().Add(time.Hour * 24 * 90)
	uuid, refreshToken, err := jwt.New(user.ID, rtExp, config.Config.Secret.RefreshToken)
	if err != nil {
		return nil, errorutil.HandleError(err)
	}

	err = a.tokenRepo.Add(uuid, rtExp)
	if err != nil {
		return nil, errorutil.HandleError(err)
	}

	if err = cookie.AddToContext(ctx, "refresh_token", refreshToken, rtExp); err != nil {
		return nil, errorutil.HandleError(err)
	}

	return &model.AuthToken{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}, nil
}
