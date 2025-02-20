package database

import (
	"context"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var ctx = context.TODO()

func DBInstance() *mongo.Client {
	MongoDB := "mongodb://localhost:27017"
	fmt.Print(MongoDB)

	// crate a mongodb client with options
	clientOptions := options.Client().ApplyURI(MongoDB)
	client, err := mongo.Connect(ctx, clientOptions)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Connected to mongodb")
	return client
}

func OpenCollection(client *mongo.Client, collectionName string) *mongo.Collection {
	var Collection *mongo.Collection = client.Database("Restaurent").Collection(collectionName)

	return Collection
}
