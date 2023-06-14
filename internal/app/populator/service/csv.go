package service

import (
	"encoding/csv"
	"fmt"
	"mongo-oplog-populator/internal/app/populator/types"

	"log"
	"os"
)

func CreateCSVFile(csvFileName string, numberOfOperations int, attributes types.Attributes) {
	file, err := os.Create(csvFileName)
	if err != nil {
		log.Fatal("Could not create file:", err)
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	//add threshold for flush
	defer writer.Flush()

	header := []string{
		"First Name", "Last Name", "Subject", "Street Address", "Position",
		"ZIP", "Phone Number", "Age", "Work Hours", "Salary",
	}

	err = writer.Write(header)
	if err != nil {
		log.Fatal("Error writing header to CSV:", err)
	}

	for i := 0; i < len(attributes.FirstNames); i++ {
		row := []string{
			attributes.FirstNames[i],
			attributes.LastNames[i],
			attributes.Subjects[i],
			attributes.StreetAddresses[i],
			attributes.Positions[i],
			attributes.Zips[i],
			attributes.PhoneNumbers[i],
			fmt.Sprintf("%d", attributes.Ages[i]),
			fmt.Sprintf("%d", attributes.Workhours[i]),
			fmt.Sprintf("%.2f", attributes.Salaries[i]),
		}
		err = writer.Write(row)
		if err != nil {
			log.Fatal("Error writing row to CSV:", err)
		}
	}
	log.Println("CSV file created successfully.")
}
