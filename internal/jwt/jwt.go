package jwt

import (
	"errors"
	"fmt"
	"os"
	"time"

	"github.com/Revolutionize-org/RevolveCMS-backend/internal/hashing"
	"github.com/Revolutionize-org/RevolveCMS-backend/internal/postgres"
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

func NewAccessToken(userID string) (string, error) {
	now := time.Now()
	claims := jwt.MapClaims{
		"id":  userID,
		"exp": jwt.NewNumericDate(now.Add(time.Hour * 1)),
		"iat": jwt.NewNumericDate(now),
		"iss": "revolvecms",
	}

	token, err := generateTokenWithClaims(claims, os.Getenv("ACCESS_TOKEN_SECRET"))
	return token, err
}

func NewRefreshToken(payload RefreshTokenPayload) (string, error) {
	now := time.Now()
	claims := jwt.MapClaims{
		"id":     payload.ID,
		"userID": payload.UserID,
		"exp":    jwt.NewNumericDate(now.Add(time.Hour * 24 * 90)),
		"iat":    jwt.NewNumericDate(now),
		"iss":    "revolvecms",
	}

	token, err := generateTokenWithClaims(claims, os.Getenv("REFRESH_TOKEN_SECRET"))
	return token, err
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

func CreateRefreshToken(userID string) (string, string, error) {
	uuid, err := uuid.NewRandom()
	if err != nil {
		return "", "", err
	}

	payload := RefreshTokenPayload{
		ID:     uuid,
		UserID: userID,
	}
	refreshToken, err := NewRefreshToken(payload)
	return uuid.String(), refreshToken, err
}

func Validate(token string, tokenRepo *postgres.TokenRepo) (jwt.MapClaims, error) {
	claims, err := Parse(token, os.Getenv("REFRESH_TOKEN_SECRET"))
	if err != nil {
		return nil, err
	}

	tokenId, ok := claims["id"].(string)
	if !ok {
		return nil, errors.New("invalid token")
	}

	hashedToken, err := tokenRepo.Get(tokenId)
	if err := postgres.CheckErrNoRows(err, "invalid token"); err != nil {
		return nil, err
	}

	if err := hashing.CompareHashAndSecret(hashedToken.Token, token); err != nil {
		return nil, errors.New("invalid token")
	}

	return claims, nil
}
