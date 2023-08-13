package service

import (
	"context"
	"fmt"
	"math"
	"math/rand"
	"sync"
	"time"

	"mongo-oplog-populator/config"
	"mongo-oplog-populator/internal/app/populator/generator"

	"go.mongodb.org/mongo-driver/mongo"
)

type Populator interface {
	PopulateData(ctx context.Context, fakeData generator.FakeData)
}

type populator struct {
	cfg config.Config

	opSize         generator.OperationSize
	noOfOperations int
	ModeFlag       bool
	Client         *mongo.Client

	sharedIdx     int
	updateIndices map[int]struct{}
	deleteIndices map[int]struct{}
	mu            sync.Mutex
}

func NewPopulator(cfg config.Config, client *mongo.Client, modeFlag bool, noOfOperations int) Populator {
	opSize := calculateOperationSize(noOfOperations)

	// Seed the random number generator with the current time
	rand.Seed(time.Now().UnixNano())

	// Create a set to keep track of update indices
	updateIndices := make(map[int]struct{})

	// Generate a random number between 0 and totalNumbers (exclusive)
	for len(updateIndices) < opSize.Update {
		randomNumber := rand.Intn(noOfOperations)

		// Add the random index to the updateIndices set
		updateIndices[randomNumber] = struct{}{}
	}

	// Create a set to keep track of update indices
	deleteIndices := make(map[int]struct{})

	// Generate a random number between 0 and totalNumbers (exclusive)
	for len(deleteIndices) < opSize.Delete {
		randomNumber := rand.Intn(noOfOperations)

		// Add the random index to the updateIndices set
		deleteIndices[randomNumber] = struct{}{}
	}

	return &populator{
		cfg:            cfg,
		opSize:         opSize,
		updateIndices:  updateIndices,
		deleteIndices:  deleteIndices,
		Client:         client,
		ModeFlag:       modeFlag,
		noOfOperations: noOfOperations,
	}
}

func (p *populator) PopulateData(ctx context.Context, fakeData generator.FakeData) {

	dataGenerator := createDataGenerator(p.noOfOperations, p.ModeFlag)
	dataChan := dataGenerator.GenerateDocument(ctx, fakeData)

	var wg sync.WaitGroup
	workerCnt := 25
	for idx := 0; idx < workerCnt; idx++ {
		wg.Add(1)
		go func() {
			p.worker(ctx, dataChan, &wg)
		}()
	}

	wg.Wait()
}

func (p *populator) worker(ctx context.Context, dataChan <-chan generator.Document, wgOuter *sync.WaitGroup) {
	defer wgOuter.Done()

	var wg sync.WaitGroup
	for doc := range dataChan {

		p.mu.Lock()
		idx := p.sharedIdx
		p.sharedIdx++
		p.mu.Unlock()

		wg.Add(1)
		go func(idx int, doc generator.Document) {
			defer wg.Done()

			for dbIdx := 1; dbIdx <= p.cfg.MaxDatabases; dbIdx++ {
				wg.Add(1)
				go func(idx, dbIdx int, doc generator.Document) {
					defer wg.Done()

					dbName := fmt.Sprintf("database%d", dbIdx)
					db := p.Client.Database(dbName)

					for collIdx := 1; collIdx <= p.cfg.MaxCollections; collIdx++ {
						wg.Add(1)
						go func(idx, collIdx int, db *mongo.Database) {
							defer wg.Done()

							collectionName := fmt.Sprintf("collection%d", collIdx)
							collection := db.Collection(collectionName)

							_, err := collection.InsertOne(context.Background(), doc)
							if err != nil {
								fmt.Printf("Failed to insert document: %v\n", err)
							}

							mapIdx := idx % p.noOfOperations
							//update
							if _, ok := p.updateIndices[mapIdx]; ok {
								update := doc.GetUpdate()
								_, err := collection.UpdateOne(ctx, doc, update)
								if err != nil {
									fmt.Printf("Failed to update document at index %d: %v\n", idx, err)
								}
							}

							//delete
							if _, ok := p.deleteIndices[mapIdx]; ok {
								_, err := collection.DeleteOne(ctx, doc)
								if err != nil {
									fmt.Printf("Failed to delete document at index %d: %v\n", idx, err)
								}
							}

						}(idx, collIdx, db)
					}
				}(idx, dbIdx, doc)
			}
		}(idx, doc)
	}
	wg.Wait()
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
	fmt.Printf("Out of total %d operations, %d: insert  %d:update %d:delete\n", totalOperation, i, u, d)

	return opSize
}
