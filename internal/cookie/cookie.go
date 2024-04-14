package cookie

import (
	"context"
	"errors"
	"net/http"
	"os"
	"time"

	"github.com/Revolutionize-org/RevolveCMS-backend/internal/middleware"
)

func AddToContext(ctx context.Context, name, value string, expires time.Time) error {
	w, ok := ctx.Value(middleware.ResponseWriterKey{}).(http.ResponseWriter)

	if !ok {
		return errors.New("could not get response writer")
	}

	var sameSiteMode http.SameSite

	if os.Getenv("ENV") == "production" {
		sameSiteMode = http.SameSiteNoneMode
	} else {
		sameSiteMode = http.SameSiteLaxMode
	}

	http.SetCookie(w, &http.Cookie{
		Name:     name,
		Value:    value,
		Expires:  expires,
		SameSite: sameSiteMode,
		HttpOnly: true,
		Secure:   os.Getenv("ENV") == "production",
	})
	return nil
}

func DeleteFromContext(ctx context.Context, name string) error {
	w, ok := ctx.Value(middleware.ResponseWriterKey{}).(http.ResponseWriter)

	if !ok {
		return errors.New("could not get response writer")
	}

	http.SetCookie(w, &http.Cookie{
		Name:     name,
		Value:    "",
		MaxAge:   -1,
		HttpOnly: true,
	})

	return nil
}

func GetFromContext(ctx context.Context, name string) (string, error) {
	req := ctx.Value(middleware.RequestKey{}).(*http.Request)

	cookie, err := req.Cookie(name)
	if err != nil {
		return "", err
	}

	return cookie.Value, nil
}
