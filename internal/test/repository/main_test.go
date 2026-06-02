package repository

import (
	"fmt"
	"os"
	"testing"

	"github.com/behnamdehghannejad/vendorservice/internal/infra/postgres"
	"github.com/behnamdehghannejad/vendorservice/internal/pkg/config"
	"github.com/behnamdehghannejad/vendorservice/internal/pkg/log"
)

var (
	vendorRepo      *postgres.VendorRepository
	productRepo     *postgres.ProductRepository
	inventoryRepo   *postgres.InventoryRepository
	transactionRepo *postgres.TransactionRepository
)

func TestMain(m *testing.M) {
	os.Setenv("CONFIG_PATH", "../../..")

	err := log.Initialize()
	if err != nil {
		os.Exit(1)
	}

	cfg, err := config.Load()
	if err != nil {
		fmt.Printf("error load config %v\n", err)
		os.Exit(1)
	}

	postgresConfig := postgres.PostgresConfig{
		Username:      cfg.Database.Username,
		Password:      cfg.Database.Password,
		Database:      cfg.Database.DatabaseTest,
		Host:          cfg.Database.Host,
		Port:          cfg.Database.Port,
		Migrate:       cfg.Database.Migrate,
		SSLMode:       cfg.Database.SSLMode,
		MigrationPath: "../../.././" + cfg.Database.MigrationPath,
	}

	migrator := postgres.NewMigrator(postgresConfig)
	if err := migrator.UP(); err != nil {
		fmt.Printf("error to migrate database %v\n", err)
		os.Exit(1)
	}

	db, err := postgres.New(postgresConfig)
	if err != nil {
		fmt.Printf("error to connect database %v\n", err)
		os.Exit(1)
	}

	vendorRepo = postgres.NewVendorRepository(db)
	productRepo = postgres.NewProductRepository(db)
	transactionRepo = postgres.NewTransactionRepository(db)
	inventoryRepo = postgres.NewInventoryRepository(db)

	code := m.Run()
	os.Exit(code)
}
