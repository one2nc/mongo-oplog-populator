package service

import (
	"context"
	"fmt"
	"math"
	"math/rand"
	"mongo-oplog-populator/internal/app/populator/generator"

	"time"

	"go.mongodb.org/mongo-driver/mongo"
)

// TODO-DONE: remove client

var ctx = context.Background()

// TODO: generateData should be called once in main and passed here
func Populate(ctx context.Context, mclient *mongo.Client, dataList []generator.Data, opSize *generator.OperationSize) {

	//TODO: calculate opsize once  and pass to populate func
	//TODO-DONE: move reader from here

	var updateCount = 0
	var deleteCount = 0
	var insertedDataList []generator.Data

	rand.Seed(time.Now().UnixNano())

	//TODO: refactor this part of code
populateLoop:
	for i := 0; i < len(dataList); i++ {
		select {
		case <-ctx.Done():
			// The context is done, stop reading Oplogs
			break populateLoop
		default:
			// Context is still active, continue reading Oplogs
		}
		_, err := insertData(dataList[i], mclient)
		if err != nil {
			fmt.Printf("Failed to insert data at index %d: %s\n", i, err.Error())
			continue
		}
		insertedDataList = append(insertedDataList, dataList[i])
		println("inserting Data")
		//update
		if isMultipleOfSevenEightOrEleven(i) {
			if updateCount < opSize.Update {
				_, err := updateData(insertedDataList[i], mclient)
				if err != nil {
					fmt.Printf("Failed to update data at index %d: %s\n", i, err.Error())
					continue
				}
				updateCount++
				println("updating data")
			}
		}

		//delete
		if isMultipleOfTwoNineortweleve(i) {
			if deleteCount < opSize.Delete {
				indx := rand.Intn(i)
				_, err := deleteData(insertedDataList[indx], mclient)
				if err != nil {
					fmt.Printf("Failed to delete data at index %d: %s\n", i, err.Error())
					continue
				}
				insertedDataList = append(insertedDataList[:indx], insertedDataList[indx:]...)
				deleteCount++
				println("Deleting Data")
			}
		}
	}
}

func CalculateOperationSize(totalOperation int) *generator.OperationSize {
	i := (85 * totalOperation) / 100
	i = int(math.Max(float64(i), 1))

	u := (10 * totalOperation) / 100
	u = int(math.Max(float64(u), 1))

	d := (5 * totalOperation) / 100
	d = int(math.Max(float64(d), 1))

	opSize := &generator.OperationSize{
		Insert: i,
		Update: u,
		Delete: d,
	}

	return opSize
}

func GenerateData(operations int, attributes generator.PersonnelInfo) []generator.Data {
	x := operations / 2
	var data []generator.Data
	index := 0
	for i := 0; i < x; i++ {
		emp := &generator.Employee{}
		empData := emp.GetData(attributes, index)
		data = append(data, empData)
		student := &generator.Student{}
		studentData := student.GetData(attributes, index)
		data = append(data, studentData)
		index++
		//to reset if attributes size < input number of operations size. Will continue to read data in a cycle
		if index > len(attributes.FirstNames)-2 {
			index = 0
		}
	}
	dataAlterTable := generateDataAlterTable(3, attributes)
	data = append(data, dataAlterTable...)
	shuffle(data)
	return data
}

func generateDataAlterTable(operations int, attributes generator.PersonnelInfo) []generator.Data {
	var data []generator.Data
	index := 0
	for i := 0; i < operations; i++ {
		emp := &generator.EmployeeA{}
		empData := emp.GetData(attributes, index)
		data = append(data, empData)
		student := &generator.StudentA{}
		studentData := student.GetData(attributes, index)
		data = append(data, studentData)
		index++
		//to reset if attributes size < input number of operations size. Will continue to read data in a cycle
		if index > len(attributes.FirstNames)-2 {
			index = 0
		}
	}
	return data
}

func shuffle(slice []generator.Data) {
	for i := range slice {
		j := rand.Intn(i + 1)
		slice[i], slice[j] = slice[j], slice[i]
	}
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

func insertData(data generator.Data, client *mongo.Client) (*mongo.InsertOneResult, error) {
	collection := data.GetCollection(client)
	InsertOneResult, err := collection.InsertOne(ctx, data)
	if err != nil {
		return nil, err
	}
	return InsertOneResult, nil
}

func updateData(data generator.Data, client *mongo.Client) (*mongo.UpdateResult, error) {
	collection := data.GetCollection(client)
	update := data.GetUpdate()
	updateOneResult, err := collection.UpdateOne(ctx, data, update)
	if err != nil {
		return nil, err
	}
	return updateOneResult, nil
}

func deleteData(data generator.Data, client *mongo.Client) (*mongo.DeleteResult, error) {
	collection := data.GetCollection(client)
	deleteResult, err := collection.DeleteOne(ctx, data)
	if err != nil {
		return nil, err
	}
	return deleteResult, nil
}
