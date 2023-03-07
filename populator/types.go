package populator

type StudentInfo struct {
	Id            string    `json:"Id"`
	Name          string    `json:"Name,omitempty"`
	Roll_no       int       `json:"Roll_no,omitempty"`
	Is_Graduated  bool      `json:"Is_Graduated,omitempty"`
	Date_Of_Birth string    `json:"Date_Of_Birth,omitempty"`
	Address       []Address `json:"Address,omitempty"`
	Phone         []Phone   `json:"Phone,omitempty"`
	Email         string    `json:"Email,omitempty"`
}
type Oplog struct {
	Type        string      `json:"op"`
	Namespace   string      `json:"ns"`
	StudentInfo StudentInfo `json:"o"`
}

type Address struct {
	Line1 string `json:"line1,omitempty"`
	Line2 string `json:"line2,omitempty"`
	Zip   string `json:"zip,omitempty"`
}

type Phone struct {
	Personal string `json:"personal,omitempty"`
	Work     string `json:"work,omitempty"`
}
