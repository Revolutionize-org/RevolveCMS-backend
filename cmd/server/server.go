package main

import (
	"log"
	"net/http"
	"os"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/Revolutionize-org/RevolveCMS-backend/internal/database/postgres"
	"github.com/Revolutionize-org/RevolveCMS-backend/internal/gql"
	"github.com/Revolutionize-org/RevolveCMS-backend/internal/gql/resolver"
	"github.com/Revolutionize-org/RevolveCMS-backend/internal/middleware"
	"github.com/go-pg/pg/v10"
	"github.com/joho/godotenv"
)

const defaultPort = "5000"

func main() {
	err := godotenv.Load()

	if err != nil {
		log.Fatal("Error loading .env file")
	}
	port := os.Getenv("API_PORT")
	if port == "" {
		port = defaultPort
	}

	db := postgres.New(&pg.Options{
		Addr:       "host.docker.internal:5432",
		User:       os.Getenv("POSTGRES_USER"),
		Password:   os.Getenv("POSTGRES_PASSWORD"),
		Database:   os.Getenv("POSTGRES_DB"),
		MaxRetries: 3,
		PoolSize:   10,
	})
	defer db.Close()

	db.AddQueryHook(postgres.DBLogger{})

	srv := handler.NewDefaultServer(gql.NewExecutableSchema(
		gql.Config{Resolvers: &resolver.Resolver{
			UserRepo:  &postgres.UserRepo{DB: db},
			TokenRepo: &postgres.TokenRepo{DB: db},
		}}))

	middlewareHandler := middleware.Writer(srv)

	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", middlewareHandler)

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
