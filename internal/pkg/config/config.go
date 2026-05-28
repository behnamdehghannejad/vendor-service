package config

import (
	"github.com/behnamdehghannejad/vendorservice/internal/handler/httphandler"
	"github.com/behnamdehghannejad/vendorservice/internal/infra/postgres"
)

type Config struct {
	App      httphandler.HttpConfig
	Database postgres.PostgresConfig
}
