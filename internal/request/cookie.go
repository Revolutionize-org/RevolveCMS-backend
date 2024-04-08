package request

import (
	"context"
	"errors"
	"net/http"
	"time"

	"github.com/Revolutionize-org/RevolveCMS-backend/internal/middleware"
)

func AddCookieToContext(ctx context.Context, name, value string) error {
	w, ok := ctx.Value(middleware.ResponseWriterKey{}).(http.ResponseWriter)

	if !ok {
		return errors.New("could not get response writer")
	}

	http.SetCookie(w, &http.Cookie{
		Name:     name,
		Value:    value,
		Expires:  time.Now().Add(time.Hour * 24 * 90),
		SameSite: http.SameSiteLaxMode,
		HttpOnly: true,
	})
	return nil
}
