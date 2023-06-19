package service

import (
	"context"
	"mongo-oplog-populator/internal/app/populator/generator"
)

type Populator interface {
	PopulateData(ctx context.Context, fakeData generator.FakeData)
}
