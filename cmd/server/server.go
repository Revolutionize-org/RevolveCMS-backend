package main

import (
	"context"
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

	"github.com/Revolutionize-org/RevolveCMS-backend/internal/gql"
	"github.com/Revolutionize-org/RevolveCMS-backend/internal/gql/resolver"
	"github.com/Revolutionize-org/RevolveCMS-backend/internal/middleware"
	"github.com/Revolutionize-org/RevolveCMS-backend/internal/postgres"
	"github.com/Revolutionize-org/RevolveCMS-backend/internal/postgres/repository"
	"github.com/Revolutionize-org/RevolveCMS-backend/internal/service/auth"
	"github.com/go-pg/pg/v10"
	"github.com/joho/godotenv"
)

const defaultPort = "5000"

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	port := getPort()
	db := connectToDB()
	defer db.Close()

	srv := createGraphQLServer(db)

	setupHTTPHandlers(srv)

	log.Printf("Connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}

func getPort() string {
	port := os.Getenv("API_PORT")
	if port == "" {
		port = defaultPort
	}
	return port
}

func connectToDB() *pg.DB {
	db := postgres.New(&pg.Options{
		Addr:       "host.docker.internal:5432",
		User:       os.Getenv("POSTGRES_USER"),
		Password:   os.Getenv("POSTGRES_PASSWORD"),
		Database:   os.Getenv("POSTGRES_DB"),
		MaxRetries: 3,
		PoolSize:   10,
	})

	db.AddQueryHook(postgres.DBLogger{})
	return db
}

func createGraphQLServer(db *pg.DB) http.Handler {
	userRepo := repository.NewUserRepo(db)

	authService := auth.New(
		userRepo,
		repository.NewTokenRepo(db),
	)

	srv := handler.New(gql.NewExecutableSchema(
		gql.Config{
			Resolvers: &resolver.Resolver{
				AuthService: authService,
				UserRepo:    userRepo,
				RoleRepo:    repository.NewRoleRepo(db),
				WebsiteRepo: repository.NewWebsiteRepo(db),
			},
		},
	))

	srv.AddTransport(transport.Options{})
	srv.AddTransport(transport.GET{})
	srv.AddTransport(transport.POST{})

	if os.Getenv("ENV") == "dev" {
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

	return srv
}

func setupHTTPHandlers(srv http.Handler) {
	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"http://localhost:8080", "http://localhost:3001"},
		AllowCredentials: true,
		AllowedMethods:   []string{"GET", "POST", "OPTIONS"},
		AllowedHeaders:   []string{"Authorization", "Content-Type"},
	})

	if os.Getenv("ENV") == "dev" {
		http.Handle("/", playground.Handler("GraphQL playground", "/graphql"))
	}

	http.Handle("/graphql", c.Handler(
		middleware.Auth(
			middleware.Request(
				middleware.Writer(srv))),
	))

}
