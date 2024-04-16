package auth

import (
	"context"

	"github.com/Revolutionize-org/RevolveCMS-backend/internal/gql/model"
	"github.com/Revolutionize-org/RevolveCMS-backend/internal/postgres/repository"
)

type Service interface {
	Login(ctx context.Context, userInfo model.UserInfo) (*model.AuthToken, error)
	Logout(ctx context.Context) (bool, error)
	RefreshToken(ctx context.Context) (string, error)
}

type auth struct {
	userRepo  *repository.UserRepo
	tokenRepo *repository.TokenRepo
}

func New(userRepo *repository.UserRepo, tokenRepo *repository.TokenRepo) Service {
	return &auth{
		userRepo:  userRepo,
		tokenRepo: tokenRepo,
	}
}
