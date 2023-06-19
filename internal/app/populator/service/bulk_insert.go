package service

import (
	"context"
	"mongo-oplog-populator/internal/app/populator/generator"
)

type BulkInsert struct {
	NoOfOperations int
}

func NewBulkInsert(noOfOperations int) Populator {
	return &BulkInsert{
		NoOfOperations: noOfOperations,
	}
}

func (bi BulkInsert) PopulateData(ctx context.Context, fakeData generator.FakeData) {
	data := GenerateData(bi.NoOfOperations, fakeData)

	Populate(ctx, Client, data, OpSize)
}
