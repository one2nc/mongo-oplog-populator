package populator

type Employee struct {
	Id       string  `json:"Id"`
	Name     string  `json:"Name"`
	Age      int     `json:"Age"`
	Salary   float64 `json:"Salary"`
	Phone    []Phone `json:"Phone"`
	Position string  `json:"Position"`
}

type EmployeeS struct {
	Id        string  `json:"Id"`
	Name      string  `json:"Name"`
	Age       int     `json:"Age"`
	Salary    float64 `json:"Salary"`
	Phone     []Phone `json:"Phone"`
	Position  string  `json:"Position"`
	WorkHours int     `json:"WorkHours"`
}

type Student struct {
	Id      string `json:"Id"`
	Name    string `json:"Name"`
	Age     int    `json:"Age"`
	Subject string `json:"Subject"`
	// Phone         []Phone
}

type StudentS struct {
	Id      string `json:"Id"`
	Name    string `json:"Name"`
	Age     int    `json:"Age"`
	Subject string `json:"Subject"`
	// Phone         []Phone
	Is_Graduated bool `json:"Is_Graduated"`
}

type Phone struct {
	Personal string `json:"Personal"`
	Work     string `json:"Work"`
}

type OperationSize struct {
	insert, update, delete int
}
