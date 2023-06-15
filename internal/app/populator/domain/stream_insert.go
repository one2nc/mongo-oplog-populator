package domain

import (
	"context"
	"fmt"
	"mongo-oplog-populator/config"
	"mongo-oplog-populator/internal/app/populator/service"
	"os"
	"os/signal"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
)

type StreamInsert struct {
	NumberOfOperations int
}

func NewStreamInsert(numberOfOperations int) Populator {
	return &StreamInsert{NumberOfOperations: numberOfOperations}
}

func (si *StreamInsert) PopulateData(client *mongo.Client, cfg config.Config, ctx context.Context) {
	ticker := time.NewTicker(time.Second * 1)
	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt)

	a := 1
	for {
		select {
		case <-ticker.C:
			println("Second : ", a)
			go service.Populate(client, si.NumberOfOperations, cfg, ctx)
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