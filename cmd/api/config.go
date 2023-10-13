package main

import (
	"github.com/caarlos0/env/v6"
)

type config struct {
	Port       int    `env:"PORT" envDefault:"8080"`
	DBHost     string `env:"DB_HOST" envDefault:"db"`
	DBPort     int    `env:"DB_PORT" envDefault:"3306"`
	DBUser     string `env:"DB_USER" envDefault:"root"`
	DBPassword string `env:"DB_PASSWORD" envDefault:"password"`
	DBName     string `env:"DB_NAME" envDefault:"payment"`
}

func newConfig() (*config, error) {
	cfg := &config{}
	if err := env.Parse(cfg); err != nil {
		return nil, err
	}
	return cfg, nil
}
