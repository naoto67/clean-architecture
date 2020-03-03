package config

import (
	"os"

	"github.com/caarlos0/env/v6"
	"github.com/joho/godotenv"
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
	err := loadEnv()
	if err != nil {
		panic(err)
	}

	if err := env.Parse(&Config); err != nil {
		panic(err)
	}
}

func loadEnv() error {
	env := os.Getenv("ENV")
	return godotenv.Load("./env/" + env + ".env")
}
