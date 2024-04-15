package auth

import (
	"context"
	"errors"
	"os"
	"time"

	"github.com/Revolutionize-org/RevolveCMS-backend/internal/cookie"
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

	_, accessToken, err := jwt.New(subject, time.Now().Add(time.Hour*1), os.Getenv("ACCESS_TOKEN_SECRET"))
	if err != nil {
		return "", err
	}

	return accessToken, nil
}
