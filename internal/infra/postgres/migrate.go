package postgres

import (
	"database/sql"
	"fmt"
	"os"
	"path/filepath"

	"github.com/behnamdehghannejad/vendorservice/internal/pkg/apperror"

	_ "github.com/lib/pq"

	migrate "github.com/rubenv/sql-migrate"
)

type Migrator struct {
	dialect    string
	host       string
	port       string
	database   string
	password   string
	username   string
	migrations *migrate.FileMigrationSource
}

func NewMigrator(cfg PostgresConfig) (*Migrator, error) {
	root, err := filepath.Abs(os.Getenv("CONFIG_PATH"))
	if err != nil {
		return nil, apperror.Wrap(err).
			Warning().
			Forbidden().
			Build()
	}

	migrationPath := filepath.Join(root, cfg.MigrationPath)
	migrations := &migrate.FileMigrationSource{
		Dir: migrationPath,
	}

	return &Migrator{
		host:       cfg.Host,
		port:       cfg.Port,
		database:   cfg.Database,
		username:   cfg.Username,
		password:   cfg.Password,
		dialect:    "postgres",
		migrations: migrations,
	}, nil
}

func (m *Migrator) UP() error {
	dsn := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		m.host,
		m.port,
		m.username,
		m.password,
		m.database,
	)

	db, err := sql.Open("postgres", dsn)
	if err != nil {
		return apperror.Wrap(err).
			UnExpected().
			Input(dsn).
			Warning().
			Log().
			Build()
	}

	defer db.Close()

	if err := db.Ping(); err != nil {
		return apperror.Wrap(err).
			UnExpected().
			Input(dsn).
			Warning().
			Log().
			Build()
	}

	_, err = migrate.Exec(db, m.dialect, m.migrations, migrate.Up)
	if err != nil {
		return apperror.Wrap(err).
			UnExpected().
			Input(dsn).
			Warning().
			Log().
			Build()
	}

	return nil
}
