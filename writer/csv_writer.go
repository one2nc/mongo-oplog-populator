package writer

import (
	"encoding/csv"
	"fmt"
	"log"
	"math/rand"
	"os"

	"github.com/brianvoe/gofakeit"
)

type CSVWriter struct {
	FilePath   string
	Operations int
}

func NewCSVWriter(filepath string, operations int) Writer {
	return &CSVWriter{FilePath: filepath, Operations: operations}
}

func (csvw *CSVWriter) WriteData() {
	_, err := os.Stat(csvw.FilePath)
	if os.IsNotExist(err) {
		var subjects = []string{"Maths", "Science", "Social Studies", "English"}
		var positions = []string{"Manager", "Engineer", "Salesman", "Developer"}

		file, err := os.Create(csvw.FilePath)
		if err != nil {
			log.Fatal("Could not create file:", err)
		}
		defer file.Close()

		writer := csv.NewWriter(file)
		// Add threshold for flush
		defer writer.Flush()

		header := []string{
			"First Name", "Last Name", "Subject", "Street Address", "Position",
			"ZIP", "Phone Number", "Age", "Work Hours", "Salary",
		}

		err = writer.Write(header)
		if err != nil {
			log.Fatal("Error writing header to CSV:", err)
		}

		if csvw.Operations > 50 {
			csvw.Operations = csvw.Operations / 4
		}

		for i := 0; i < csvw.Operations; i++ {
			row := []string{
				gofakeit.FirstName(),
				gofakeit.LastName(),
				subjects[rand.Intn(len(subjects))],
				gofakeit.Address().Street,
				positions[rand.Intn(len(positions))],
				gofakeit.Zip(),
				gofakeit.Phone(),
				fmt.Sprintf("%d", rand.Intn(30)+20),
				fmt.Sprintf("%d", rand.Intn(8)+4),
				fmt.Sprintf("%.2f", rand.Float64()*10000),
			}

			err = writer.Write(row)
			if err != nil {
				log.Fatal("Error writing row to CSV:", err)
			}
		}

		log.Println("CSV file created successfully.")
	}
}
