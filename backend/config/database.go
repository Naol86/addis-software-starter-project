package config

import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

func NewMongoDBConfig(env *Env) *mongo.Client {
	client, err := mongo.Connect(options.Client().ApplyURI(env.DBUri))
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Connected to MongoDB")
	return client
}

func CloseMongoDBConnection(client *mongo.Client) {
	ctx := context.TODO()
	err := client.Disconnect(ctx)
	if err != nil {
		panic(err)
	}
	log.Println("Connection to MongoDB closed.")
}