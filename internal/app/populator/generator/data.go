package generator

import (
	"go.mongodb.org/mongo-driver/mongo"
)

type Data interface {
	GetCollection(client *mongo.Client) *mongo.Collection
	GetData(attributes PersonnelInfo, index int) Data
	GetUpdateSet() interface{}
	GetUpdateUnset() interface{}
	GetUpdate() interface{}
}
