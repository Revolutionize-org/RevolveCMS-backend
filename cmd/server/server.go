package main

import (
	"log"
	"net/http"
	"os"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/rs/cors"

	"github.com/Revolutionize-org/RevolveCMS-backend/internal/gql"
	"github.com/Revolutionize-org/RevolveCMS-backend/internal/gql/resolver"
	"github.com/Revolutionize-org/RevolveCMS-backend/internal/middleware"
	"github.com/Revolutionize-org/RevolveCMS-backend/internal/postgres"
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
	authService := auth.New(&postgres.UserRepo{DB: db}, &postgres.TokenRepo{DB: db})

	return middleware.Writer(handler.NewDefaultServer(gql.NewExecutableSchema(
		gql.Config{
			Resolvers: &resolver.Resolver{
				AuthService: authService,
			},
		},
	)))
}

func setupHTTPHandlers(srv http.Handler) {
	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"http://localhost:8080", "http://localhost:3001"},
		AllowCredentials: true,
		AllowedMethods:   []string{"GET", "POST", "OPTIONS"},
	})
	http.Handle("/", playground.Handler("GraphQL playground", "/graphql"))
	http.Handle("/graphql", c.Handler(
		middleware.Request(
			middleware.Writer(srv))))
}
