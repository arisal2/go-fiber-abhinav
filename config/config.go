package config

import (
	"fmt"
	"os"
	"github.com/joho/godotenv"
)

// Config function to get env value
func Config(key string) string {
	err := godotenv.Load(".env")
	if err != nil {
		fmt.Print("Error loading .env file")
	}
	return os.Getenv(key)
}