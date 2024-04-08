package jwt

import (
	"time"

	"github.com/Revolutionize-org/RevolveCMS-backend/internal/gql/model"
	"github.com/golang-jwt/jwt/v5"
)

func New(user *model.User, secret string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.RegisteredClaims{
		ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * 1)),
		IssuedAt:  jwt.NewNumericDate(time.Now()),
		Subject:   user.ID,
		Issuer:    "revolvecms",
	})
	return token.SignedString([]byte(secret))
}
