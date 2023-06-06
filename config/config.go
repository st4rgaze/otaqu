package config

import (
	"log"

	"github.com/joho/godotenv"
)

// load .env file
func LoadEnv() error {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file")
		return err
	}

	return nil
}
