package config

import (
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

const (
	MONGO_URI      = "MONGO_URI"
	CSV_FILE_NAME  = "CSV_FILE_NAME"
	MAX_DB         = "MAX_DB"
	MAX_COLLECTION = "MAX_COLLECTION"
	ENV            = ".env"
)

type Config struct {
	MongoUri       string
	CsvFileName    string
	MaxDatabases   int
	MaxCollections int
}

func Load() Config {
	err := godotenv.Load(ENV)
	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	maxDBs, err := strconv.Atoi(ReadFromEnvFile(MAX_DB))
	if err != nil {
		log.Printf("Invalid env variable MAX_DB")
		maxDBs = 1
	}

	maxCollections, err := strconv.Atoi(ReadFromEnvFile(MAX_COLLECTION))
	if err != nil {
		log.Printf("Invalid env variable MAX_COLLECTION")
		maxCollections = 1
	}

	cfg := Config{
		MongoUri:       ReadFromEnvFile(MONGO_URI),
		CsvFileName:    ReadFromEnvFile(CSV_FILE_NAME),
		MaxDatabases:   maxDBs,
		MaxCollections: maxCollections,
	}
	return cfg
}

func ReadFromEnvFile(key string) string {
	return os.Getenv(key)
}
