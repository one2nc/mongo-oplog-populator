package service

import (
	"mongo-oplog-populator/internal/app/populator/types"

	"go.mongodb.org/mongo-driver/mongo"
)

type Data interface {
	GetCollection() *mongo.Collection
	GetData(attributes types.PersonnelInfo, index int) Data
	GetUpdateSet() interface{}
	GetUpdateUnset() interface{}
	GetUpdate() interface{}
}
