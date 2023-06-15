package writer

import (
	"encoding/csv"
	"fmt"
	"log"
	"mongo-oplog-populator/internal/app/populator/types"
	"os"
)

type CSVWriter struct {
	FilePath string
}

func NewCSVWriter(filepath string) Writer {
	return &CSVWriter{FilePath: filepath}
}

func (csvw *CSVWriter) WriteData(personnelInfo types.PersonnelInfo) {
	file, err := os.Create(csvw.FilePath)
	if err != nil {
		log.Fatal("Could not create file:", err)
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	// TODO-DONE:Add threshold for flush

	//TODO-DONE: create a method for CSV headers in CSV writer
	headers := getCSVHeaders()
	err = writer.Write(headers)
	if err != nil {
		log.Fatal("Error writing header to CSV:", err)
	}

	//TODO-DONE: generate randomInt once and modify it
	//TODO-DONE: Modify this Code

	for i := 0; i < len(personnelInfo.FirstNames); i++ {
		row := []string{
			personnelInfo.FirstNames[i],
			personnelInfo.LastNames[i],
			personnelInfo.Subjects[i],
			personnelInfo.StreetAddresses[i],
			personnelInfo.Positions[i],
			personnelInfo.Zips[i],
			personnelInfo.PhoneNumbers[i],
			fmt.Sprintf("%d", personnelInfo.Ages[i]),
			fmt.Sprintf("%d", personnelInfo.Workhours[i]),
			fmt.Sprintf("%f", personnelInfo.Salaries[i]),
		}
		err = writer.Write(row)
		if err != nil {
			log.Fatal("Error writing row to CSV:", err)
		}
		if i%10 == 0 {
			writer.Flush()
		}
	}
	writer.Flush()
	log.Println("CSV file created successfully.")
}

func getCSVHeaders() []string {
	return []string{
		"First Name", "Last Name", "Subject", "Street Address", "Position",
		"ZIP", "Phone Number", "Age", "Work Hours", "Salary",
	}
}
