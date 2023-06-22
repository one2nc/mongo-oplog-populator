package service

import (
	"context"
	"fmt"
	"math"
	"sync"

	"mongo-oplog-populator/internal/app/populator/generator"

	"go.mongodb.org/mongo-driver/mongo"
)

type Populator interface {
	PopulateData(ctx context.Context, fakeData generator.FakeData)
}

type populator struct {
	OpSize         generator.OperationSize
	NoOfOperations int
	ModeFlag       bool
	Client         *mongo.Client
}

func NewPopulator(client *mongo.Client, modeFlag bool, noOfOperations int) Populator {
	return &populator{
		OpSize:         calculateOperationSize(noOfOperations),
		Client:         client,
		ModeFlag:       modeFlag,
		NoOfOperations: noOfOperations,
	}
}

func (p populator) PopulateData(ctx context.Context, fakeData generator.FakeData) {
	// Generate data
	dataGenerator := createDataGenerator(p.NoOfOperations, p.ModeFlag)
	dataChan := dataGenerator.Generate(ctx, fakeData)

	var wg sync.WaitGroup

	workerCnt := 25
	for i := 0; i < workerCnt; i++ {
		wg.Add(1)
		go func(workerID int) {
			defer wg.Done()
			//fmt.Printf("workerID: %v\n", workerID)
			p.worker(ctx, dataChan)
		}(i)
	}

	wg.Wait()
}

func (p populator) worker(ctx context.Context, dataChan <-chan generator.Data) {
	updateCount := 0
	deleteCount := 0
	idx := 0
	for data := range dataChan {
		select {
		case <-ctx.Done():
			// The context is done, stop reading Oplogs
			return
		default:
			// Context is still active, continue reading Oplogs
		}

		_, err := p.insertData(ctx, data)
		if err != nil {
			fmt.Printf("Failed to insert data at index %d: %s\n", idx, err.Error())
			return
		}

		//update
		if isMultipleOfSevenEightOrEleven(idx) {
			updateCount, err = p.updateData(ctx, data, updateCount)
			if err != nil {
				fmt.Printf("Failed to update data at index %d: %s\n", idx, err.Error())
				return
			}
		}

		//delete
		if isMultipleOfTwoNineortweleve(idx) {
			deleteCount, err = p.deleteData(ctx, data, deleteCount)
			if err != nil {
				fmt.Printf("Failed to delete data at index %d: %s\n", idx, err.Error())
				return
			}
		}

		idx++

		// reset update and delete count on each set of operations
		if idx%p.NoOfOperations == 0 {
			updateCount, deleteCount = 0, 0
		}
	}
}

func createDataGenerator(noOfOperations int, modeFlag bool) generator.Generator {
	var dataGenerator generator.Generator
	if modeFlag {
		dataGenerator = generator.NewStreamDataGenerator(noOfOperations)
	} else {
		dataGenerator = generator.NewBulkDataGenerator(noOfOperations)
	}
	return dataGenerator
}

func calculateOperationSize(totalOperation int) generator.OperationSize {
	i := (85 * totalOperation) / 100
	i = int(math.Max(float64(i), 1))

	u := (10 * totalOperation) / 100
	u = int(math.Max(float64(u), 1))

	d := (5 * totalOperation) / 100
	d = int(math.Max(float64(d), 1))

	opSize := generator.OperationSize{
		Insert: i,
		Update: u,
		Delete: d,
	}
	fmt.Printf("Out of total %d operations, %d: insert  %d:update %d:delete", totalOperation, i, u, d)

	return opSize
}

func (p populator) insertData(ctx context.Context, data generator.Data) (*mongo.InsertOneResult, error) {
	select {
	case <-ctx.Done():
		// The context is done, stop reading Oplogs
		return nil, nil
	default:
	}

	collection := data.GetCollection(p.Client)
	InsertOneResult, err := collection.InsertOne(ctx, data)
	if err != nil {
		return nil, err
	}
	// fmt.Println("Inserted data")
	return InsertOneResult, nil
}

func (p populator) updateData(ctx context.Context, data generator.Data, updateCount int) (int, error) {
	select {
	case <-ctx.Done():
		// The context is done, stop reading Oplogs
		return updateCount, nil
	default:
	}

	if updateCount < p.OpSize.Update {
		collection := data.GetCollection(p.Client)
		update := data.GetUpdate()
		_, err := collection.UpdateOne(ctx, data, update)
		if err != nil {
			return updateCount, err
		}
		updateCount++
		// fmt.Println("Updated data")
	}

	return updateCount, nil
}

func (p populator) deleteData(ctx context.Context, data generator.Data, deleteCount int) (int, error) {
	select {
	case <-ctx.Done():
		// The context is done, stop reading Oplogs
		return deleteCount, nil
	default:
	}

	if deleteCount < p.OpSize.Delete {
		collection := data.GetCollection(p.Client)
		_, err := collection.DeleteOne(ctx, data)
		if err != nil {
			return deleteCount, err
		}
		deleteCount++
		// fmt.Println("Deleted data")
	}
	return deleteCount, nil
}

func isMultipleOfSevenEightOrEleven(n int) bool {
	if n == 0 {
		return false
	}
	return n%7 == 0 || n%8 == 0 || n%11 == 0 || n == 10
}

func isMultipleOfTwoNineortweleve(n int) bool {
	if n == 0 {
		return false
	}
	return n%2 == 0 || n%9 == 0 || n%12 == 0 || n == 10
}
