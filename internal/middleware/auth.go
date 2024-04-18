package middleware

import (
	"bytes"
	"context"
	"encoding/json"
	"io"
	"net/http"
	"os"

	"github.com/golang-jwt/jwt/v5"
	"github.com/golang-jwt/jwt/v5/request"
)

// Constants for operation names that do not require authentication
const (
	OpLogin        = "Login"
	OpLogout       = "Logout"
	OpRefreshToken = "RefreshToken"
)

// GraphQLRequest represents a GraphQL request structure
type GraphQLRequest struct {
	Query         string                 `json:"query"`
	Variables     map[string]interface{} `json:"variables"`
	OperationName string                 `json:"operationName"`
}

type UserKey struct{}

func Auth(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		body, err := io.ReadAll(r.Body)
		if err != nil {
			http.Error(w, "Failed to read request body: "+err.Error(), http.StatusBadRequest)
			return
		}

		r.Body = io.NopCloser(bytes.NewBuffer(body))

		gqlRequest, err := parseGraphQLRequest(body)
		if err != nil {
			sendError(w, err, http.StatusBadRequest)
			return
		}

		if operationExemptFromAuth(gqlRequest.OperationName) {
			next.ServeHTTP(w, r)
			return
		}

		claims, err := validateToken(r)
		if err != nil {
			sendError(w, err, http.StatusUnauthorized)
			return
		}

		ctx := context.WithValue(r.Context(), UserKey{}, claims["sub"])
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func parseGraphQLRequest(body []byte) (*GraphQLRequest, error) {
	var gqlRequest GraphQLRequest
	if err := json.Unmarshal(body, &gqlRequest); err != nil {
		return nil, err
	}
	return &gqlRequest, nil
}

func operationExemptFromAuth(operationName string) bool {
	return operationName == OpLogin || operationName == OpLogout || operationName == OpRefreshToken
}

func validateToken(r *http.Request) (jwt.MapClaims, error) {
	token, err := request.ParseFromRequest(r, request.AuthorizationHeaderExtractor, func(token *jwt.Token) (interface{}, error) {
		secretKey := os.Getenv("ACCESS_TOKEN_SECRET")
		return []byte(secretKey), nil
	})
	if err != nil {
		return nil, err
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok || !token.Valid {
		return nil, err
	}

	return claims, nil
}

func sendError(w http.ResponseWriter, err error, status int) {
	errorResponse := map[string]string{
		"message": err.Error(),
	}

	w.Header().Set("Content-Type", "application/json")

	w.WriteHeader(status)
	json.NewEncoder(w).Encode(errorResponse)
}
