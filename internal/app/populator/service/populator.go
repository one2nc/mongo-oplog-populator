package service

import (
	"context"
	"fmt"
	"math"
	"math/rand"
	"mongo-oplog-populator/config"
	"mongo-oplog-populator/internal/app/populator/types"

	"time"

	"go.mongodb.org/mongo-driver/mongo"
)

// remove client
var client *mongo.Client

var ctx = context.Background()

//TODO : generateData should be called once in main and passed here
func Populate(ctx context.Context, mclient *mongo.Client, operations int, cfg config.Config, personnelInfo types.PersonnelInfo) []interface{} {
	client = mclient

	//TODO : calculate once  and pass to populate func
	opSize := calculateOperationSize(operations)

	//TODO-DONE: move reader from here

	dataList := generateData(opSize.insert, personnelInfo)

	var updateCount = 0
	var deleteCount = 0
	var insertedDataList []Data

	var results []interface{}
	rand.Seed(time.Now().UnixNano())

	//TODO : refactor this part of code
populateLoop:
	for i := 0; i < len(dataList); i++ {
		select {
		case <-ctx.Done():
			// The context is done, stop reading Oplogs
			break populateLoop
		default:
			// Context is still active, continue reading Oplogs
		}
		insertedData, err := insertData(dataList[i])
		if err != nil {
			fmt.Printf("Failed to insert data at index %d: %s\n", i, err.Error())
			continue
		}
		insertedDataList = append(insertedDataList, dataList[i])
		println("inserting Data")
		//update
		if isMultipleOfSevenEightOrEleven(i) {
			if updateCount < opSize.update {
				updateResult, err := updateData(insertedDataList[i])
				if err != nil {
					fmt.Printf("Failed to update data at index %d: %s\n", i, err.Error())
					continue
				}
				updateCount++
				println("updating data")
				results = append(results, updateResult)
			}
		}

		//delete
		if isMultipleOfTwoNineortweleve(i) {
			if deleteCount < opSize.delete {
				indx := rand.Intn(i)
				deleteResult, err := deleteData(insertedDataList[indx])
				if err != nil {
					fmt.Printf("Failed to delete data at index %d: %s\n", i, err.Error())
					continue
				}
				insertedDataList = append(insertedDataList[:indx], insertedDataList[indx:]...)
				deleteCount++
				results = append(results, deleteResult)
				println("Deleting Data")
			}
		}
		results = append(results, insertedData)
	}

	//insert data for alter table
	data := generateDataAlterTable(3, personnelInfo)
alterLoop:
	for i := 0; i < len(data); i++ {
		select {
		case <-ctx.Done():
			// The context is done, stop reading Oplogs
			break alterLoop
		default:
			// Context is still active, continue reading Oplogs
		}
		fmt.Println("Alter successfull")
		insertedData, err := insertData(data[i])
		if err != nil {
			fmt.Printf("Failed to insert data at index %d: %s\n", i, err.Error())
			continue
		}
		results = append(results, insertedData)
	}
	return results
}

func calculateOperationSize(totalOperation int) *OperationSize {
	i := (85 * totalOperation) / 100
	i = int(math.Max(float64(i), 1))

	u := (10 * totalOperation) / 100
	u = int(math.Max(float64(u), 1))

	d := (5 * totalOperation) / 100
	d = int(math.Max(float64(d), 1))

	opSize := &OperationSize{
		insert: i,
		update: u,
		delete: d,
	}

	return opSize
}

func generateData(operations int, attributes types.PersonnelInfo) []Data {
	x := operations / 2
	var data []Data
	index := 0
	for i := 0; i < x; i++ {
		emp := &Employee{}
		empData := emp.GetData(attributes, index)
		data = append(data, empData)
		student := &Student{}
		studentData := student.GetData(attributes, index)
		data = append(data, studentData)
		index++
		//to reset if attributes size < input number of operations size. Will continue to read data in a cycle
		if index > len(attributes.FirstNames)-2 {
			index = 0
		}
	}
	shuffle(data)
	return data
}

func generateDataAlterTable(operations int, attributes types.PersonnelInfo) []Data {
	var data []Data
	index := 0
	for i := 0; i < operations; i++ {
		emp := &EmployeeA{}
		empData := emp.GetData(attributes, index)
		data = append(data, empData)
		student := &StudentA{}
		studentData := student.GetData(attributes, index)
		data = append(data, studentData)
		index++
		//to reset if attributes size < input number of operations size. Will continue to read data in a cycle
		if index > len(attributes.FirstNames)-2 {
			index = 0
		}
	}
	shuffle(data)
	return data
}

func shuffle(slice []Data) {
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

func insertData(data Data) (*mongo.InsertOneResult, error) {
	collection := data.GetCollection()
	InsertOneResult, err := collection.InsertOne(ctx, data)
	if err != nil {
		return nil, err
	}
	return InsertOneResult, nil
}

func updateData(data Data) (*mongo.UpdateResult, error) {
	collection := data.GetCollection()
	update := data.GetUpdate()
	updateOneResult, err := collection.UpdateOne(ctx, data, update)
	if err != nil {
		return nil, err
	}
	return updateOneResult, nil
}

func deleteData(data Data) (*mongo.DeleteResult, error) {
	collection := data.GetCollection()
	deleteResult, err := collection.DeleteOne(ctx, data)
	if err != nil {
		return nil, err
	}
	return deleteResult, nil
}
