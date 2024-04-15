package resolver

import (
	"context"
	"fmt"

	"github.com/Revolutionize-org/RevolveCMS-backend/internal/gql"
	"github.com/Revolutionize-org/RevolveCMS-backend/internal/gql/model"
)

type userResolver struct{ *Resolver }

func (r *Resolver) User() gql.UserResolver { return &userResolver{r} }

func (r *userResolver) Role(ctx context.Context, obj *model.User) (*model.Role, error) {
	panic(fmt.Errorf("not implemented: Role - role"))
}

func (r *userResolver) Website(ctx context.Context, obj *model.User) (*model.Website, error) {
	panic(fmt.Errorf("not implemented: Website - website"))
}
