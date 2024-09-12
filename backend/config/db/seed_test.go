package db_test

import (
	"main/config"
	"main/config/db"
	"testing"
)

func TestSQLInyection(t *testing.T) {

	sql_client := config.GetPostgreSQLClient("../../.env")

	err := db.SQLInjection("init.sql", sql_client)

	if err != nil {
		t.Error("Erorr")
	}
}
