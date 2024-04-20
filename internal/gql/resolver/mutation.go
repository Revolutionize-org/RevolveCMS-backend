package resolver

import (
	"context"

	"github.com/Revolutionize-org/RevolveCMS-backend/internal/cookie"
	"github.com/Revolutionize-org/RevolveCMS-backend/internal/gql"
	"github.com/Revolutionize-org/RevolveCMS-backend/internal/gql/model"
	"github.com/Revolutionize-org/RevolveCMS-backend/internal/validation"
	"github.com/vektah/gqlparser/v2/gqlerror"
)

type mutationResolver struct{ *Resolver }

func (r *Resolver) Mutation() gql.MutationResolver { return &mutationResolver{r} }

func (r *mutationResolver) Login(ctx context.Context, userInfo model.UserInfo) (*model.AuthToken, error) {
	_, err := cookie.GetFromContext(ctx, "refresh_token")
	if err == nil {
		return nil, gqlerror.Errorf("Already logged in")
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
	if err := validation.ValidateInput[model.HeaderInput](ctx, header); err != nil {
		return nil, err
	}

	return r.WebsiteService.CreateHeader(ctx, header)
}

func (r *mutationResolver) ModifyHeader(ctx context.Context, header model.HeaderInput) (*model.Header, error) {
	if err := validation.ValidateInput[model.HeaderInput](ctx, header); err != nil {
		return nil, err
	}

	return r.WebsiteService.ModifyHeader(ctx, header)
}

func (r *mutationResolver) DeleteHeader(ctx context.Context, id string) (bool, error) {
	return r.WebsiteService.DeleteHeader(ctx, id)
}

func (r *mutationResolver) CreatePage(ctx context.Context, page model.PageInput) (*model.Page, error) {
	if err := validation.ValidateInput[model.PageInput](ctx, page); err != nil {
		return nil, err
	}

	return r.WebsiteService.CreatePage(ctx, page)
}

func (r *mutationResolver) DeletePage(ctx context.Context, id string) (bool, error) {
	return r.WebsiteService.DeletePage(ctx, id)
}

func (r *mutationResolver) ModifyPage(ctx context.Context, page model.PageInput) (*model.Page, error) {
	if err := validation.ValidateInput[model.PageInput](ctx, page); err != nil {
		return nil, err
	}

	return r.WebsiteService.ModifyPage(ctx, page)
}

func (r *mutationResolver) CreateFooter(ctx context.Context, footer model.FooterInput) (*model.Footer, error) {
	if err := validation.ValidateInput[model.FooterInput](ctx, footer); err != nil {
		return nil, err
	}

	return r.WebsiteService.CreateFooter(ctx, footer)
}

func (r *mutationResolver) DeleteFooter(ctx context.Context, id string) (bool, error) {
	return r.WebsiteService.DeleteFooter(ctx, id)
}

func (r *mutationResolver) ModifyFooter(ctx context.Context, footer model.FooterInput) (*model.Footer, error) {
	if err := validation.ValidateInput[model.FooterInput](ctx, footer); err != nil {
		return nil, err
	}
	return r.WebsiteService.ModifyFooter(ctx, footer)
}

func (r *mutationResolver) ModifyWebsiteTheme(ctx context.Context, id string, themeID string) (*model.Website, error) {
	return r.WebsiteService.ModifyWebsiteTheme(ctx, id, themeID)
}
