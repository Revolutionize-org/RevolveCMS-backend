//go:generate go run github.com/99designs/gqlgen
package resolver

import "github.com/Revolutionize-org/RevolveCMS-backend/internal/database/postgres"

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	UserRepo  *postgres.UserRepo
	TokenRepo *postgres.TokenRepo
}
