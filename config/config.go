package config

import (
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func GetClient() *mongo.Client {
	clientOptions := options.Client().ApplyURI("mongodb://127.0.0.1:30001").SetDirect(true)
	client, err := mongo.Connect(context.Background(), clientOptions)
	if err != nil {
		fmt.Println("Error connecting to MongoDB:", err)
		return nil
	}
	return client
}

func DisconnectClient(client *mongo.Client) {
	client.Disconnect(context.Background())
}
