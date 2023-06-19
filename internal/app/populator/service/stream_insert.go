package service

import (
	"context"
	"fmt"
	"mongo-oplog-populator/internal/app/populator/generator"
	"os"
	"os/signal"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
)

type StreamInsert struct {
	NoOfOperations int
	Client         *mongo.Client
	OpSize         *generator.OperationSize
}

func NewStreamInsert(noOfOperations int) Populator {
	return &StreamInsert{
		NoOfOperations: noOfOperations,
	}
}

func (si *StreamInsert) PopulateData(ctx context.Context, fakeData generator.FakeData) {
	data := GenerateData(si.NoOfOperations, fakeData)
	ticker := time.NewTicker(time.Second * 1)
	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt)

	a := 1
	for {
		select {
		case <-ticker.C:
			println("Second : ", a)
			go Populate(ctx, si.Client, data, si.OpSize)
			a++
		case <-interrupt:
			fmt.Println("Interrupt signal received, stopping program...")
			ticker.Stop()
			return
		case <-ctx.Done():
			return
		}
	}
}
