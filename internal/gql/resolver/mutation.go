package resolver

import (
	"context"
	"errors"
	"fmt"

	"github.com/Revolutionize-org/RevolveCMS-backend/internal/cookie"
	"github.com/Revolutionize-org/RevolveCMS-backend/internal/gql"
	"github.com/Revolutionize-org/RevolveCMS-backend/internal/gql/model"
	"github.com/Revolutionize-org/RevolveCMS-backend/internal/validation"
	"github.com/google/uuid"
)

type mutationResolver struct{ *Resolver }

func (r *Resolver) Mutation() gql.MutationResolver { return &mutationResolver{r} }

func (r *mutationResolver) Login(ctx context.Context, userInfo model.UserInfo) (*model.AuthToken, error) {
	_, err := cookie.GetFromContext(ctx, "refresh_token")
	if err == nil {
		return nil, errors.New("already logged in")
	}

	if err := validation.ValidateInput[model.UserInfo](ctx, userInfo); err != nil {
		return nil, err
	}

	return r.AuthService.Login(ctx, userInfo)
}

func (r *mutationResolver) Logout(ctx context.Context) (bool, error) {
	return r.AuthService.Logout(ctx)
}

func (r *mutationResolver) RefreshToken(ctx context.Context) (string, error) {
	return r.AuthService.RefreshToken(ctx)
}

func (r *mutationResolver) CreateHeader(ctx context.Context, header model.HeaderInput) (*model.Header, error) {
	panic(fmt.Errorf("not implemented: CreateHeader - createHeader"))
}

// DeleteHeader is the resolver for the deleteHeader field.
func (r *mutationResolver) DeleteHeader(ctx context.Context, id uuid.UUID) (bool, error) {
	panic(fmt.Errorf("not implemented: DeleteHeader - deleteHeader"))
}

// ModifyHeader is the resolver for the modifyHeader field.
func (r *mutationResolver) ModifyHeader(ctx context.Context, header model.HeaderInput) (*model.Header, error) {
	panic(fmt.Errorf("not implemented: ModifyHeader - modifyHeader"))
}

// CreatePage is the resolver for the createPage field.
func (r *mutationResolver) CreatePage(ctx context.Context, page model.PageInput) (*model.Page, error) {
	panic(fmt.Errorf("not implemented: CreatePage - createPage"))
}

// DeletePage is the resolver for the deletePage field.
func (r *mutationResolver) DeletePage(ctx context.Context, id uuid.UUID) (bool, error) {
	panic(fmt.Errorf("not implemented: DeletePage - deletePage"))
}

// ModifyPage is the resolver for the modifyPage field.
func (r *mutationResolver) ModifyPage(ctx context.Context, page model.PageInput) (*model.Page, error) {
	panic(fmt.Errorf("not implemented: ModifyPage - modifyPage"))
}

// CreateFooter is the resolver for the createFooter field.
func (r *mutationResolver) CreateFooter(ctx context.Context, footer model.FooterInput) (*model.Footer, error) {
	panic(fmt.Errorf("not implemented: CreateFooter - createFooter"))
}

// DeleteFooter is the resolver for the deleteFooter field.
func (r *mutationResolver) DeleteFooter(ctx context.Context, id uuid.UUID) (bool, error) {
	panic(fmt.Errorf("not implemented: DeleteFooter - deleteFooter"))
}

// ModifyFooter is the resolver for the modifyFooter field.
func (r *mutationResolver) ModifyFooter(ctx context.Context, footer model.FooterInput) (*model.Footer, error) {
	panic(fmt.Errorf("not implemented: ModifyFooter - modifyFooter"))
}

// ModifyWebsiteTheme is the resolver for the modifyWebsiteTheme field.
func (r *mutationResolver) ModifyWebsiteTheme(ctx context.Context, id uuid.UUID, themeID uuid.UUID) (*model.Website, error) {
	panic(fmt.Errorf("not implemented: ModifyWebsiteTheme - modifyWebsiteTheme"))
}
