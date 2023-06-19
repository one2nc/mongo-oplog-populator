package generator

import (
	"github.com/brianvoe/gofakeit"
)

// TODO: Decide from where to take this number, .env?
var noOfFakeDataOperations int = 1000

type GoFakeItGenerator struct {
}

func NewGoFakeItGenerator() CustomDataGenerator {
	return &GoFakeItGenerator{}
}

// generateAge implements CustomGenerator
func (*GoFakeItGenerator) GenerateAge(i int) int {
	return (i + 1) % 80
}

// generateFirstName implements CustomGenerator
func (*GoFakeItGenerator) GenerateFirstName() string {
	return gofakeit.FirstName()

}

// generateLastName implements CustomGenerator
func (*GoFakeItGenerator) GenerateLastName() string {
	return gofakeit.LastName()
}

// generatePhone implements CustomGenerator
func (*GoFakeItGenerator) GeneratePhone() string {
	return gofakeit.Phone()
}

// generatePositions implements CustomGenerator
func (*GoFakeItGenerator) GeneratePosition(i int) string {
	positions := []string{"Manager", "Engineer", "Salesman", "Developer"}
	return positions[i%len(positions)]
}

// generateSalary implements CustomGenerator
func (*GoFakeItGenerator) GenerateSalary(i int) float64 {
	return float64((i + 20000) % 20000)
}

// generateStreetAddress implements CustomGenerator
func (*GoFakeItGenerator) GenerateStreetAddress() string {
	return gofakeit.Street()
}

// generateSubjects implements CustomGenerator
func (*GoFakeItGenerator) GenerateSubject(i int) string {
	subjects := []string{"Maths", "Science", "Social Studies", "English"}
	return subjects[i%len(subjects)]
}

// generateWorkHours implements CustomGenerator
func (*GoFakeItGenerator) GenerateWorkHours(i int) int {
	return (i + 1) % 13
}

// generateZip implements CustomGenerator
func (*GoFakeItGenerator) GenerateZip() string {
	return gofakeit.Zip()
}

// TODO-DONE: generate 1000 data
func (g *GoFakeItGenerator) GenerateFakeData() FakeData {
	var personnelInfo FakeData
	for i := 0; i < noOfFakeDataOperations; i++ {
		personnelInfo.FirstNames = append(personnelInfo.FirstNames, g.GenerateFirstName())
		personnelInfo.LastNames = append(personnelInfo.LastNames, g.GenerateLastName())
		personnelInfo.Subjects = append(personnelInfo.Subjects, g.GenerateSubject(i))
		personnelInfo.StreetAddresses = append(personnelInfo.StreetAddresses, g.GenerateStreetAddress())
		personnelInfo.Positions = append(personnelInfo.Positions, g.GeneratePosition(i))
		personnelInfo.Zips = append(personnelInfo.Zips, g.GenerateZip())
		personnelInfo.PhoneNumbers = append(personnelInfo.PhoneNumbers, g.GeneratePhone())
		personnelInfo.Ages = append(personnelInfo.Ages, g.GenerateAge(i))
		personnelInfo.Workhours = append(personnelInfo.Workhours, g.GenerateWorkHours(i))
		personnelInfo.Salaries = append(personnelInfo.Salaries, g.GenerateSalary(i))
	}

	return personnelInfo
}
