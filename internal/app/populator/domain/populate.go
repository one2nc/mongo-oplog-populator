package domain

import (
	"context"
	"mongo-oplog-populator/config"
	"mongo-oplog-populator/internal/app/populator/service"

	"go.mongodb.org/mongo-driver/mongo"
)

type Populator interface {
	PopulateData(ctx context.Context, client *mongo.Client, cfg config.Config, dataList []service.Data, opSize *service.OperationSize)
}
