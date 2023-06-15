package generator

type CustomGenerator interface {
	generateFirstName() string
	generateLastName() string
	generateSubjects() []string
	generateStreetAddress() string
	generatePositions() []string
	generateZip() string
	generatePhone() string
	generateAge() int
	generateWorkHours() int
	generateSalary() float64
}
