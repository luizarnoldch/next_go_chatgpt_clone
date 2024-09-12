package config

type (
	CONFIG struct {
		MICRO MICRO
		ENV string
	}

	MICRO struct {
		DB  DB
		API API
	}

	DB struct {
		PSQL PSQL
	}

	API struct {
		API_HOST string
		API_PORT string
	}

	PSQL struct {
		PSQL_HOST   string
		PSQL_PORT   string
		PSQL_USER   string
		PSQL_PASS   string
		PSQL_SCHEMA string
	}
)