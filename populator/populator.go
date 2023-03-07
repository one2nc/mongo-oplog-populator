package populator

import (
	"encoding/json"
	"fmt"
	"strconv"

	"github.com/brianvoe/gofakeit/v6"
	"go.mongodb.org/mongo-driver/mongo"
)

var database, table string

const (
	insert = "i"
	update = "u"
	delete = "d"
)

func MakeInsertJson(mongoConnection *mongo.Collection, args []string) error {
	database = args[1]
	table = args[2]
	size, _ := strconv.Atoi(args[3])
	for i := 0; i < size; i++ {
		stud := StudentInfo{Id: gofakeit.UUID(), Name: gofakeit.Name(), Roll_no: gofakeit.Number(0, 50), Is_Graduated: gofakeit.Bool(), Date_Of_Birth: gofakeit.Date().Format("02-01-2006")}
		insertData := Oplog{Type: insert, Namespace: database + "." + table, StudentInfo: stud}
		jsonData, err := convertToJson(insertData)
		if err != nil {
			fmt.Println("Error in MakeInserJson. The error is:", err)
		}
		fmt.Println("json data is:", jsonData)
	}
	// _, err = mongoConnection.InsertOne(context.Background(), jsonData)
	// if err != nil {
	// 	fmt.Println("Error inserting oplog data:", err)
	// 	return err
	// }

	// fmt.Println("Oplog data inserted successfully!")
	return nil
}

func MakeDeleteJson(mongoConnection *mongo.Collection, args []string) error {
	database = args[1]
	table = args[2]
	size, _ := strconv.Atoi(args[3])
	for i := 0; i < size; i++ {
		stud := StudentInfo{Id: gofakeit.UUID()}
		insertData := Oplog{Type: delete, Namespace: database + "." + table, StudentInfo: stud}
		jsonData, err := convertToJson(insertData)
		if err != nil {
			fmt.Println("Error in MakeDelete. The error is:", err)
		}
		fmt.Println("json data is:", jsonData)
	}
	return nil
}

func MakeUpdateJson(mongoConnection *mongo.Collection, args []string) {
	status := make([]string, 0)
	status = append(status, "$set")
	status = append(status, "$unset")

}

func MakeInsertAllJson(mongoConnection *mongo.Collection, args []string) error {
	database = args[1]
	table = args[2]
	size, _ := strconv.Atoi(args[3])
	for i := 0; i < size; i++ {
		stud := StudentInfo{Id: gofakeit.UUID(), Name: gofakeit.Name(), Roll_no: gofakeit.Number(0, 50), Is_Graduated: gofakeit.Bool(), Date_Of_Birth: gofakeit.Date().Format("02-01-2006"), Address: []Address{{Line1: gofakeit.Address().Address, Line2: gofakeit.Address().Address, Zip: gofakeit.Zip()}}, Phone: []Phone{{Personal: gofakeit.Phone(), Work: gofakeit.Phone()}}, Email: gofakeit.Email()}
		insertData := Oplog{Type: insert, Namespace: database + "." + table, StudentInfo: stud}
		jsonData, err := convertToJson(insertData)
		if err != nil {
			fmt.Println("Error in MakeInsertAllJson. The error is:", err)
			return err
		}
		fmt.Println("json data is:", jsonData)
	}
	return nil
}

func convertToJson(insertData Oplog) (string, error) {
	jsonData, err := json.Marshal(insertData)
	if err != nil {
		fmt.Println("Error marshaling JSON:", err)
		return "", err
	}
	return string(jsonData), nil
}
