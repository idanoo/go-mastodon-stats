package gomastodonstats

// Host information
var (
	POSTGRESQL_HOST        string
	POSTGRESQL_USER        string
	POSTGRESQL_PASS        string
	POSTGRESQL_STATS_DB    string
	POSTGRESQL_STATS_TABLE = "statsdb"

	TIMEZONE string

	MATRIX_WEBHOOK_URL     string
	MATRIX_WEBHOOK_API_KEY string
	MATRIX_WEBHOOK_CHANNEL string

	MASTODON_INSTANCE_URL string
	MASTODON_USERNAME     string
	MASTODON_PASSWORD     string

	// UserCount metric name
	METRICNAME_USERCOUNT = "userCount"

	// Pixelfed
	PIXELFED_DB_SCHEMA  string
	PIXELFED_USER_QUERY = "SELECT count(*) FROM users WHERE status IS NULL;"
	PIXELFED_IDENTIFIER = "pixelfed"

	// Matrix
	MATRIX_DB_SCHEMA   string
	MATRIX_USER_QUERY  = "SELECT count(*) FROM users WHERE deactivated = 0;"
	MATRIX_IDENTIFIDER = "matrix"

	// Mastodon
	MASTODON_DB_SCHEMA  string
	MASTODON_USER_QUERY = "SELECT count(*) FROM users WHERE disabled = False;"
	MASTODON_IDENTIFIER = "mastodon"

	// Mobilizon
	MOBILIZON_DB_SCHEMA  string
	MOBILIZON_USER_QUERY = "SELECT count(*) FROM users WHERE disabled = False;"
	MOBILIZON_IDENTIFIER = "mobilizon"

	// Peertube
	PEERTUBE_DB_SCHEMA  string
	PEERTUBE_USER_QUERY = "SELECT count(*) FROM \"user\" WHERE blocked = False;"
	PEERTUBE_IDENTIFIER = "peertube"
)
