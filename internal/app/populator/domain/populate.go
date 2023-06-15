package domain

import (
	"context"
	"mongo-oplog-populator/config"

	"go.mongodb.org/mongo-driver/mongo"
)

type Populator interface {
	PopulateData(client *mongo.Client, cfg config.Config, ctx context.Context)
}
