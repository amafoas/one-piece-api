package utils

import (
	"os"
)

func GetEnvVariable(key string) string {
	// err := godotenv.Load()
	// if err != nil {
	// 	log.Fatalf("Error loading .env file")
	// }
	return os.Getenv(key)
}
