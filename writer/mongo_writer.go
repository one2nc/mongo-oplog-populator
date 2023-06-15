package writer

import (
	"context"
	"fmt"
	"mongo-oplog-populator/config"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func GetClient(cfg config.Config,ctx context.Context) *mongo.Client {
	clientOptions := options.Client().ApplyURI(cfg.MongoUri).SetDirect(true)
	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		fmt.Println("Error connecting to MongoDB:", err)
		return nil
	}
	return client
}

func DisconnectClient(client *mongo.Client,ctx context.Context) {
	client.Disconnect(ctx)
}
