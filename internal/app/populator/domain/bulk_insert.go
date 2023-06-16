package domain

import (
	"context"
	"mongo-oplog-populator/internal/app/populator/generator"
	"mongo-oplog-populator/internal/app/populator/service"

	"go.mongodb.org/mongo-driver/mongo"
)

type BulkInsert struct {
}

func NewBulkInsert(numberOfOperations int) Populator {
	return &BulkInsert{}
}

func (bi *BulkInsert) PopulateData(ctx context.Context, client *mongo.Client, dataList []generator.Data, opSize *generator.OperationSize) {
	service.Populate(ctx, client, dataList, opSize)
}
