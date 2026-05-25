package app

import (
	"log"
	"os"

	"github.com/spf13/viper"
)

type Config struct {
	App      AppConfig
	GRPC     GRPCConfig
	Database DatabaseConfig
	Redis    RedisConfig
}

type AppConfig struct {
	Name        string
	HttpPort    int
	GatewayPort int
}

type GRPCConfig struct {
	Host string
	Port int
}

type DatabaseConfig struct {
	Host     string
	Port     int
	User     string
	Password string
	Name     string
	SSLMode  string
}

type RedisConfig struct {
	Host     string
	Port     int
	Password string
	DB       int
}

func LoadConfig() *Config {
	env := os.Getenv("APP_ENV")

	if env == "" {
		env = "develop"
	}

	viper.SetConfigName(env)
	viper.SetConfigType("yaml")
	viper.AddConfigPath("./")

	viper.AutomaticEnv()

	err := viper.ReadInConfig()
	if err != nil {
		log.Fatalf("failed to read config: %v", err)
	}

	var cfg Config

	err = viper.Unmarshal(&cfg)
	if err != nil {
		log.Fatalf("failed to unmarshal config: %v", err)
	}

	return &cfg
}
