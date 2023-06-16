package domain

import (
	"context"
	"mongo-oplog-populator/config"
	"mongo-oplog-populator/internal/app/populator/service"

	"go.mongodb.org/mongo-driver/mongo"
)

type BulkInsert struct {
	NumberOfOperations int
}

func NewBulkInsert(numberOfOperations int) Populator {
	return &BulkInsert{NumberOfOperations: numberOfOperations}
}

func (bi *BulkInsert) PopulateData(ctx context.Context, client *mongo.Client, cfg config.Config, dataList []service.Data, opSize *service.OperationSize) {
	service.Populate(ctx, client, bi.NumberOfOperations, cfg, dataList, opSize)
}
