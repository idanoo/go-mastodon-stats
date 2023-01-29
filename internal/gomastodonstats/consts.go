package gomastodonstats

// Host information
var (
	POSTGRESQL_HOST        string
	POSTGRESQL_USER        string
	POSTGRESQL_PASS        string
	POSTGRESQL_STATS_DB    string
	POSTGRESQL_STATS_TABLE = "statsdb"

	TIMEZONE string

	// Pixelfed
	PIXELFED_DB_SCHEMA  string
	PIXELFED_USER_QUERY = "SELECT count(*) FROM users WHERE status IS NULL;"

	// Matrix
	MATRIX_DB_SCHEMA  string
	MATRIX_USER_QUERY = "SELECT count(*) FROM users WHERE deactivated = 0;"

	// Mastodon
	MASTODON_DB_SCHEMA  string
	MASTODON_USER_QUERY = "SELECT count(*) FROM users WHERE disabled = False;"

	// Mobilizon
	MOBILIZON_DB_SCHEMA  string
	MOBILIZON_USER_QUERY = "SELECT count(*) FROM users WHERE disabled = False;"

	// Peertube
	PEERTUBE_DB_SCHEMA  string
	PEERTUBE_USER_QUERY = "SELECT count(*) FROM \"user\" WHERE blocked = False;"
)
