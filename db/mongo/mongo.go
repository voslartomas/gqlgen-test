package mongo

import (
	"context"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var dbClient *mongo.Client

func Connect() context.Context {
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)

	client, err := mongo.NewClient(options.Client().ApplyURI("mongodb://root:test@localhost:27017/?authSource=admin"))
	dbClient = client

	dbClient.Connect(ctx)

	if err != nil {
		panic("DB error")
	}

	log.Println("Connected to mongo db.")

	return ctx
}

func Disconnect(ctx context.Context) {
	dbClient.Disconnect(ctx)
	log.Println("Disconnected from mongo db.")
}

func GetDatabase() *mongo.Database {
	return dbClient.Database("gqlgentest")
}
