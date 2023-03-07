package config

import (
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func GetMongoConnection(args []string) *mongo.Collection {
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")
	client, err := mongo.Connect(context.Background(), clientOptions)
	if err != nil {
		fmt.Println("Error connecting to MongoDB:", err)
		return nil
	}

	defer client.Disconnect(context.Background())
	oplogCollection := client.Database(args[1]).Collection("oplog.rs")
	return oplogCollection
}
