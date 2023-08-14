package main

import (
	"log"
	"net/http"
	"os"
	"skeleton-service/graph"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"go.uber.org/zap"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

const defaultPort = "8080"

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}
	logger, err := zap.NewProduction()
	if err != nil {
		// handle error
		log.Printf("There was an error creating the logger: %v", err)
	}
	defer logger.Sync()

	db, err := gorm.Open(postgres.New(postgres.Config{
		DSN: os.Getenv("PSQL_CONNECTION_STRING"),
	}))

	if err != nil {
		logger.Error("failed to setup db connection ", zap.Error(err))
	} else {
		logger.Info("connected to db")
	}

	srv := handler.NewDefaultServer(graph.NewExecutableSchema(graph.Config{Resolvers: graph.NewResolver(logger, db)}))

	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", srv)

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
