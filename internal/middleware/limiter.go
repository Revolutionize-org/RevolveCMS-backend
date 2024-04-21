package middleware

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"strings"

	"github.com/Revolutionize-org/RevolveCMS-backend/internal/errorutil"
)

type GraphQLRequestKey struct{}

func Limiter(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		gqlReq, err := readGraphqlRequest(r)
		if err != nil {
			sendError(w, errorutil.HandleErrorDependingEnv(err), http.StatusBadRequest)
			return
		}

		if strings.Count(gqlReq.Query, "login") > 1 {
			sendError(w, errors.New("too many login requests"), http.StatusTooManyRequests)
			return
		}

		ctx := context.WithValue(r.Context(), GraphQLRequestKey{}, gqlReq)

		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func readGraphqlRequest(r *http.Request) (*GraphQLRequest, error) {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		return nil, errors.New("cannot read body")
	}

	r.Body = io.NopCloser(bytes.NewBuffer(body))

	return praseGraphQlRequest(body)
}

func praseGraphQlRequest(body []byte) (*GraphQLRequest, error) {
	var gqlRequest GraphQLRequest
	if err := json.Unmarshal(body, &gqlRequest); err != nil {
		return nil, err
	}
	return &gqlRequest, nil
}
