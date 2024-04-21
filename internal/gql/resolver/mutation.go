package resolver

import (
	"context"

	"github.com/Revolutionize-org/RevolveCMS-backend/internal/cookie"
	"github.com/Revolutionize-org/RevolveCMS-backend/internal/errorutil"
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

	if errs := validation.ValidateInput[model.UserInfo](ctx, userInfo); errs != nil {
		errorutil.AddGraphQLErrors(ctx, errs)
		return nil, nil
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
	if err := validation.ValidateHeaderInput(header); err != nil {
		return nil, err
	}

	if errs := validation.ValidateInput[model.HeaderInput](ctx, header); errs != nil {
		errorutil.AddGraphQLErrors(ctx, errs)
		return nil, nil
	}

	return r.WebsiteService.CreateHeader(ctx, header)
}

func (r *mutationResolver) ModifyHeader(ctx context.Context, header model.HeaderInput) (*model.Header, error) {
	if err := validation.ValidateHeaderInput(header); err != nil {
		return nil, err
	}

	if errs := validation.ValidateInput[model.HeaderInput](ctx, header); errs != nil {
		errorutil.AddGraphQLErrors(ctx, errs)
		return nil, nil
	}

	return r.WebsiteService.ModifyHeader(ctx, header)
}

func (r *mutationResolver) DeleteHeader(ctx context.Context, id string) (bool, error) {
	header := &model.Header{
		ID: id,
	}

	if errs := validation.ValidateInput[*model.Header](ctx, header); errs != nil {
		errorutil.AddGraphQLErrors(ctx, errs)
		return false, nil
	}

	return r.WebsiteService.DeleteHeader(ctx, header)
}

func (r *mutationResolver) CreatePage(ctx context.Context, page model.PageInput) (*model.Page, error) {
	if err := validation.ValidatePageInput(page); err != nil {
		return nil, err
	}

	if errs := validation.ValidateInput[model.PageInput](ctx, page); errs != nil {
		errorutil.AddGraphQLErrors(ctx, errs)
		return nil, nil
	}

	return r.WebsiteService.CreatePage(ctx, page)
}

func (r *mutationResolver) DeletePage(ctx context.Context, id string) (bool, error) {
	page := &model.Page{
		ID: id,
	}

	if errs := validation.ValidateInput[*model.Page](ctx, page); errs != nil {
		errorutil.AddGraphQLErrors(ctx, errs)
		return false, nil
	}

	return r.WebsiteService.DeletePage(ctx, page)
}

func (r *mutationResolver) ModifyPage(ctx context.Context, page model.PageInput) (*model.Page, error) {
	if err := validation.ValidatePageInput(page); err != nil {
		return nil, err
	}

	if errs := validation.ValidateInput[model.PageInput](ctx, page); errs != nil {
		errorutil.AddGraphQLErrors(ctx, errs)
		return nil, nil
	}

	return r.WebsiteService.ModifyPage(ctx, page)
}

func (r *mutationResolver) CreateFooter(ctx context.Context, footer model.FooterInput) (*model.Footer, error) {
	if err := validation.ValidateFooterInput(footer); err != nil {
		return nil, err
	}

	if errs := validation.ValidateInput[model.FooterInput](ctx, footer); errs != nil {
		errorutil.AddGraphQLErrors(ctx, errs)
		return nil, nil
	}

	return r.WebsiteService.CreateFooter(ctx, footer)
}

func (r *mutationResolver) DeleteFooter(ctx context.Context, id string) (bool, error) {
	footer := &model.Footer{
		ID: id,
	}

	if errs := validation.ValidateInput[*model.Footer](ctx, footer); errs != nil {
		errorutil.AddGraphQLErrors(ctx, errs)
		return false, nil
	}

	return r.WebsiteService.DeleteFooter(ctx, footer)
}

func (r *mutationResolver) ModifyFooter(ctx context.Context, footer model.FooterInput) (*model.Footer, error) {
	if err := validation.ValidateFooterInput(footer); err != nil {
		return nil, err
	}

	if errs := validation.ValidateInput[model.FooterInput](ctx, footer); errs != nil {
		errorutil.AddGraphQLErrors(ctx, errs)
		return nil, nil
	}

	return r.WebsiteService.ModifyFooter(ctx, footer)
}

func (r *mutationResolver) ModifyWebsiteTheme(ctx context.Context, id string, themeID string) (*model.Website, error) {
	website := &model.Website{
		ID: id,
	}

	if errs := validation.ValidateInput[*model.Website](ctx, website); errs != nil {
		errorutil.AddGraphQLErrors(ctx, errs)
		return nil, nil
	}

	theme := &model.Theme{
		ID: themeID,
	}

	if errs := validation.ValidateInput[*model.Theme](ctx, theme); errs != nil {
		errorutil.AddGraphQLErrors(ctx, errs)
		return nil, nil
	}

	theme, err := r.WebsiteService.GetService().WebsiteRepo.GetThemeByID(theme.ID)
	if err != nil {
		return nil, err
	}

	return r.WebsiteService.ModifyWebsiteTheme(ctx, website, theme)
}
