package gomastodonstats

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

// Returns valid DSL for PSQL
func getMySQLConnectionString(schema string) string {
	return fmt.Sprintf(
		"%s:%s@%s/%s",
		MYSQL_USER,
		MYSQL_PASS,
		MYSQL_HOST,
		schema,
	)
}

// Returns DB connection
func getMySQLConnection(schema string) (*sql.DB, error) {
	return sql.Open("mysql", getMySQLConnectionString(schema))
}

func runMySqlIntQuery(schema string, q string) (int, error) {
	var res int
	db, err := getMySQLConnection(schema)
	if err != nil {
		return res, err
	}

	rows, err := db.Query(q)
	if err != nil {
		return res, err
	}
	defer rows.Close()

	for rows.Next() {
		rows.Scan(&res)
	}

	return res, err
}
