package populator

import (
	"context"
	"log"
	"math/rand"

	"github.com/brianvoe/gofakeit"
	"go.mongodb.org/mongo-driver/mongo"
)

var database, table string

const (
	insert = "i"
	update = "u"
	delete = "d"
)

func MakePopulateJson(client *mongo.Client, operations int) {
	subjects := []string{"Maths", "Science", "Social Studies", "English"}
	positions := []string{"Manager", "Engineer", "Salesman", "Developer"}

	studentsCount := operations / 2
	employeesCount := operations - studentsCount
	studentsCollection := client.Database("student").Collection("students")
	employeesCollection := client.Database("Employee").Collection("employees")

	// ds := (5 * studentsCount) / 100
	// ds = int(math.Max(float64(d), 1))
	// a := 0
	// b := 0
	for i := 0; i < studentsCount; i++ {
		ctx := context.Background()
		id := gofakeit.UUID()
		data := &Student{
			Id:      id,
			Name:    gofakeit.FirstName() + " " + gofakeit.LastName(),
			Age:     rand.Intn(10) + 18,
			Subject: subjects[rand.Intn(len(subjects))],
		}
		_, errS := studentsCollection.InsertOne(ctx, data)
		if errS != nil {
			log.Fatal(errS)
		}

		//Add uuid to array
		//Add delete here
		// if a <= ds {
		// 	_, errD := studentsCollection.DeleteOne(ctx, id)
		// 	if errD != nil {
		// 		log.Fatal(errD)
		// 	}
		// 	a++
		// }
	}

	for i := 0; i < employeesCount; i++ {
		ctx := context.Background()
		id := gofakeit.UUID()
		data := &Employee{
			Id:       id,
			Name:     gofakeit.FirstName() + " " + gofakeit.LastName(),
			Age:      rand.Intn(30) + 20,
			Salary:   rand.Float64() * 10000,
			Position: positions[rand.Intn(len(positions))],
		}
		_, errE := employeesCollection.InsertOne(ctx, data)
		if errE != nil {
			log.Fatal(errE)
		}

		//add delete here
		_, errDe := employeesCollection.DeleteOne(ctx, id)
		if errDe != nil {
			log.Fatal(errDe)
		}
	}

	//insert data for alter table
	for i := 0; i < 3; i++ {
		ctx := context.Background()
		dataS := &StudentS{
			Id:           gofakeit.UUID(),
			Name:         gofakeit.FirstName() + " " + gofakeit.LastName(),
			Age:          rand.Intn(10) + 18,
			Subject:      subjects[rand.Intn(len(subjects))],
			Is_Graduated: gofakeit.Bool(),
		}
		_, errS := studentsCollection.InsertOne(ctx, dataS)
		if errS != nil {
			log.Fatal(errS)
		}

		dataE := &EmployeeS{
			Id:        gofakeit.UUID(),
			Name:      gofakeit.FirstName() + " " + gofakeit.LastName(),
			Age:       rand.Intn(30) + 20,
			Salary:    rand.Float64() * 10000,
			Position:  positions[rand.Intn(len(positions))],
			WorkHours: rand.Intn(8) + 4,
		}
		_, errE := employeesCollection.InsertOne(ctx, dataE)
		if errE != nil {
			log.Fatal(errE)
		}
	}

}
