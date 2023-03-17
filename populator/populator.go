package populator

import (
	"context"
	"math"
	"strconv"

	"github.com/brianvoe/gofakeit"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

var database, table string

const (
	insert = "i"
	update = "u"
	delete = "d"
)

func MakePopulateJson(mongoConnection *mongo.Collection, args []string) error {
	database = args[1]
	table = args[2]
	size, _ := strconv.Atoi(args[3])
	d := (5 * size) / 100
	d = int(math.Max(float64(d), 1))

	u := (10 * size) / 100
	u = int(math.Max(float64(u), 1))

	a := 0
	b := 0

	for i := 0; i < size; i++ {
		id := gofakeit.UUID()
		name := gofakeit.Name()
		stud := StudentInfo{Id: id, Name: name, Roll_no: gofakeit.Number(0, 50), Is_Graduated: gofakeit.Bool(), Date_Of_Birth: gofakeit.Date().Format("02-01-2006")}
		_, err := mongoConnection.InsertOne(context.Background(), stud)
		if err != nil {
			panic(err)
		}

		if b <= u {
			filter := bson.M{"Id": id, "Name": name}
			updateSet := bson.M{
				"$set": bson.M{
					"Name":         gofakeit.Name(),
					"Is_Graduated": gofakeit.Bool(),
				},
			}
			updateUnset := bson.M{
				"$unset": bson.M{
					"Date_Of_Birth": "",
				},
			}

			randomBool := gofakeit.Bool()
			var update interface{}
			if randomBool {
				update = updateSet
			} else {
				update = updateUnset
			}
			_, errUpdate := mongoConnection.UpdateOne(context.Background(), filter, update)
			if errUpdate != nil {
				panic(errUpdate)
			}
			b++
		}

		if a <= d {
			_, errDelete := mongoConnection.DeleteOne(context.Background(), stud)
			if errDelete != nil {
				panic(errDelete)
			}
			a++
		}
	}
	return nil
}
