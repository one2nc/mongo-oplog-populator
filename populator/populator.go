package populator

import (
	"context"
	"log"
	"math"
	"math/rand"
	"sync"

	"github.com/brianvoe/gofakeit"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func MakePopulateJson(client *mongo.Client, operations int) {
	ctx := context.Background()

	subjects := []string{"Maths", "Science", "Social Studies", "English"}
	positions := []string{"Manager", "Engineer", "Salesman", "Developer"}

	op := (85 * operations) / 100
	op = int(math.Max(float64(op), 1))

	d := (5 * operations) / 100
	d = int(math.Max(float64(d), 1))

	u := (10 * operations) / 100
	u = int(math.Max(float64(u), 1))

	studentsCollection := client.Database("student").Collection("students")
	employeesCollection := client.Database("Employee").Collection("employees")

	a := 0
	b := 0
	aMutex := &sync.Mutex{}

	ch := make(chan string, op)
	for i := 0; i < op; i++ {
		go func() {
			var collection *mongo.Collection
			var data interface{}
			var id string
			randomBool := gofakeit.Bool()

			aMutex.Lock()
			defer aMutex.Unlock()

			if randomBool {
				// insert a student
				id = gofakeit.UUID()
				name := gofakeit.FirstName() + " " + gofakeit.LastName()
				data = &Student{
					Id:      id,
					Name:    name,
					Age:     rand.Intn(10) + 18,
					Subject: subjects[rand.Intn(len(subjects))],
				}

				collection = studentsCollection
				deleteBool := gofakeit.Bool()
				studentUpdate := gofakeit.Bool()
				if studentUpdate {
					if b < u {
						filter := bson.M{"Id": id}
						// r := studentsCollection.FindOne(ctx, filter)
						updateSet := bson.M{
							"$set": bson.M{
								"Name": gofakeit.Name(),
								"Age":  rand.Intn(10) + 18,
							},
						}

						updateUnset := bson.M{
							"$unset": bson.M{
								"Subject": "",
							},
						}

						updateB := gofakeit.Bool()
						var update interface{}
						if updateB {
							update = updateSet
						} else {
							update = updateUnset
						}
						_, errUpdate := studentsCollection.UpdateOne(context.Background(), filter, update)
						if errUpdate != nil {
							panic(errUpdate)
						}

						b++
					}
				}

				if deleteBool {
					if a < d {
						filter := bson.M{"Id": id, "Name": name}
						r := studentsCollection.FindOne(ctx, filter)
						if r != nil {
							_, errDelete := studentsCollection.DeleteOne(ctx, data)
							if errDelete != nil {
								panic(errDelete)
							}
							a++
						}
					}
				}

			} else {
				// insert an employee
				id := gofakeit.UUID()
				name := gofakeit.FirstName() + " " + gofakeit.LastName()
				phone := []Phone{{gofakeit.Phone(), gofakeit.Phone()}}
				data = &Employee{
					Id:       id,
					Name:     name,
					Age:      rand.Intn(30) + 20,
					Salary:   rand.Float64() * 10000,
					Phone:    phone,
					Position: positions[rand.Intn(len(positions))],
				}

				collection = employeesCollection
				deleteBool := gofakeit.Bool()
				employeeUpdate := gofakeit.Bool()

				if employeeUpdate {
					if b < u {
						filter := bson.M{"Id": id, "Name": name}
						updateSet := bson.M{
							"$set": bson.M{
								"Name": gofakeit.Name(),
							},
						}

						updateUnset := bson.M{
							"$unset": bson.M{
								"Position": "",
							},
						}

						updateB := gofakeit.Bool()
						var update interface{}
						if updateB {
							update = updateSet
						} else {
							update = updateUnset
						}
						_, errUpdate := employeesCollection.UpdateOne(context.Background(), filter, update)
						if errUpdate != nil {
							panic(errUpdate)
						}

						b++
					}
				}

				if deleteBool {
					if a < d {
						// employeesCollection.FindOne(ctx, data)
						_, errDelete := studentsCollection.DeleteOne(ctx, data)
						if errDelete != nil {
							panic(errDelete)
						}
						a++
					}
				}
			}

			//Insert Data
			_, err := collection.InsertOne(ctx, data)
			if err != nil {
				log.Fatal(err)
			}

			ch <- id // signal that an insertion has been completed
		}()
	}

	// Wait for all insertions to be completed
	for i := 0; i < op; i++ {
		<-ch // wait until an insertion is completed
	}
	close(ch) // close the channel

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
