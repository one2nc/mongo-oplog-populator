package populator

type Employee struct {
	Name     string
	Age      int
	Salary   float64
	Phone    []Phone
	Position string
}

type EmployeeU struct {
	Name      string
	Age       int
	Salary    float64
	Phone     []Phone
	Position  string
	WorkHours int
}

type Student struct {
	Name    string
	Age     int
	Subject string
}

type StudentU struct {
	Name         string
	Age          int
	Subject      string
	Is_Graduated bool
}

type Phone struct {
	Id       string
	Personal string
	Work     string
}

type OperationSize struct {
	insert, update, delete int
}
