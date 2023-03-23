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
			var deletecollection *mongo.Collection
			var updatecollection *mongo.Collection
			var update interface{}
			var updateSet interface{}
			var updateUnset interface{}

			var data interface{}
			var sid string
			randomBool := gofakeit.Bool()

			aMutex.Lock()
			defer aMutex.Unlock()

			if randomBool {
				// insert a student
				sid = gofakeit.UUID()
				sname := gofakeit.FirstName() + " " + gofakeit.LastName()
				data = &Student{
					Id:      sid,
					Name:    sname,
					Age:     rand.Intn(10) + 18,
					Subject: subjects[rand.Intn(len(subjects))],
				}

				collection = studentsCollection
				deleteBoolS := gofakeit.Bool()
				studentUpdate := gofakeit.Bool()
				if studentUpdate {
					if b < u {
						updatecollection = studentsCollection
						// updateSet = bson.D{{"$set", bson.D{{"Name", gofakeit.Name()}}}}
						// updateUnset = bson.D{{"$unset", bson.D{{"Name", gofakeit.Name()}}}}

						updateSet = bson.M{
							"$set": bson.M{
								"Age":  rand.Intn(10) + 18,
							},
						}

						updateUnset = bson.M{
							"$unset": bson.M{
								"Subject": "",
							},
						}
						updateS := gofakeit.Bool()
						if updateS {
							update = updateSet
						} else {
							update = updateUnset
						}

						b++
					}

				}
				if deleteBoolS {
					if a < d {
						deletecollection = studentsCollection
						a++

					}
				}
			} else {
				// insert an employee
				eid := gofakeit.UUID()
				ename := gofakeit.FirstName() + " " + gofakeit.LastName()
				phone := []Phone{{gofakeit.Phone(), gofakeit.Phone()}}
				data = &Employee{
					Id:       eid,
					Name:     ename,
					Age:      rand.Intn(30) + 20,
					Salary:   rand.Float64() * 10000,
					Phone:    phone,
					Position: positions[rand.Intn(len(positions))],
				}

				collection = employeesCollection
				deleteBoolE := gofakeit.Bool()
				employeeUpdate := gofakeit.Bool()

				if employeeUpdate {
					if b < u {
						updatecollection = employeesCollection
						// updateSet = bson.D{{"$set", bson.D{{"Name", gofakeit.Name()}}}}
						// updateUnset = bson.D{{"$unset", bson.D{{"Name", ""}}}}

						updateSet = bson.M{
							"$set": bson.M{
								"Age": rand.Intn(10) + 18,
							},
						}
						updateUnset = bson.M{
							"$unset": bson.M{
								"Position": " ",
							},
						}
						updateE := gofakeit.Bool()
						if updateE {
							update = updateSet
						} else {
							update = updateUnset
						}

						b++
					}
				}

				if deleteBoolE {
					if a < d {
						deletecollection = employeesCollection
						a++
					}
				}
			}

			//Insert Data
			_, err := collection.InsertOne(ctx, data)
			if err != nil {
				log.Fatal(err)
			}

			if updatecollection != nil {
				_, errUpdate := updatecollection.UpdateOne(ctx, data, update)
				if errUpdate != nil {
					log.Fatal(errUpdate)
				}
			}

			if deletecollection != nil {
				_, errDelete := deletecollection.DeleteOne(ctx, data)
				if errDelete != nil {
					log.Fatal(errDelete)
				}
			}
			ch <- sid // signal that an insertion has been completed
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

		//add random data here too
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
