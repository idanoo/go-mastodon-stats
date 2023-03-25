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

	MASTODON_CLIENT_NAME     = "go-mastodon-stats"
	MASTODON_INSTANCE_URL    string
	MASTODON_CLIENT_ID       string
	MASTODON_CLIENT_SECRET   string
	MASTODON_CLIENT_USERNAME string
	MASTODON_CLIENT_PASSWORD string

	// UserCount metric name
	METRICNAME_USERCOUNT    = "userCount"
	METRICNAME_1W_USERCOUNT = "userCount1W"

	// This is hardcoded because.. well configs are annoying
	SERVICE_LINKS = map[string]string{
		PIXELFED_IDENTIFIER:  "https://pixelfed.nz",
		MATRIX_IDENTIFIDER:   "https://mtrx.nz",
		MASTODON_IDENTIFIER:  "https://mastodon.nz",
		MOBILIZON_IDENTIFIER: "https://openevents.nz",
		PEERTUBE_IDENTIFIER:  "https://peertube.nz",
		BOOKWYRM_IDENTIFIER:  "https://bookworm.nz",
	}

	// Pixelfed
	PIXELFED_DB_SCHEMA  string
	PIXELFED_USER_QUERY = "SELECT count(*) FROM users WHERE status IS NULL;"
	PIXELFED_IDENTIFIER = "pixelfed"

	// Matrix
	MATRIX_DB_SCHEMA   string
	MATRIX_USER_QUERY  = "SELECT count(*) FROM users WHERE deactivated = 0;"
	MATRIX_IDENTIFIDER = "matrix"

	// Mastodon
	MASTODON_DB_SCHEMA            string
	MASTODON_USER_QUERY           = "SELECT count(*) FROM users WHERE disabled = False AND confirmed_at IS NOT NULL AND approved = True;"
	MASTODON_1W_ACTIVE_USER_QUERY = "SELECT count(*) FROM users WHERE disabled = False AND confirmed_at IS NOT NULL AND approved = True AND current_sign_in_at >=  now() - interval '1 week';"
	MASTODON_IDENTIFIER           = "mastodon"

	// Mobilizon
	MOBILIZON_DB_SCHEMA  string
	MOBILIZON_USER_QUERY = "SELECT count(*) FROM users WHERE disabled = False;"
	MOBILIZON_IDENTIFIER = "mobilizon"

	// Peertube
	PEERTUBE_DB_SCHEMA  string
	PEERTUBE_USER_QUERY = "SELECT count(*) FROM \"user\" WHERE blocked = False;"
	PEERTUBE_IDENTIFIER = "peertube"

	// BookWyrm
	BOOKWYRM_DB_SCHEMA  string
	BOOKWYRM_USER_QUERY = "SELECT count(*) FROM bookwyrm_user WHERE local = True AND is_active = True;"
	BOOKWYRM_IDENTIFIER = "bookwyrm"
)
