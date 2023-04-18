package populator

type Employee struct {
	Id       string    `json:"Id"`
	Name     string    `json:"Name"`
	Age      int       `json:"Age"`
	Salary   float64   `json:"Salary"`
	Phone    Phone     `json:"Phone"`
	Address  []Address `json:"Address"`
	Position string    `json:"Position"`
}

type EmployeeA struct {
	Id        string    `json:"Id"`
	Name      string    `json:"Name"`
	Age       int       `json:"Age"`
	Salary    float64   `json:"Salary"`
	Phone     Phone     `json:"Phone"`
	Address   []Address `json:"Address"`
	Position  string    `json:"Position"`
	WorkHours int       `json:"WorkHours"`
}

type Student struct {
	Id      string `json:"Id"`
	Name    string `json:"Name"`
	Age     int    `json:"Age"`
	Subject string `json:"Subject"`
	// Phone         []Phone
}

type StudentA struct {
	Id      string `json:"Id"`
	Name    string `json:"Name"`
	Age     int    `json:"Age"`
	Subject string `json:"Subject"`
	// Phone         []Phone
	Is_Graduated bool `json:"Is_Graduated"`
}

type Phone struct {
	// Id       string `json:"Id"`
	Personal string `json:"Personal"`
	Work     string `json:"Work"`
}

type Address struct {
	Zip   string `json:"zip"`
	Line1 string `json:"line1"`
}
type OperationSize struct {
	insert, update, delete int
}
