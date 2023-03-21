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

// func GetMongoConnection(args int, client *mongo.Client) *mongo.Collection {
// 	oplogCollection := client.Database("test").Collection("test")
// 	stud := StudentInfo{Id: gofakeit.UUID(), Name: gofakeit.Name(), Roll_no: gofakeit.Number(0, 50), Is_Graduated: gofakeit.Bool(), Date_Of_Birth: gofakeit.Date().Format("02-01-2006")}
// 	_, err := oplogCollection.InsertOne(context.Background(), stud)
// 	if err != nil {
// 		panic(err)
// 	}
// 	return oplogCollection
// }

func DisconnectClient(client *mongo.Client) {
	client.Disconnect(context.Background())
}
