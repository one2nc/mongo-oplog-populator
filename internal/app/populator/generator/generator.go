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
	Generate(ctx context.Context, fakeData FakeData) <-chan Data
}

func generateData(noOfOperations int, attributes FakeData) []Data {
	tempAlterOpsCnt := float64(noOfOperations) * 0.1
	alterOpsCnt := int(math.Round(tempAlterOpsCnt))

	x := (noOfOperations - alterOpsCnt*2) / 2
	var data []Data
	index := 0
	for i := 0; i < x; i++ {
		emp := &Employee{}
		empData := emp.GetData(attributes, index)
		data = append(data, empData)
		student := &Student{}
		studentData := student.GetData(attributes, index)
		data = append(data, studentData)
		index++
		//to reset if attributes size < input number of operations size. Will continue to read data in a cycle
		if index > len(attributes.FirstNames)-2 {
			index = 0
		}
	}
	dataAlterTable := generateDataAlterTable(alterOpsCnt, attributes)
	data = append(data, dataAlterTable...)
	shuffle(data)
	return data
}

func generateDataAlterTable(operations int, attributes FakeData) []Data {
	var data []Data
	index := 0
	for i := 0; i < operations; i++ {
		emp := &EmployeeA{}
		empData := emp.GetData(attributes, index)
		data = append(data, empData)
		student := &StudentA{}
		studentData := student.GetData(attributes, index)
		data = append(data, studentData)
		index++
		//to reset if attributes size < input number of operations size. Will continue to read data in a cycle
		if index > len(attributes.FirstNames)-2 {
			index = 0
		}
	}
	return data
}

func shuffle(slice []Data) {
	for i := range slice {
		j := rand.Intn(i + 1)
		slice[i], slice[j] = slice[j], slice[i]
	}
}
