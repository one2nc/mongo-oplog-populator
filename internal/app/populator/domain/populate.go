package domain

import "go.mongodb.org/mongo-driver/mongo"

type Populator interface {
	PopulateData(client *mongo.Client)
}
