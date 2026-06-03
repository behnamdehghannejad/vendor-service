package config

import (
	discount "github.com/behnamdehghannejad/vendorservice/internal/adapter/discount_client"
	"github.com/behnamdehghannejad/vendorservice/internal/adapter/handler/httphandler"
	"github.com/behnamdehghannejad/vendorservice/internal/infra/postgres"
)

type Config struct {
	App            httphandler.HttpConfig        `mapstructure:"app"`
	Database       postgres.PostgresConfig       `mapstructure:"database"`
	DiscountClient discount.DiscountClientConfig `mapstructure:"discount_client"`
}
