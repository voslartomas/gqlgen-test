package main

import (
	"log"
	"net/http"
	"os"

	"github.com/voslartomas/gqlgen-test/cache"
	mongodb "github.com/voslartomas/gqlgen-test/db/mongo"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/voslartomas/gqlgen-test/graph"
	"github.com/voslartomas/gqlgen-test/graph/generated"
)

const defaultPort = "8080"

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	srv := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: &graph.Resolver{}}))

	ctx := mongodb.Connect()
	defer mongodb.Disconnect(ctx)

	cache.Connect()

	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", srv)

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))

	log.Printf("Program exits.")
}
