package gomastodonstats

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

// Returns valid DSL for PSQL
func getConnectionString(schema string) string {
	return fmt.Sprintf(
		"postgresql://%s:%s@%s/%s?sslmode=disable",
		POSTGRESQL_USER,
		POSTGRESQL_PASS,
		POSTGRESQL_HOST,
		schema,
	)
}

// Returns DB connection
func getConnection(schema string) (*sql.DB, error) {
	return sql.Open("postgres", getConnectionString(schema))
}

// insertMetric to stats DB
func insertValues(m metric) error {
	db, err := getConnection(POSTGRESQL_STATS_DB)
	if err != nil {
		return err
	}

	_, err = db.Exec(fmt.Sprintf(
		"INSERT into %s (service,metric_name,metric_time,metric_value) "+
			"VALUES ($1, $2, $3, $4)",
		POSTGRESQL_STATS_TABLE,
	), m.Service, m.MetricName, m.MetricTime, m.MetricValue)
	return err
}

func runIntQuery(schema string, q string) (int, error) {
	var res int
	db, err := getConnection(schema)
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
