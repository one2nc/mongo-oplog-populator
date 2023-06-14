package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

const (
	MONGO_URI     = "MONGO_URI"
	CSV_FILE_NAME = "CSV_FILE_NAME"
	ENV           = ".env"
)

type Config struct {
	MongoUri    string
	CsvFileName string
}

func Load() Config {
	err := godotenv.Load(ENV)
	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	cfg := Config{
		MongoUri:    ReadFromEnvFile(MONGO_URI),
		CsvFileName: ReadFromEnvFile(CSV_FILE_NAME),
	}
	return cfg
}

func ReadFromEnvFile(key string) string {
	return os.Getenv(key)
}
