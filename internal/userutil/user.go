package userutil

import (
	"context"
	"errors"

	"github.com/Revolutionize-org/RevolveCMS-backend/internal/database/repository"
	"github.com/Revolutionize-org/RevolveCMS-backend/internal/gql/model"
	"github.com/Revolutionize-org/RevolveCMS-backend/internal/middleware"
)

func RetrieveUser(ctx context.Context, userRepo repository.UserRepo) (*model.User, error) {
	userID, ok := ctx.Value(middleware.UserKey{}).(string)
	if !ok {

		return nil, errors.New("could not get user from context")
	}
	return userRepo.GetByID(userID)
}
