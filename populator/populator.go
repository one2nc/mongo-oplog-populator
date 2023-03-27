package populator

import (
	"context"
	"errors"
	"fmt"
	"math"
	"math/rand"
	"reflect"
	"time"

	"github.com/brianvoe/gofakeit"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

var client *mongo.Client
var ctx, _ = context.WithTimeout(context.Background(), 10*time.Minute)
var subjects = []string{"Maths", "Science", "Social Studies", "English"}
var positions = []string{"Manager", "Engineer", "Salesman", "Developer"}

var updateCount = 0
var totalUpdate int
var deleteCount = 0
var totalDelete int

var insertedData []interface{}
var insertedDataCount int

func Populate(mclient *mongo.Client, operations int) {
	client = mclient

	//extract to function
	op := (85 * operations) / 100
	op = int(math.Max(float64(op), 1))

	u := (10 * operations) / 100
	u = int(math.Max(float64(u), 1))
	totalUpdate = u

	d := (5 * operations) / 100
	d = int(math.Max(float64(d), 1))
	totalDelete = d

	dataList := generateData(op)
	rand.Seed(time.Now().UnixNano())

	for i := 0; i < len(dataList); i++ {
		if err := insertData(dataList[i]); err != nil {
			fmt.Printf("Failed to insert data at index %d: %s\n", i, err.Error())
			continue
		}

		//fix this
		insertedData = append(insertedData, dataList[i])

		//update
		if isMultipleOfSevenEightOrEleven(i) {
			if updateCount < totalUpdate {
				if err := updateData(insertedData[rand.Intn(i)]); err != nil {
					fmt.Printf("Failed to update data at index %d: %s\n", i, err.Error())
					continue
				}
			}
		}

		//delete
		if isMultipleOfTwoNineortweleve(i) {
			if deleteCount < totalDelete {
				indx := rand.Intn(i)
				if err := deleteData(insertedData[indx]); err != nil {
					fmt.Printf("Failed to delete data at index %d: %s\n", i, err.Error())
					continue
				}
				deleteCount++
				insertedData = append(insertedData[:indx], insertedData[indx:]...)
			}
		}
	}
	//insert data for alter table
	// // for i := 0; i < 3; i++ {

	// 	dataS := &StudentS{
	// 		Id:           gofakeit.UUID(),
	// 		Name:         gofakeit.FirstName() + " " + gofakeit.LastName(),
	// 		Age:          rand.Intn(10) + 18,
	// 		Subject:      subjects[rand.Intn(len(subjects))],
	// 		Is_Graduated: gofakeit.Bool(),
	// 	}
	// 	_, errS := studentsCollection.InsertOne(ctx, dataS)
	// 	if errS != nil {
	// 		log.Fatal(errS)
	// 	}

	// 	dataE := &EmployeeS{
	// 		Id:        gofakeit.UUID(),
	// 		Name:      gofakeit.FirstName() + " " + gofakeit.LastName(),
	// 		Age:       rand.Intn(30) + 20,
	// 		Salary:    rand.Float64() * 10000,
	// 		Position:  positions[rand.Intn(len(positions))],
	// 		WorkHours: rand.Intn(8) + 4,
	// 	}
	// 	_, errE := employeesCollection.InsertOne(ctx, dataE)
	// 	if errE != nil {
	// 		log.Fatal(errE)
	// 	}
	// }
}
func generateData(operations int) []interface{} {
	x := operations / 2
	var data []interface{}
	for i := 0; i < x; i++ {
		emp := &Employee{}
		empData := emp.GetData()
		data = append(data, empData)

		student := &Student{}
		studentData := student.GetData()
		data = append(data, studentData)

	}
	shuffle(data)
	return data
}

func shuffle(slice []interface{}) {
	for i := range slice {
		j := rand.Intn(i + 1)
		slice[i], slice[j] = slice[j], slice[i]
	}
}

func getType(data interface{}) reflect.Type {
	t := reflect.TypeOf(data)
	if t.Kind() == reflect.Ptr {
		t = t.Elem()
	}
	return t
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
	return n%9 == 0 || n%12 == 0
}

func insertData(data interface{}) error {
	value := reflect.ValueOf(data)
	method := value.MethodByName("GetCollection")
	if !method.IsValid() {
		return errors.New("GetCollection method not found")
	}

	result := method.Call(nil)
	if len(result) != 1 || result[0].IsNil() {
		return errors.New("collection not found")
	}
	collection := result[0].Interface().(*mongo.Collection)
	//extract

	_, err := collection.InsertOne(ctx, data)
	if err != nil {
		return err
	}
	return nil
}

func updateData(data interface{}) error {
	value := reflect.ValueOf(data)
	method := value.MethodByName("GetCollection")
	if !method.IsValid() {
		return errors.New("GetCollection method not found")
	}

	result := method.Call(nil)
	if len(result) != 1 || result[0].IsNil() {
		return errors.New("collection not found")
	}

	updateMethod := value.MethodByName("GetUpdate")
	if !updateMethod.IsValid() {
		return errors.New("GetUpdate method not found")
	}

	updateResult := updateMethod.Call(nil)
	if len(updateResult) != 1 || updateResult[0].IsNil() {
		return errors.New("update is empty or nil")
	}

	update := updateResult[0].Interface().(bson.M)
	collection := result[0].Interface().(*mongo.Collection)

	_, err := collection.UpdateOne(ctx, data, update)
	if err != nil {
		return err
	}
	return nil
}

func deleteData(data interface{}) error {
	value := reflect.ValueOf(data)
	method := value.MethodByName("GetCollection")
	if !method.IsValid() {
		return errors.New("GetCollection method not found")
	}

	result := method.Call(nil)
	if len(result) != 1 || result[0].IsNil() {
		return errors.New("collection not found")
	}
	collection := result[0].Interface().(*mongo.Collection)
	//extract

	_, err := collection.DeleteOne(ctx, data)
	if err != nil {
		return err
	}
	return nil
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

func (s *Student) GetData() interface{} {
	return &Student{
		Id:      gofakeit.UUID(),
		Name:    gofakeit.FirstName() + " " + gofakeit.LastName(),
		Age:     rand.Intn(10) + 18,
		Subject: subjects[rand.Intn(len(subjects))],
	}
}

func (s *Student) GetUpdateSet() interface{} {
	return bson.M{
		"$set": bson.M{
			"Age": rand.Intn(10) + 18,
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

func (s *Student) GetUpdate() interface{} {
	updateS := gofakeit.Bool()
	if updateS {
		updateCount++
		return s.GetUpdateSet()
	} else {
		updateCount++
		return s.GetUpdateUnset()
	}
}

func (e *Employee) GetCollection() *mongo.Collection {
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

func (e *Employee) GetUpdateSet() interface{} {
	return bson.M{
		"$set": bson.M{
			"Age": rand.Intn(10) + 18,
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

func (e *Employee) GetUpdate() interface{} {
	updateE := gofakeit.Bool()
	if updateE {
		updateCount++
		return e.GetUpdateSet()
	} else {
		updateCount++
		return e.GetUpdateUnset()
	}
}
