package db

import (
	"os"

	fiberlog "github.com/gofiber/fiber/v2/log"
	"github.com/jmoiron/sqlx"
)

// SQLInjection reads the content of the init.sql file and executes it in the database if not all tables already exist.
func SQLInjection(filePath string, sqlClient *sqlx.DB) error {
	// Check if the file exists
	if _, err := os.Stat(filePath); os.IsNotExist(err) {
		fiberlog.Fatalf("File %s not found: %v", filePath, err)
		return err
	}

	// Read the content of the file
	content, err := os.ReadFile(filePath)
	if err != nil {
		fiberlog.Fatalf("Error reading file %s: %v", filePath, err)
		return err
	}

	sqlCode := string(content)

	// Execute the SQL code in the database
	_, err = sqlClient.Exec(sqlCode)
	if err != nil {
		fiberlog.Fatalf("Error executing SQL code: %v", err)
		return err
	}

	fiberlog.Info("SQL script executed successfully.")
	return nil
}
