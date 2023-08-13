package generator

import (
	"context"
	"math"
	"math/rand"
)

type DataGenerator interface {
	GenerateFirstName() string
	GenerateLastName() string
	GenerateSubject(i int) string
	GenerateStreetAddress() string
	GeneratePosition(i int) string
	GenerateZip() string
	GeneratePhone() string
	GenerateAge(i int) int
	GenerateWorkHours(i int) int
	GenerateSalary(i int) float64
	GenerateFakeData() FakeData
}

type Generator interface {
	GenerateDocument(ctx context.Context, fakeData FakeData) <-chan Document
}

func generateDocument(noOfOperations int, attributes FakeData) []Document {
	tempAlterOpsCnt := float64(noOfOperations) * 0.1
	alterOpsCnt := int(math.Round(tempAlterOpsCnt))

	insertOpCnt := noOfOperations - alterOpsCnt
	var docs []Document
	for index := 0; index < insertOpCnt; index++ {
		doc := Document{
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
		docs = append(docs, doc)
	}

	docsAlterTable := generateDocsAlterTable(alterOpsCnt, attributes)
	docs = append(docs, docsAlterTable...)
	shuffleDocs(docs)
	return docs
}

func generateDocsAlterTable(operations int, attributes FakeData) []Document {
	var docs []Document
	for index := 0; index < operations; index++ {
		doc := Document{
			Name:   attributes.FirstNames[index] + " " + attributes.LastNames[index],
			Age:    attributes.Ages[index],
			Salary: attributes.Salaries[index],
			Phone:  Phone{attributes.PhoneNumbers[index], attributes.PhoneNumbers[index+1]},
			Address: []Address{
				{attributes.Zips[index], attributes.StreetAddresses[index]},
				{attributes.Zips[index+1], attributes.StreetAddresses[index+1]},
			},
			Position:  attributes.Positions[index%len(attributes.Positions)],
			WorkHours: attributes.Workhours[index],
		}
		docs = append(docs, doc)
	}
	return docs
}

func shuffleDocs(slice []Document) {
	for i := range slice {
		j := rand.Intn(i + 1)
		slice[i], slice[j] = slice[j], slice[i]
	}
}
