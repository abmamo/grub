package main

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

// function to take a key of an env
// variable and return its value
func getEnvironment(key string, filename string) string {
	// load .env file
	err := godotenv.Load(filename)
	// check if loaded
	if err != nil {
		log.Fatalf("Error loading .env file")
	}
	// return value
	return os.Getenv(key)
}
