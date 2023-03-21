package populator

type Employee struct {
	Id       string  `json:"Id,omitempty"`
	Name     string  `json:"Name,omitempty"`
	Age      int     `json:"Age,omitempty"`
	Salary   float64 `json:"Salary,omitempty"`
	Phone    []Phone `json:"Phone,omitempty"`
	Position string  `json:"Position,omitempty"`
}

type EmployeeS struct {
	Id        string  `json:"Id,omitempty"`
	Name      string  `json:"Name,omitempty"`
	Age       int     `json:"Age,omitempty"`
	Salary    float64 `json:"Salary,omitempty"`
	Phone     []Phone `json:"Phone,omitempty"`
	Position  string  `json:"Position,omitempty"`
	WorkHours int     `json:"WorkHours,omitempty"`
}

type Student struct {
	Id      string `json:"Id,omitempty"`
	Name    string `json:"Name,omitempty"`
	Age     int    `json:"Age,omitempty"`
	Subject string `json:"Subject,omitempty"`
	// Phone         []Phone
}

type StudentS struct {
	Id      string `json:"Id,omitempty"`
	Name    string `json:"Name,omitempty"`
	Age     int    `json:"Age,omitempty"`
	Subject string `json:"Subject,omitempty"`
	// Phone         []Phone
	Is_Graduated bool `json:"Is_Graduated,omitempty"`
}

type Phone struct {
	Personal string `json:"Personal,omitempty"`
	Work     string `json:"Work,omitempty"`
}
