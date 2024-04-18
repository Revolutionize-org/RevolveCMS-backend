package website

import (
	"context"

	"github.com/Revolutionize-org/RevolveCMS-backend/internal/gql/model"
	"github.com/Revolutionize-org/RevolveCMS-backend/internal/postgres/repository"
)

type Service interface {
	GetHeader(ctx context.Context) (*model.Header, error)
}

type websiteService struct {
	WebsiteRepo repository.WebsiteRepo
	UserRepo    repository.UserRepo
}

func New(webisteRepo repository.WebsiteRepo, userRepo repository.UserRepo) Service {
	return &websiteService{
		WebsiteRepo: webisteRepo,
		UserRepo:    userRepo,
	}
}
