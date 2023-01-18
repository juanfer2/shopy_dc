package utilities

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

func GetEnv(key string) string {
	// load .env file
	return os.Getenv(key)
}

func ValidateEnvFile() error {
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	return err
}

func GetUrlServer() string {
	return "http://localhost:" + GetEnv("PORT")
}
