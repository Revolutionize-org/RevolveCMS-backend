package resolver

import (
	"context"
	"errors"
	"net/http"
	"os"
	"time"

	"github.com/99designs/gqlgen/graphql"
	"github.com/Revolutionize-org/RevolveCMS-backend/internal/gql"
	"github.com/Revolutionize-org/RevolveCMS-backend/internal/gql/model"
	"github.com/Revolutionize-org/RevolveCMS-backend/internal/middleware"
	"github.com/Revolutionize-org/RevolveCMS-backend/internal/validation"
	"github.com/go-pg/pg"
	"github.com/golang-jwt/jwt/v5"

	"github.com/vektah/gqlparser/v2/gqlerror"
	"golang.org/x/crypto/bcrypt"
)

type mutationResolver struct{ *Resolver }

func (r *Resolver) Mutation() gql.MutationResolver { return &mutationResolver{r} }

func (r *mutationResolver) Login(ctx context.Context, userInfo model.UserInfo) (*model.AuthToken, error) {
	validErr := validation.ValidateStruct(userInfo)

	for _, err := range validErr {
		graphql.AddError(ctx, &gqlerror.Error{
			Message: err.Error(),
			Extensions: map[string]interface{}{
				"field": err.Field(),
			},
		})
	}

	if validErr != nil {
		return nil, errors.New("invalid input received")
	}

	user, err := r.UserRepo.GetByEmail(userInfo.Email)

	if err != nil {
		if err.Error() == pg.ErrNoRows.Error() {
			return nil, errors.New("invalid email or password")
		}
		return nil, err
	}

	bytePassword := []byte(userInfo.Password)
	byteHashedPassword := []byte(user.PasswordHash)

	if err := bcrypt.CompareHashAndPassword(byteHashedPassword, bytePassword); err != nil {
		return nil, errors.New("invalid email or password")
	}

	accessToken := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.RegisteredClaims{
		ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * 1)),
		IssuedAt:  jwt.NewNumericDate(time.Now()),
		Subject:   user.ID,
		Issuer:    "revolvecms",
	})

	accessTokenSigned, err := accessToken.SignedString([]byte(os.Getenv("ACCESS_TOKEN_SECRET")))

	if err != nil {
		return nil, err
	}

	refreshToken := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.RegisteredClaims{
		ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * 24 * 90)),
		IssuedAt:  jwt.NewNumericDate(time.Now()),
		Subject:   user.ID,
		Issuer:    "revolvecms",
	})

	refreshTokenSigned, err := refreshToken.SignedString([]byte(os.Getenv("REFRESH_TOKEN_SECRET")))

	if err != nil {
		return nil, err
	}

	if err := r.TokenRepo.Add(refreshTokenSigned); err != nil {
		return nil, err
	}

	w, ok := ctx.Value(middleware.ResponseWriterKey{}).(http.ResponseWriter)

	if !ok {
		return nil, errors.New("could not get response writer")
	}

	http.SetCookie(w, &http.Cookie{
		Name:     "refresh_token",
		Value:    refreshTokenSigned,
		Expires:  time.Now().Add(time.Hour * 24 * 90),
		SameSite: http.SameSiteLaxMode,
		HttpOnly: true,
	})

	return &model.AuthToken{
		AccessToken:  accessTokenSigned,
		RefreshToken: refreshTokenSigned,
	}, nil
}
