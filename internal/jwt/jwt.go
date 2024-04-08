package jwt

import (
	"errors"
	"fmt"
	"os"
	"time"

	"github.com/Revolutionize-org/RevolveCMS-backend/internal/gql/model"
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
)

func New(data any, secret string) (string, error) {
	var token *jwt.Token

	switch v := data.(type) {
	case *model.User:
		token = jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * 1)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			Subject:   v.ID,
			Issuer:    "revolvecms",
		})
	case uuid.UUID:
		token = jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * 1)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			Subject:   v.String(),
			Issuer:    "revolvecms",
		})
	default:
		return "", errors.New("invalid data type")
	}
	return token.SignedString([]byte(secret))
}

func Parse(t, secret string) (jwt.MapClaims, error) {
	token, err := jwt.Parse(t, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(secret), nil
	})

	if err != nil {
		return nil, err
	}

	if !token.Valid {
		return nil, errors.New("invalid token")
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return claims, nil
	}
	return nil, errors.New("invalid token")
}

func CreateRefreshToken() (string, string, error) {
	uuid, err := uuid.NewRandom()
	if err != nil {
		return "", "", err
	}

	refreshToken, err := New(uuid, os.Getenv("REFRESH_TOKEN_SECRET"))
	return uuid.String(), refreshToken, err
}
