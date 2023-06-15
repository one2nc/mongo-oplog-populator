package domain

import (
	"context"
	"mongo-oplog-populator/config"
	"mongo-oplog-populator/internal/app/populator/service"
	"mongo-oplog-populator/internal/app/populator/types"

	"go.mongodb.org/mongo-driver/mongo"
)

type BulkInsert struct {
	NumberOfOperations int
}

func NewBulkInsert(numberOfOperations int) Populator {
	return &BulkInsert{NumberOfOperations: numberOfOperations}
}

func (bi *BulkInsert) PopulateData(ctx context.Context, client *mongo.Client, cfg config.Config, personnelInfo types.PersonnelInfo) {
	service.Populate(ctx, client, bi.NumberOfOperations, cfg, personnelInfo)
}
