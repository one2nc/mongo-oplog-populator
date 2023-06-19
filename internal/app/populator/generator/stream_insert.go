package generator

import (
	"context"
	"fmt"
	"sync"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
)

type StreamInsert struct {
	NoOfOperations int
	Client         *mongo.Client
	OpSize         *OperationSize
}

func NewStreamDataGenerator(noOfOperations int) Generator {
	return &StreamInsert{
		NoOfOperations: noOfOperations,
	}
}

func (si *StreamInsert) Generate(ctx context.Context, fakeData FakeData) <-chan Data {
	data := make(chan Data)
	var wg sync.WaitGroup
	go func() {
		tempData := generateData(si.NoOfOperations, fakeData)
		ticker := time.NewTicker(time.Second * 1)
		seconds := 1
		for {
			select {
			case <-ticker.C:
				fmt.Printf("seconds: %v\n", seconds)
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
