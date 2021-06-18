package config

import (
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

var (
	// DatabaseConnection is a connection string with Arango
	DatabaseConnection = ""
	// Port on which the API will run
	Port = 0
)

// Initialize will init as environment variables
func Initialize() {
	var err error

	if err = godotenv.Load(); err != nil {
		log.Fatal(err)
	}
	Port, err = strconv.Atoi(os.Getenv("API_PORT"))
	if err != nil {
		Port = 9000
	}

}
