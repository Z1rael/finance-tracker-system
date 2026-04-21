package config

import (
	"log"

	"github.com/caarlos0/env/v11"
)

type Config struct {
	DB_address  string `env:"DB_ADDRESS" envDefault:"localhost"`
	DB_user     string `env:"DB_USER" envDefault: "user"`
	DB_password string `env:"DB_PASSWORD" envDefault:"secret"`
	DB_name     string `env:"DB_NAME" envDefault:"finance-tracker-db"`
	DB_port     string `env:"DB_PORT" envDefault:"5432"`
}

func NewConfig() *Config {
	cfg := Config{}
	if err := env.Parse(&cfg); err != nil {
		log.Fatalf("Failed to parse config: %v", err)
		return nil
	}

	return &cfg
}
