package middleware

import (
	"context"
	"net/http"
)

type ResponseWriterKey struct{}

func Writer(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		ctx = context.WithValue(ctx, ResponseWriterKey{}, w)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
