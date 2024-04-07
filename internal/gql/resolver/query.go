package resolver

import (
	"context"
	"fmt"

	"github.com/Revolutionize-org/RevolveCMS-backend/internal/gql"
	"github.com/Revolutionize-org/RevolveCMS-backend/internal/gql/model"
)

func (r *queryResolver) Me(ctx context.Context, id string) (*model.User, error) {
	panic(fmt.Errorf("not implemented: Me - me"))
}

// Mutation returns gql.MutationResolver implementation.

// Query returns gql.QueryResolver implementation.
func (r *Resolver) Query() gql.QueryResolver { return &queryResolver{r} }

type queryResolver struct{ *Resolver }
