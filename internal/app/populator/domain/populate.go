package domain

import (
	"context"
	"mongo-oplog-populator/internal/app/populator/generator"

	"go.mongodb.org/mongo-driver/mongo"
)

type Populator interface {
	PopulateData(ctx context.Context, client *mongo.Client, dataList []generator.Data, opSize *generator.OperationSize)
}
