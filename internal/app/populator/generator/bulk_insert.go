package generator

import (
	"context"
)

type BulkInsert struct {
	NoOfOperations int
}

func NewBulkDataGenerator(noOfOperations int) Generator {
	return &BulkInsert{
		NoOfOperations: noOfOperations,
	}
}

func (bi BulkInsert) Generate(ctx context.Context, fakeData FakeData) <-chan Data {
	data := make(chan Data)
	go func() {
		tempData := generateData(bi.NoOfOperations, fakeData)
		for _, td := range tempData {
			data <- td
		}
		close(data)
	}()
	return data
}
