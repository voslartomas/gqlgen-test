package mongo

import (
	"context"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var dbClient *mongo.Client
var dbCtx context.Context

func Connect() {
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	dbCtx = ctx

	client, err := mongo.NewClient(options.Client().ApplyURI("mongodb://root:test@localhost:27017/?authSource=admin"))
	dbClient = client

	dbClient.Connect(ctx)

	if err != nil {
		panic("DB error")
	}

	fmt.Println("Connected to mongo db.")
}

func Disconnect() {
	dbClient.Disconnect(dbCtx)
	fmt.Println("Disconnected from mongo db.")
}

func GetDatabase() *mongo.Database {
	return dbClient.Database("gqlgentest")
}

func GetContext() context.Context {
	return dbCtx
}
