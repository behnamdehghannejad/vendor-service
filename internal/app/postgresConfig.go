package app

import (
	"fmt"
	"log"
	"time"

	_ "github.com/lib/pq"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func (cfg *Config) PostgresDSN() string {
	return fmt.Sprintf(
		"host=%s port=%d user=%s password=%s dbname=%s sslmode=%s",
		cfg.Database.Host,
		cfg.Database.Port,
		cfg.Database.User,
		cfg.Database.Password,
		cfg.Database.Name,
		cfg.Database.SSLMode,
	)
}

func NewPostgres(cfg *Config) (*gorm.DB, error) {
	dsn := cfg.PostgresDSN()

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		PrepareStmt:                              true,
		SkipDefaultTransaction:                   false,
		DisableForeignKeyConstraintWhenMigrating: true,
		Logger:                                   logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		return nil, err
	}

	sqlDB, err := db.DB()
	if err != nil {
		return nil, err
	}

	sqlDB.SetMaxIdleConns(10)
	sqlDB.SetMaxOpenConns(100)
	sqlDB.SetConnMaxIdleTime(5 * time.Minute)
	sqlDB.SetConnMaxLifetime(30 * time.Minute)

	err = sqlDB.Ping()
	if err != nil {
		log.Fatalf("failed to connect repository with ip: %s and port: %s -> ERROR: %v", cfg.Database.Host, cfg.Database.Port, err)
		return nil, err
	}

	return db, nil
}
