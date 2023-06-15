package writer

import (
	"encoding/csv"
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
	_, err := os.Stat(csvw.FilePath)
	if os.IsNotExist(err) {

		file, err := os.Create(csvw.FilePath)
		if err != nil {
			log.Fatal("Could not create file:", err)
		}
		defer file.Close()

		writer := csv.NewWriter(file)
		// TODO:Add threshold for flush
		defer writer.Flush()

		header := []string{
			"First Name", "Last Name", "Subject", "Street Address", "Position",
			"ZIP", "Phone Number", "Age", "Work Hours", "Salary",
		}

		err = writer.Write(header)
		if err != nil {
			log.Fatal("Error writing header to CSV:", err)
		}

		//TODO-DONE: generate randomInt once and modify it
		//TODO: Modify this Code

		// for i := 0; i < csvw.Operations; i++ {
		// 	//TODO:
		//
		// 	row := []string{
		// 		gofakeit.FirstName(),
		// 		gofakeit.LastName(),
		// 		subjects[rand.Intn(len(subjects))],
		// 		gofakeit.Address().Street,
		// 		positions[rand.Intn(len(positions))],
		// 		gofakeit.Zip(),
		// 		gofakeit.Phone(),
		// 		fmt.Sprintf("%d", rand.Intn(30)+20),
		// 		fmt.Sprintf("%d", rand.Intn(8)+4),
		// 		fmt.Sprintf("%.2f", rand.Float64()*10000),
		// 	}

		// 	err = writer.Write(row)
		// 	if err != nil {
		// 		log.Fatal("Error writing row to CSV:", err)
		// 	}
		// }

		log.Println("CSV file created successfully.")
	}
}
