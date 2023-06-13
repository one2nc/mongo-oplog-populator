package service

type Employee struct {
	Name     string
	Age      int
	Salary   float64
	Phone    Phone
	Address  []Address
	Position string
}

type EmployeeA struct {
	Name      string
	Age       int
	Salary    float64
	Phone     Phone
	Address   []Address
	Position  string
	WorkHours int
}

type Student struct {
	Name    string
	Age     int
	Subject string
}

type StudentA struct {
	Name         string
	Age          int
	Subject      string
	Is_Graduated bool
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
	insert, update, delete int
}
