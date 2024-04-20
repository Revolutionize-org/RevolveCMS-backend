package auth

import (
	"context"
	"errors"

	"github.com/Revolutionize-org/RevolveCMS-backend/internal/cookie"
	"github.com/Revolutionize-org/RevolveCMS-backend/internal/errorutil"
	"github.com/Revolutionize-org/RevolveCMS-backend/internal/jwt"
)

func (a *auth) Logout(ctx context.Context) (bool, error) {
	refreshToken, err := cookie.GetFromContext(ctx, "refresh_token")
	if err != nil {
		return false, nil
	}

	claims, err := jwt.Validate(refreshToken, a.tokenRepo)
	if err != nil {
		return a.deleteRefreshToken(ctx, err)
	}

	jti, ok := claims["jti"].(string)
	if !ok {
		return a.deleteRefreshToken(ctx, errors.New("invalid token"))
	}

	if err := a.deleteTokenByID(jti); err != nil {
		return a.deleteRefreshToken(ctx, err)
	}

	return a.deleteRefreshToken(ctx, nil)
}

func (a *auth) deleteRefreshToken(ctx context.Context, err error) (bool, error) {
	cookieErr := cookie.DeleteFromContext(ctx, "refresh_token")
	if cookieErr != nil {
		return false, cookieErr
	}
	return true, err
}

func (a *auth) deleteTokenByID(jti string) error {
	isDeleted, err := a.tokenRepo.Delete(jti)
	if err != nil {
		return errorutil.HandleErrorDependingEnv(err)
	}

	if !isDeleted {
		return errors.New("token not found")
	}

	return nil
}
