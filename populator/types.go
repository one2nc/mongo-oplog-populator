package populator

type Employee struct {
	Id       string  `json:"Id"`
	Name     string  `json:"Name"`
	Age      int     `json:"Age"`
	Salary   float64 `json:"Salary"`
	Phone    []Phone `json:"Phone"`
	Position string  `json:"Position"`
}

type EmployeeU struct {
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

type StudentU struct {
	Id      string `json:"Id"`
	Name    string `json:"Name"`
	Age     int    `json:"Age"`
	Subject string `json:"Subject"`
	// Phone         []Phone
	Is_Graduated bool `json:"Is_Graduated"`
}

type Phone struct {
	Id       string `json:"Id"`
	Personal string `json:"Personal"`
	Work     string `json:"Work"`
}

type OperationSize struct {
	insert, update, delete int
}
