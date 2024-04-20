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
)

func (a *auth) Login(ctx context.Context, userInfo model.UserInfo) (*model.AuthToken, error) {
	user, err := a.userRepo.GetByEmail(userInfo.Email)
	if err != nil {
		return nil, errorutil.HandleErrorOrNoRows(err, "invalid email or password")
	}

	if err := hashing.CompareHashAndSecret(user.PasswordHash, userInfo.Password); err != nil {
		return nil, errors.New("invalid email or password")
	}

	now := time.Now()
	_, accessToken, err := jwt.New(user.ID, now.Add(time.Hour*1), config.Config.Secret.AccessToken)
	if err != nil {
		return nil, errorutil.HandleErrorDependingEnv(err)
	}

	rtExp := time.Now().Add(time.Hour * 24 * 90)
	uuid, refreshToken, err := jwt.New(user.ID, rtExp, config.Config.Secret.RefreshToken)
	if err != nil {
		return nil, errorutil.HandleErrorDependingEnv(err)
	}

	err = a.tokenRepo.Add(uuid, rtExp)
	if err != nil {
		return nil, errorutil.HandleErrorDependingEnv(err)
	}

	if err = cookie.AddToContext(ctx, "refresh_token", refreshToken, rtExp); err != nil {
		return nil, errorutil.HandleErrorDependingEnv(err)
	}

	return &model.AuthToken{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}, nil
}
