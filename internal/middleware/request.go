package middleware

import (
	"context"
	"net/http"
)

type RequestKey struct{}

func Request(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		ctx = context.WithValue(ctx, RequestKey{}, r)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
