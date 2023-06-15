package domain

import (
	"context"
	"mongo-oplog-populator/config"
	"mongo-oplog-populator/internal/app/populator/types"

	"go.mongodb.org/mongo-driver/mongo"
)

type Populator interface {
	PopulateData(ctx context.Context, client *mongo.Client, cfg config.Config, personnelInfo types.PersonnelInfo)
}
