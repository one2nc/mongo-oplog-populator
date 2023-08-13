package generator

import (
	"context"
	"fmt"
	"sync"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
)

type StreamInsert struct {
	noOfOperations int
	client         *mongo.Client
	opSize         *OperationSize
}

func NewStreamDataGenerator(noOfOperations int) Generator {
	return &StreamInsert{
		noOfOperations: noOfOperations,
	}
}

func (si *StreamInsert) GenerateDocument(ctx context.Context, fakeData FakeData) <-chan Document {
	data := make(chan Document)
	var wg sync.WaitGroup
	go func() {
		tempData := generateDocument(si.noOfOperations, fakeData)
		ticker := time.NewTicker(time.Second * 1)
		seconds := 1
		for {
			select {
			case <-ticker.C:
				fmt.Printf("seconds: %v, %d\n", seconds, len(tempData))
				seconds++
				wg.Add(1)
				go func() {
					defer wg.Done()

					for _, td := range tempData {
						data <- td
					}
				}()
			case <-ctx.Done():
				ticker.Stop()

				wg.Wait()

				close(data)
				return
			}
		}
	}()
	return data
}
