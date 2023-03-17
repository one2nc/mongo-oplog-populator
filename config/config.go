package config

import (
	"context"
	"fmt"

	"github.com/brianvoe/gofakeit"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type StudentInfo struct {
	Id            string `json:"Id"`
	Name          string `json:"Name,omitempty"`
	Roll_no       int    `json:"Roll_no,omitempty"`
	Is_Graduated  bool   `json:"Is_Graduated,omitempty"`
	Date_Of_Birth string `json:"Date_Of_Birth,omitempty"`
	// Address       []Address `json:"Address,omitempty"`
	// Phone         []Phone   `json:"Phone,omitempty"`
	Email string `json:"Email,omitempty"`
}

func GetClient() *mongo.Client {
	clientOptions := options.Client().ApplyURI("mongodb://127.0.0.1:2717").SetDirect(true)
	client, err := mongo.Connect(context.Background(), clientOptions)
	if err != nil {
		fmt.Println("Error connecting to MongoDB:", err)
		return nil
	}
	return client
}
func GetMongoConnection(args []string, client *mongo.Client) *mongo.Collection {
	oplogCollection := client.Database(args[1]).Collection(args[2])
	stud := StudentInfo{Id: gofakeit.UUID(), Name: gofakeit.Name(), Roll_no: gofakeit.Number(0, 50), Is_Graduated: gofakeit.Bool(), Date_Of_Birth: gofakeit.Date().Format("02-01-2006")}
	_, err := oplogCollection.InsertOne(context.Background(), stud)
	if err != nil {
		panic(err)
	}
	return oplogCollection
}

func DisconnectClient(client *mongo.Client) {
	client.Disconnect(context.Background())
}
