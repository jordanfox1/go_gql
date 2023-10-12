package main

import (
	"fmt"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/rs/cors"
	"go_gql/graph"
	custom_middleware "go_gql/middleware"
	"go_gql/postgres"
	"log"
	"net/http"
	"os"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/go-pg/pg/v10"
)

func main() {
	DB := postgres.NewDB(&pg.Options{
		User:     "postgres",
		Password: "password",
		Database: "meetmeup_dev",
	})

	defer DB.Close()
	DB.AddQueryHook(postgres.DBLogger{})

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	// We only want ONE instance of this
	userRepo := postgres.UsersRepo{DB: DB}

	router := chi.NewRouter()

	router.Use(cors.New(cors.Options{
		AllowedOrigins:   []string{fmt.Sprintf("http://localhost:%s", port)},
		AllowCredentials: true,
		Debug:            true,
	}).Handler)
	router.Use(middleware.RequestID)
	router.Use(middleware.Logger)
	router.Use(custom_middleware.AuthMiddleware(userRepo))

	config := graph.Config{Resolvers: &graph.Resolver{
		MeetupsRepo: postgres.MeetupsRepo{DB: DB},
		UsersRepo:   userRepo,
	}}

	srv := handler.NewDefaultServer(graph.NewExecutableSchema(config))

	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", graph.DataloaderMiddleware(DB, srv))

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
