package domain

import (
	"mongo-oplog-populator/internal/app/populator/service"

	"go.mongodb.org/mongo-driver/mongo"
)

type BulkInsert struct {
	NumberOfOperations int
}

func NewBulkInsert(numberOfOperations int) Populator {
	return &BulkInsert{NumberOfOperations: numberOfOperations}
}

func (bi *BulkInsert) PopulateData(client *mongo.Client) {
	service.Populate(client, bi.NumberOfOperations)
}
