package resolver

import (
	"context"
	"fmt"
	"time"

	"github.com/Revolutionize-org/RevolveCMS-backend/internal/gql"
	"github.com/Revolutionize-org/RevolveCMS-backend/internal/gql/model"
	"github.com/google/uuid"
)

type userResolver struct{ *Resolver }

func (r *Resolver) User() gql.UserResolver { return &userResolver{r} }

// ID is the resolver for the id field.
func (r *userResolver) ID(ctx context.Context, obj *model.User) (uuid.UUID, error) {
	panic(fmt.Errorf("not implemented: ID - id"))
}

// CreatedAt is the resolver for the created_at field.
func (r *userResolver) CreatedAt(ctx context.Context, obj *model.User) (*time.Time, error) {
	panic(fmt.Errorf("not implemented: CreatedAt - created_at"))
}

// Role is the resolver for the role field.
func (r *userResolver) Role(ctx context.Context, obj *model.User) (*model.Role, error) {
	panic(fmt.Errorf("not implemented: Role - role"))
}

// Website is the resolver for the website field.
func (r *userResolver) Website(ctx context.Context, obj *model.User) (*model.Website, error) {
	panic(fmt.Errorf("not implemented: Website - website"))
}
