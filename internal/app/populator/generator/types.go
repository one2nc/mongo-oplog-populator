package generator

import (
	"math/rand"

	"go.mongodb.org/mongo-driver/bson"
)

type Document struct {
	// ID       int       `bson:"_id"`
	Name      string    `bson:"name"`
	Age       int       `bson:"age"`
	Salary    float64   `bson:"salary"`
	Phone     Phone     `bson:"phone,omitempty"`
	Address   []Address `bson:"address,omitempty"`
	Position  string    `bson:"position,omitempty"`
	WorkHours int       `bson:"workhours,omitempty"`
}

type Phone struct {
	Personal string
	Work     string
}

type Address struct {
	Zip   string
	Line1 string
}
type OperationSize struct {
	Insert, Update, Delete int
}

type FakeData struct {
	FirstNames, LastNames, Subjects, StreetAddresses, Positions, Zips, PhoneNumbers []string
	Ages, Workhours                                                                 []int
	Salaries                                                                        []float64
}

func (d *Document) GetUpdate() interface{} {
	updateS := getRandBoolean()
	if updateS {
		return bson.M{"$set": bson.M{"age": rand.Intn(10) + 18}}
	} else {
		return bson.M{"$unset": bson.M{"salary": ""}}
	}
}

func getRandBoolean() bool {
	ri := rand.Intn(20)
	if (5*ri)%3 == 0 {
		return false
	}
	return true
}
