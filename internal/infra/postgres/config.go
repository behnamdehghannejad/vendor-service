package postgres

type PostgresConfig struct {
	Username string
	Password string
	Database string
	Host     string
	Port     string
	SSLMode  string
	Migrate  bool
}
