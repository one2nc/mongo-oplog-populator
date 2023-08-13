package generator

import (
	"context"
)

type BulkInsert struct {
	noOfOperations int
}

func NewBulkDataGenerator(noOfOperations int) Generator {
	return &BulkInsert{
		noOfOperations: noOfOperations,
	}
}

func (bi BulkInsert) GenerateDocument(ctx context.Context, fakeData FakeData) <-chan Document {
	data := make(chan Document)
	go func() {
		tempData := generateDocument(bi.noOfOperations, fakeData)
		for _, td := range tempData {
			data <- td
		}
		close(data)
	}()
	return data
}
