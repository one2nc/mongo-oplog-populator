package service

import (
	"math/rand"
	"mongo-oplog-populator/internal/app/populator/types"

	"github.com/brianvoe/gofakeit"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func (s *Student) GetCollection() *mongo.Collection {
	return client.Database("student").Collection("students")
}

func (s *Student) GetData(attributes types.Attributes, index int) Data {
	return &Student{
		Name:    attributes.FirstNames[index] + " " + attributes.LastNames[index],
		Age:     attributes.Ages[index],
		Subject: attributes.Subjects[index%len(attributes.Subjects)],
	}
}

func (s *Student) GetUpdateSet() interface{} {
	return bson.M{"$set": bson.M{"age": rand.Intn(10) + 18}}
}

func (s *Student) GetUpdateUnset() interface{} {
	return bson.M{"$unset": bson.M{"subject": ""}}
}

func (s *Student) GetUpdate() interface{} {
	updateS := gofakeit.Bool()
	if updateS {
		return s.GetUpdateSet()
	} else {
		return s.GetUpdateUnset()
	}
}

func (s *StudentA) GetCollection() *mongo.Collection {
	return client.Database("student").Collection("students")
}

func (s *StudentA) GetData(attributes types.Attributes, index int) Data {
	return &StudentA{
		Name:         attributes.FirstNames[index] + " " + attributes.LastNames[index],
		Age:          attributes.Ages[index],
		Subject:      attributes.Subjects[index%len(attributes.Subjects)],
		Is_Graduated: gofakeit.Bool(),
	}
}
func (s *StudentA) GetUpdateSet() interface{} {
	return bson.M{"$set": bson.M{"age": rand.Intn(10) + 18}}
}

func (s *StudentA) GetUpdateUnset() interface{} {
	return bson.M{"$unset": bson.M{"subject": ""}}
}

func (s *StudentA) GetUpdate() interface{} {
	updateS := gofakeit.Bool()
	if updateS {
		return s.GetUpdateSet()
	} else {
		return s.GetUpdateUnset()
	}
}
