package main

import (
	"context"
	"crypto/tls"
	"crypto/x509"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/99designs/gqlgen/graphql"
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/handler/extension"
	"github.com/99designs/gqlgen/graphql/handler/transport"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/rs/cors"
	"github.com/vektah/gqlparser/v2/gqlerror"

	"github.com/Revolutionize-org/RevolveCMS-backend/internal/config"
	"github.com/Revolutionize-org/RevolveCMS-backend/internal/database"
	"github.com/Revolutionize-org/RevolveCMS-backend/internal/database/repository"
	"github.com/Revolutionize-org/RevolveCMS-backend/internal/gql"
	"github.com/Revolutionize-org/RevolveCMS-backend/internal/gql/resolver"
	"github.com/Revolutionize-org/RevolveCMS-backend/internal/middleware"

	"github.com/Revolutionize-org/RevolveCMS-backend/internal/service/auth"
	"github.com/Revolutionize-org/RevolveCMS-backend/internal/service/website"
	"github.com/go-pg/pg/v10"
)

func main() {
	db := connectToDB()
	defer db.Close()

	srv := createGraphQLServer(db)

	setupHTTPHandlers(srv)

	log.Printf("Connect to http://localhost:%s/ for GraphQL playground", config.Config.Api.Port)
	log.Fatal(http.ListenAndServe(":"+config.Config.Api.Port, nil))
}

func connectToDB() *pg.DB {
	opts, err := pg.ParseURL(config.Config.Postgres.URL)
	if err != nil {
		panic(err)
	}

	opts.PoolSize = 10
	opts.MaxRetries = 3

	if config.Config.Api.Env != "dev" {
		configTLS(opts)
	}

	db := database.New(opts)

	db.AddQueryHook(database.DBLogger{})
	return db
}

func configTLS(opts *pg.Options) {
	caCeret, err := os.ReadFile("ca.pem")
	if err != nil {
		panic(err)
	}

	caCertPool := x509.NewCertPool()
	caCertPool.AppendCertsFromPEM(caCeret)

	tlsConfig := &tls.Config{
		RootCAs:            caCertPool,
		InsecureSkipVerify: false,
		ServerName:         config.Config.Postgres.Host,
	}
	opts.TLSConfig = tlsConfig
}

func createGraphQLServer(db *pg.DB) http.Handler {
	userRepo := repository.NewUserRepo(db)

	webisteRepo := website.New(repository.NewWebsiteRepo(db), userRepo)

	authService := auth.New(
		userRepo,
		repository.NewTokenRepo(db),
	)

	srv := handler.New(gql.NewExecutableSchema(
		gql.Config{
			Resolvers: &resolver.Resolver{
				AuthService:    authService,
				WebsiteService: webisteRepo,
				UserRepo:       userRepo,
				RoleRepo:       repository.NewRoleRepo(db),
			},
		},
	))

	srv.AddTransport(transport.Options{})
	srv.AddTransport(transport.GET{})
	srv.AddTransport(transport.POST{})
	srv.Use(extension.FixedComplexityLimit(50))

	enableIntrospection(srv)

	return srv
}

func setupHTTPHandlers(srv http.Handler) {
	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"http://localhost:8080", "http://localhost:3001", "https://client-revolve-cms.vercel.app", "https://revolve-cms-frontend.vercel.app/"},
		AllowCredentials: true,
		AllowedMethods:   []string{"GET", "POST", "OPTIONS"},
		AllowedHeaders:   []string{"Authorization", "Content-Type"},
	})

	enablePlayground()

	http.Handle("/graphql", c.Handler(
		middleware.Limiter(
			middleware.Auth(
				middleware.Request(
					middleware.Writer(srv),
				),
			),
		),
	))
}

func enablePlayground() {
	if config.Config.Api.Env == "dev" {
		http.Handle("/", playground.Handler("GraphQL playground", "/graphql"))
	}
}

func enableIntrospection(srv *handler.Server) {
	if config.Config.Api.Env == "dev" {
		srv.Use(extension.Introspection{})
	} else {
		srv.SetErrorPresenter(func(ctx context.Context, e error) *gqlerror.Error {
			err := graphql.DefaultErrorPresenter(ctx, e)

			if strings.Contains(err.Message, "Cannot query field") {
				err.Message = "internal server error"
			}
			return err
		})
	}
}
