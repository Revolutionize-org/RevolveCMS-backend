package resolver

import (
	"context"
	"fmt"

	"github.com/Revolutionize-org/RevolveCMS-backend/internal/gql"
	"github.com/Revolutionize-org/RevolveCMS-backend/internal/gql/model"
)

type userResolver struct{ *Resolver }

func (r *Resolver) User() gql.UserResolver { return &userResolver{r} }

func (r *userResolver) Role(ctx context.Context, user *model.User) (*model.Role, error) {
	role, err := r.RoleRepo.GetByID(user.RoleID)
	if err != nil {
		return nil, err
	}
	return &model.Role{ID: role.ID, Name: role.Name}, nil
}

func (r *userResolver) Website(ctx context.Context, user *model.User) (*model.Website, error) {
	panic(fmt.Errorf("not implemented: Website - website"))
}
