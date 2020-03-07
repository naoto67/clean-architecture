package config

import (
	"github.com/caarlos0/env/v6"
)

type config struct {
	DatabaseUser     string `env:"DATABASE_USER"`
	DatabasePassword string `env:"DATABASE_PASSWORD"`
	DatabasePort     string `env:"DATABASE_PORT"`
	DatabaseHost     string `env:"DATABASE_HOST"`
	DatabaseName     string `env:"DATABASE_NAME"`
}

var Config config

func Setup() {
	err := setup()
	if err != nil {
		panic(err)
	}

	if err := env.Parse(&Config); err != nil {
		panic(err)
	}
}
