package auth

import (
	"context"
	"errors"

	"github.com/Revolutionize-org/RevolveCMS-backend/internal/cookie"
	"github.com/Revolutionize-org/RevolveCMS-backend/internal/jwt"
)

func (a *auth) Logout(ctx context.Context) (bool, error) {
	token, err := cookie.GetFromContext(ctx, "refresh_token")
	if err != nil {
		return false, nil
	}

	claims, err := jwt.Validate(token, a.tokenRepo)
	if err != nil {
		return false, err
	}

	id, ok := claims["id"].(string)
	if !ok {
		return false, errors.New("invalid token")
	}

	deleted, err := a.tokenRepo.Delete(id)
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
