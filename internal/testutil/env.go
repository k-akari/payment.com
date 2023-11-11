package testutil

import (
	"sync"

	"github.com/caarlos0/env/v6"
)

var (
	loadOnce sync.Once
	cfg      testEnv
)

type testEnv struct {
	DBHost    string `env:"DB_HOST" envDefault:"db"`
	DBPort    uint16 `env:"DB_PORT" envDefault:"3306"`
	DBUser    string `env:"DB_USER" envDefault:"root"`
	DBPass    string `env:"DB_PASS" envDefault:"password"`
	RedisHost string `env:"REDIS_HOST" envDefault:"redis"`
	RedisPort int    `env:"REDIS_PORT" envDefault:"6379"`
}

func loadEnv() *testEnv {
	loadOnce.Do(func() {
		if err := env.Parse(&cfg); err != nil {
			panic(err)
		}
	})

	return &cfg
}
