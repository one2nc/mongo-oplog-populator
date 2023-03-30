package populator

import (
	"context"
	"fmt"
	"math"
	"math/rand"
	"os"
	"os/signal"
	"time"

	"github.com/brianvoe/gofakeit"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

var client *mongo.Client
var ctx = context.Background()
var subjects = []string{"Maths", "Science", "Social Studies", "English"}
var positions = []string{"Manager", "Engineer", "Salesman", "Developer"}

func BatchInsert(mclient *mongo.Client, batchInsert int) {

	ticker := time.NewTicker(time.Second * 1)
	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt)

	a := 1
	for {
		select {
		case <-ticker.C:
			println("Second : ", a)
			go Populate(mclient, batchInsert)
			a++
		case <-interrupt:
			fmt.Println("Interrupt signal received, stopping program...")
			ticker.Stop()
			return
		}
	}
}

func Populate(mclient *mongo.Client, operations int) []interface{} {

	client = mclient
	opSize := calculateOperationSize(operations)
	dataList := generateData(opSize.insert)
	var updateCount = 0
	var deleteCount = 0
	var insertedDataList []Data
	//var updatedDataList []Data
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
	data := generateDataAlterTable(3)
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
func generateData(operations int) []Data {
	x := operations / 2
	var data []Data
	for i := 0; i < x; i++ {
		data = append(data, &Employee{})
		data = append(data, &Student{})
	}
	shuffle(data)
	return data
}

func generateDataAlterTable(operations int) []Data {
	var data []Data
	for i := 0; i < operations; i++ {
		data = append(data, &Employee{})
		data = append(data, &Student{})
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
	//extract
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

type Data interface {
	GetCollection() *mongo.Collection
	GetData() interface{}
	GetUpdateSet() interface{}
	GetUpdateUnset() interface{}
	GetUpdate() interface{}
}

func (s *Student) GetCollection() *mongo.Collection {
	return client.Database("student").Collection("students")
}

func (s *StudentU) GetCollection() *mongo.Collection {
	return client.Database("student").Collection("students")
}

func (s *Student) GetData() interface{} {
	return &Student{
		Id:      gofakeit.UUID(),
		Name:    gofakeit.FirstName() + " " + gofakeit.LastName(),
		Age:     rand.Intn(10) + 18,
		Subject: subjects[rand.Intn(len(subjects))],
	}
}

func (s *StudentU) GetData() interface{} {
	return &StudentU{
		Id:           gofakeit.UUID(),
		Name:         gofakeit.FirstName() + " " + gofakeit.LastName(),
		Age:          rand.Intn(10) + 18,
		Subject:      subjects[rand.Intn(len(subjects))],
		Is_Graduated: gofakeit.Bool(),
	}
}
func (s *StudentU) GetUpdateSet() interface{} {
	return bson.M{
		"$set": bson.M{
			"Age": rand.Intn(10) + 18,
		},
	}
}

func (s *Student) GetUpdateSet() interface{} {
	return bson.M{
		"$set": bson.M{
			"Age": rand.Intn(10) + 18,
		},
	}
}

func (s *StudentU) GetUpdateUnset() interface{} {
	return bson.M{
		"$unset": bson.M{
			"Subject": "",
		},
	}
}

func (s *Student) GetUpdateUnset() interface{} {
	return bson.M{
		"$unset": bson.M{
			"Subject": "",
		},
	}
}

func (s *StudentU) GetUpdate() interface{} {
	updateS := gofakeit.Bool()
	if updateS {
		return s.GetUpdateSet()
	} else {
		return s.GetUpdateUnset()
	}
}

func (s *Student) GetUpdate() interface{} {
	updateS := gofakeit.Bool()
	if updateS {
		return s.GetUpdateSet()
	} else {
		return s.GetUpdateUnset()
	}
}

func (e *Employee) GetCollection() *mongo.Collection {
	return client.Database("Employee").Collection("employees")
}

func (e *EmployeeU) GetCollection() *mongo.Collection {
	return client.Database("Employee").Collection("employees")
}

func (e *Employee) GetData() interface{} {
	return &Employee{
		Id:       gofakeit.UUID(),
		Name:     gofakeit.FirstName() + " " + gofakeit.LastName(),
		Age:      rand.Intn(30) + 20,
		Salary:   rand.Float64() * 10000,
		Phone:    []Phone{{gofakeit.Phone(), gofakeit.Phone()}},
		Position: positions[rand.Intn(len(positions))],
	}
}

func (e *EmployeeU) GetData() interface{} {
	return &EmployeeU{
		Id:        gofakeit.UUID(),
		Name:      gofakeit.FirstName() + " " + gofakeit.LastName(),
		Age:       rand.Intn(30) + 20,
		Salary:    rand.Float64() * 10000,
		Position:  positions[rand.Intn(len(positions))],
		WorkHours: rand.Intn(8) + 4,
	}
}

func (e *EmployeeU) GetUpdateSet() interface{} {
	return bson.M{
		"$set": bson.M{
			"Age": rand.Intn(10) + 18,
		},
	}
}

func (e *Employee) GetUpdateSet() interface{} {
	return bson.M{
		"$set": bson.M{
			"Age": rand.Intn(10) + 18,
		},
	}
}
func (e *EmployeeU) GetUpdateUnset() interface{} {
	return bson.M{
		"$unset": bson.M{
			"Position": "",
		},
	}
}

func (e *Employee) GetUpdateUnset() interface{} {
	return bson.M{
		"$unset": bson.M{
			"Position": "",
		},
	}
}

func (e *EmployeeU) GetUpdate() interface{} {
	updateE := gofakeit.Bool()
	if updateE {
		return e.GetUpdateSet()
	} else {
		return e.GetUpdateUnset()
	}
}

func (e *Employee) GetUpdate() interface{} {
	updateE := gofakeit.Bool()
	if updateE {
		return e.GetUpdateSet()
	} else {
		return e.GetUpdateUnset()
	}
}
