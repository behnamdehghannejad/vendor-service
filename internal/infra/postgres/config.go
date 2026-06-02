package postgres

type PostgresConfig struct {
	MigrationPath string `mapstructure:"migration_path"`
	Username      string `mapstructure:"username"`
	Password      string `mapstructure:"password"`
	Database      string `mapstructure:"database"`
	DatabaseTest  string `mapstructure:"database_test"`
	Host          string `mapstructure:"host"`
	Port          string `mapstructure:"port"`
	SSLMode       string `mapstructure:"sslmode"`
	Migrate       bool   `mapstructure:"migrate"`
}
