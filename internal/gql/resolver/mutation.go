package resolver

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/Revolutionize-org/RevolveCMS-backend/internal/cookie"
	"github.com/Revolutionize-org/RevolveCMS-backend/internal/errorutil"
	"github.com/Revolutionize-org/RevolveCMS-backend/internal/gql"
	"github.com/Revolutionize-org/RevolveCMS-backend/internal/gql/model"
	"github.com/Revolutionize-org/RevolveCMS-backend/internal/postgres"
	"github.com/Revolutionize-org/RevolveCMS-backend/internal/userutil"
	"github.com/Revolutionize-org/RevolveCMS-backend/internal/validation"
	"github.com/google/uuid"
)

type mutationResolver struct{ *Resolver }

func (r *Resolver) Mutation() gql.MutationResolver { return &mutationResolver{r} }

func (r *mutationResolver) Login(ctx context.Context, userInfo model.UserInfo) (*model.AuthToken, error) {
	_, err := cookie.GetFromContext(ctx, "refresh_token")
	if err == nil {
		return nil, errors.New("you are already logged in")
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

func (r *mutationResolver) CreateHeader(ctx context.Context, h model.HeaderInput) (*model.Header, error) {
	user, err := userutil.RetrieveUser(ctx, r.UserRepo)
	if err != nil {
		if err := postgres.CheckErrNoRows(err, "user not found"); err != nil {
			return nil, err
		}
		return nil, errorutil.HandleError(err)
	}

	uuid, err := uuid.NewRandom()
	if err != nil {
		return nil, errorutil.HandleError(err)
	}

	header := &model.Header{
		ID:        uuid.String(),
		Name:      h.Name,
		Data:      h.Data,
		WebsiteID: user.WebsiteID,
	}

	if err := r.WebsiteService.GetService().WebsiteRepo.CreateHeader(header); err != nil {
		return nil, errorutil.HandleError(err)
	}
	return header, nil
}

func (r *mutationResolver) DeleteHeader(ctx context.Context, id string) (bool, error) {
	isDeleted, err := r.WebsiteService.GetService().WebsiteRepo.DeleteHeader(id)
	if err != nil {
		return false, errorutil.HandleError(err)
	}

	if !isDeleted {
		return false, errors.New("header not found")
	}
	return true, nil
}

func (r *mutationResolver) ModifyHeader(ctx context.Context, h model.HeaderInput) (*model.Header, error) {
	user, err := userutil.RetrieveUser(ctx, r.UserRepo)
	if err != nil {
		if err := postgres.CheckErrNoRows(err, "user not found"); err != nil {
			return nil, err
		}
		return nil, errorutil.HandleError(err)
	}

	timestampz := time.Now().Format(time.RFC3339)

	header := &model.Header{
		ID:        *h.ID,
		Name:      h.Name,
		Data:      h.Data,
		UpdatedAt: timestampz,
		WebsiteID: user.WebsiteID,
	}

	if err := r.WebsiteService.GetService().WebsiteRepo.ModifyHeader(header); err != nil {
		return nil, errorutil.HandleError(err)
	}
	return header, nil
}

func (r *mutationResolver) CreatePage(ctx context.Context, page model.PageInput) (*model.Page, error) {
	panic(fmt.Errorf("not implemented: CreatePage - createPage"))
}

// DeletePage is the resolver for the deletePage field.
func (r *mutationResolver) DeletePage(ctx context.Context, id string) (bool, error) {
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
func (r *mutationResolver) DeleteFooter(ctx context.Context, id string) (bool, error) {
	panic(fmt.Errorf("not implemented: DeleteFooter - deleteFooter"))
}

// ModifyFooter is the resolver for the modifyFooter field.
func (r *mutationResolver) ModifyFooter(ctx context.Context, footer model.FooterInput) (*model.Footer, error) {
	panic(fmt.Errorf("not implemented: ModifyFooter - modifyFooter"))
}

// ModifyWebsiteTheme is the resolver for the modifyWebsiteTheme field.
func (r *mutationResolver) ModifyWebsiteTheme(ctx context.Context, id string, themeID string) (*model.Website, error) {
	panic(fmt.Errorf("not implemented: ModifyWebsiteTheme - modifyWebsiteTheme"))
}
