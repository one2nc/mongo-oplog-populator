package reader

import (
	"encoding/csv"
	"log"
	"mongo-oplog-populator/internal/app/populator/types"

	"os"
	"strconv"
)

type CSVReader struct {
	FilePath string
}

func NewCSVReader(filepath string) Reader {
	return &CSVReader{
		FilePath: filepath,
	}
}

func (csvr *CSVReader) ReadData() types.Attributes {
	var attributes types.Attributes
	file, err := os.Open(csvr.FilePath)
	if err != nil {
		log.Fatal("Could not open file:", err)
	}
	defer file.Close()
	reader := csv.NewReader(file)
	records, err := reader.ReadAll()
	if err != nil {
		log.Fatal("Error reading CSV:", err)
	}

	//ignoring the headers from CSV
	records = records[1:]

	for _, row := range records {
		attributes.FirstNames = append(attributes.FirstNames, row[0])
		attributes.LastNames = append(attributes.LastNames, row[1])
		attributes.Subjects = append(attributes.Subjects, row[2])
		attributes.StreetAddresses = append(attributes.StreetAddresses, row[3])
		attributes.Positions = append(attributes.Positions, row[4])
		attributes.Zips = append(attributes.Zips, row[5])
		attributes.PhoneNumbers = append(attributes.PhoneNumbers, row[6])

		age, err := strconv.Atoi(row[7])
		if err != nil {
			log.Fatal("Error converting age:", err)
		}
		attributes.Ages = append(attributes.Ages, age)

		workhours, err := strconv.Atoi(row[8])
		if err != nil {
			log.Fatal("Error converting work hours:", err)
		}
		attributes.Workhours = append(attributes.Workhours, workhours)

		salary, err := strconv.ParseFloat(row[9], 64)
		if err != nil {
			log.Fatal("Error converting salary:", err)
		}
		attributes.Salaries = append(attributes.Salaries, salary)
	}

	return attributes
}
