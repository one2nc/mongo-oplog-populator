package service

import (
	"context"
	"fmt"
	"math"
	"math/rand"
	"mongo-oplog-populator/config"
	"mongo-oplog-populator/internal/app/populator/reader"
	"mongo-oplog-populator/internal/app/populator/types"

	"os"
	"time"

	"github.com/brianvoe/gofakeit"
	"go.mongodb.org/mongo-driver/mongo"
)

// remove client
var client *mongo.Client

var ctx = context.Background()

var attributes types.Attributes

func Populate(mclient *mongo.Client, operations int, cfg config.Config) []interface{} {
	client = mclient

	opSize := calculateOperationSize(operations)

	// if csv file does not exist, generate some random/fake data, and populate it to the CSV file
	_, err := os.Stat(cfg.CsvFileName)
	if os.IsNotExist(err) {
		attributes = generateDataOnce(operations)
		CreateCSVFile(cfg.CsvFileName, operations, attributes)
	}

	//read from csv
	csvReader := reader.NewCSVReader(cfg.CsvFileName)
	var attributes types.Attributes
	attributes = csvReader.ReadData()

	dataList := generateData(opSize.insert, attributes)

	var updateCount = 0
	var deleteCount = 0
	var insertedDataList []Data

	var results []interface{}
	rand.Seed(time.Now().UnixNano())

	for i := 0; i < len(dataList); i++ {
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
				results = append(results, updateResult)
				println("updating Data")
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
	data := generateDataAlterTable(3, attributes)
	for i := 0; i < len(data); i++ {
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

func generateDataOnce(n int) types.Attributes {
	var subjects = []string{"Maths", "Science", "Social Studies", "English"}
	var positions = []string{"Manager", "Engineer", "Salesman", "Developer"}

	if n > 50 {
		n = n / 4
	}
	for i := 0; i < n; i++ {
		attributes.FirstNames = append(attributes.FirstNames, gofakeit.FirstName())
		attributes.LastNames = append(attributes.LastNames, gofakeit.LastName())
		attributes.Subjects = append(attributes.Subjects, subjects[rand.Intn(len(subjects))])
		attributes.StreetAddresses = append(attributes.StreetAddresses, gofakeit.Address().Street)
		attributes.Positions = append(attributes.Positions, positions[rand.Intn(len(positions))])
		attributes.Zips = append(attributes.Zips, gofakeit.Zip())
		attributes.PhoneNumbers = append(attributes.PhoneNumbers, gofakeit.Phone())
		attributes.Ages = append(attributes.Ages, rand.Intn(30)+20)
		attributes.Workhours = append(attributes.Workhours, rand.Intn(8)+4)
		attributes.Salaries = append(attributes.Salaries, rand.Float64()*10000)
	}

	return attributes
}

func generateData(operations int, attributes types.Attributes) []Data {
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

func generateDataAlterTable(operations int, attributes types.Attributes) []Data {
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
