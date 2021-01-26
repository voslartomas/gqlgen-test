package main

import (
	"log"
	"net/http"
	"os"

	mongodb "github.com/voslartomas/gqlgen-test/db/mongo"

	"github.com/voslartomas/gqlgen-test/graph"
	"github.com/voslartomas/gqlgen-test/graph/generated"
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
)

const defaultPort = "8080"

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	srv := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: &graph.Resolver{}}))

	mongodb.Connect()
	defer mongodb.Disconnect()

	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", srv)

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}