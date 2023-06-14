package service

import (
	"math/rand"
	"mongo-oplog-populator/internal/app/populator/types"

	"github.com/brianvoe/gofakeit"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func (e *Employee) GetCollection() *mongo.Collection {
	return client.Database("employee").Collection("employees")
}

func (e *Employee) GetData(attributes types.Attributes, index int) Data {
	return &Employee{
		Name:   attributes.FirstNames[index] + " " + attributes.LastNames[index],
		Age:    attributes.Ages[index],
		Salary: attributes.Salaries[index],
		Phone:  Phone{attributes.PhoneNumbers[index], attributes.PhoneNumbers[index+1]},
		Address: []Address{
			{attributes.Zips[index], attributes.StreetAddresses[index]},
			{attributes.Zips[index+1], attributes.StreetAddresses[index+1]},
		},
		Position: attributes.Positions[index%len(attributes.Positions)],
	}
}

func (e *Employee) GetUpdateSet() interface{} {
	return bson.M{"$set": bson.M{"age": rand.Intn(10) + 18}}
}

func (e *Employee) GetUpdateUnset() interface{} {
	return bson.M{"$unset": bson.M{"position": ""}}
}

func (e *Employee) GetUpdate() interface{} {
	updateE := gofakeit.Bool()
	if updateE {
		return e.GetUpdateSet()
	} else {
		return e.GetUpdateUnset()
	}
}

func (e *EmployeeA) GetCollection() *mongo.Collection {
	return client.Database("employee").Collection("employees")
}

func (e *EmployeeA) GetData(attributes types.Attributes, index int) Data {
	return &EmployeeA{
		Name:   attributes.FirstNames[index] + " " + attributes.LastNames[index],
		Age:    attributes.Ages[index],
		Salary: attributes.Salaries[index],
		Phone:  Phone{attributes.PhoneNumbers[index], attributes.PhoneNumbers[index+1]},
		Address: []Address{
			{attributes.Zips[index], attributes.StreetAddresses[index]},
			{attributes.Zips[index], attributes.StreetAddresses[index]},
		},
		Position:  attributes.Positions[index%len(attributes.Positions)],
		WorkHours: attributes.Workhours[index],
	}
}

func (e *EmployeeA) GetUpdateSet() interface{} {
	return bson.M{"$set": bson.M{"age": rand.Intn(10) + 18}}
}

func (e *EmployeeA) GetUpdateUnset() interface{} {
	return bson.M{"$unset": bson.M{"position": ""}}
}

func (e *EmployeeA) GetUpdate() interface{} {
	updateE := gofakeit.Bool()
	if updateE {
		return e.GetUpdateSet()
	} else {
		return e.GetUpdateUnset()
	}
}
