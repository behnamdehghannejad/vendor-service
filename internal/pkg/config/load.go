package config

import (
	"log"
	"os"

	"github.com/behnamdehghannejad/vendorservice/internal/pkg/apperror"
	"github.com/spf13/viper"
)

func Load() (Config, error) {
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
		return Config{}, apperror.Wrap(err).UnExpected().Log().Build()
	}

	return cfg, nil
}
