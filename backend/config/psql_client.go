package config

import (
	"fmt"
	"time"

	fiberlog "github.com/gofiber/fiber/v2/log"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

// GetPostgreSQLClient retrieves a PostgreSQL client.
//
// It loads the configuration from the .env file and establishes a connection with the PostgreSQL database.
// It sets the connection properties and returns the PostgreSQL client.
//
// Returns:
// - *sqlx.DB: The PostgreSQL client.
func GetPostgreSQLClient(filepath string) *sqlx.DB {
	config, err := LoadConfig(filepath)
	if err != nil {
		fiberlog.Fatalf("error loading .env from psql client: %v", err)
	}

	dbUser := config.MICRO.DB.PSQL.PSQL_USER
	dbPasswd := config.MICRO.DB.PSQL.PSQL_PASS
	dbHost := config.MICRO.DB.PSQL.PSQL_HOST
	dbPort := config.MICRO.DB.PSQL.PSQL_PORT
	dbSchema := config.MICRO.DB.PSQL.PSQL_SCHEMA

	dataSource := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", dbHost, dbPort, dbUser, dbPasswd, dbSchema)

	client, err := sqlx.Open("postgres", dataSource)
	if err != nil {
		fiberlog.Fatalf("error open connection with psql: %v", err)
	}

	err = client.Ping()
	if err != nil {
		fiberlog.Fatalf("error while ping psql connection: %v", err)
	}

	client.SetConnMaxLifetime(time.Minute * 3)
	client.SetMaxOpenConns(10)
	client.SetMaxIdleConns(10)

	fiberlog.Info("Database successfully connected!")
	return client
}
