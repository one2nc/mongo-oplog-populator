package generator

import "mongo-oplog-populator/internal/app/populator/types"

type CustomGenerator interface {
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
	GenerateFakeData() types.PersonnelInfo
}
