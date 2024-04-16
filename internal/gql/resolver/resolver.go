//go:generate go run github.com/99designs/gqlgen
package resolver

import (
	"github.com/Revolutionize-org/RevolveCMS-backend/internal/postgres/repository"
	"github.com/Revolutionize-org/RevolveCMS-backend/internal/service/auth"
)

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	AuthService auth.Service
	UserRepo    *repository.UserRepo
	RoleRepo    *repository.RoleRepo
}
