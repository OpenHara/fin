package env

import (
	"github.com/caarlos0/env"
)

type Config struct {
	Port             int    `env:"PORT" envDefault:"3000"`
	DatabaseUrl      string `env:"DATABASE_URL" envDefault:"127.0.0.1"`
	DatabasePassword string `env:"DATABASE_PASSWORD,required"`
}

func Get() (*Config, error) {
	cfg := Config{}
	if err := env.Parse(&cfg); err != nil {
		return nil, err
	}

	return &cfg, nil
}
