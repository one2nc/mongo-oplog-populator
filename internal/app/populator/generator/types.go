package generator

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
	Insert, Update, Delete int
}

type PersonnelInfo struct {
	FirstNames, LastNames, Subjects, StreetAddresses, Positions, Zips, PhoneNumbers []string
	Ages, Workhours                                                                 []int
	Salaries                                                                        []float64
}
