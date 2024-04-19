package auth

import (
	"context"
	"errors"
	"time"

	"github.com/Revolutionize-org/RevolveCMS-backend/internal/config"
	"github.com/Revolutionize-org/RevolveCMS-backend/internal/cookie"
	"github.com/Revolutionize-org/RevolveCMS-backend/internal/errorutil"
	"github.com/Revolutionize-org/RevolveCMS-backend/internal/jwt"
)

func (a *auth) RefreshToken(ctx context.Context) (string, error) {
	token, err := cookie.GetFromContext(ctx, "refresh_token")
	if err != nil {
		return "", err
	}

	claims, err := jwt.Validate(token, a.tokenRepo)
	if err != nil {
		if err := cookie.DeleteFromContext(ctx, "refresh_token"); err != nil {
			return "", err
		}
		return "", err
	}

	subject, ok := claims["sub"].(string)
	if !ok {
		return "", errors.New("invalid token")
	}

	_, accessToken, err := jwt.New(subject, time.Now().Add(time.Hour*1), config.Config.Secret.AccessToken)
	if err != nil {
		return "", errorutil.HandleError(err)
	}

	return accessToken, nil
}
