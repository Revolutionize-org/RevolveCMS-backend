package jwt

import (
	"errors"
	"fmt"
	"os"
	"time"

	"github.com/Revolutionize-org/RevolveCMS-backend/internal/postgres"
	"github.com/Revolutionize-org/RevolveCMS-backend/internal/postgres/repository"
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
)

type RefreshTokenPayload struct {
	ID     uuid.UUID
	UserID string
}

func generateTokenWithClaims(claims jwt.Claims, secret string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(secret))
}

func New(subject string, expiration time.Time, secret string) (string, string, error) {
	uuid, err := uuid.NewRandom()
	if err != nil {
		return "", "", err
	}

	now := time.Now()
	claims := jwt.MapClaims{
		"jti": uuid.String(),
		"sub": subject,
		"exp": jwt.NewNumericDate(expiration),
		"iat": jwt.NewNumericDate(now),
		"iss": "revolvecms",
	}

	token, err := generateTokenWithClaims(claims, secret)
	return uuid.String(), token, err
}

func parse(t, secret string) (jwt.MapClaims, error) {
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

func Validate(t string, tokenRepo *repository.TokenRepo) (jwt.MapClaims, error) {
	claims, err := parse(t, os.Getenv("REFRESH_TOKEN_SECRET"))
	if err != nil {
		return nil, err
	}

	jti, ok := claims["jti"].(string)
	if !ok {
		return nil, errors.New("invalid token")
	}

	token, err := tokenRepo.Get(jti)
	if err := postgres.CheckErrNoRows(err, "invalid token"); err != nil {
		return nil, err
	}

	if token.ExpiresAt.Before(time.Now()) {
		return nil, errors.New("invalid token")
	}

	return claims, nil
}
