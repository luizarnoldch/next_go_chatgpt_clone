package config

import (
	"fmt"
	"os"

	fiberlog "github.com/gofiber/fiber/v2/log"
	"github.com/lpernett/godotenv"
)

// envSanityCheck ensures that all required environment variables are defined.
func envSanityCheck() {
	requiredEnvVariables := []string{
		"API_HOST",
		"API_PORT",

		// Uncomment these if you need database variables
		"PSQL_HOST",
		"PSQL_PORT",
		"PSQL_USER",
		"PSQL_PASS",
		"PSQL_SCHEMA",
	}

	for _, envVar := range requiredEnvVariables {
		if value := os.Getenv(envVar); value == "" {
			fiberlog.Fatalf("Environment variable %s not defined. Please set up the application again.", envVar)
		}
	}
}

// LoadConfig loads environment variables from a .env file and checks for required variables.
func LoadConfig(filePath string) (*CONFIG, error) {
	err := godotenv.Load(filePath)
	if err != nil {
		fiberlog.Fatalf("Error loading .env file: %v", err)
		return nil, fmt.Errorf("error loading .env file: %v", err)
	}

	envSanityCheck()

	return &CONFIG{
		MICRO: MICRO{
			API: API{
				API_HOST: os.Getenv("API_HOST"),
				API_PORT: os.Getenv("API_PORT"),
			},
			DB: DB{
				PSQL: PSQL{
					PSQL_HOST:   os.Getenv("PSQL_HOST"),
					PSQL_PORT:   os.Getenv("PSQL_PORT"),
					PSQL_USER:   os.Getenv("PSQL_USER"),
					PSQL_PASS:   os.Getenv("PSQL_PASS"),
					PSQL_SCHEMA: os.Getenv("PSQL_SCHEMA"),
				}},
		},
		ENV: os.Getenv("ENV"),
	}, nil
}
