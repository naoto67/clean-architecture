// +build local

package config

import (
	"os"

	"github.com/joho/godotenv"
)

func setup() error {
	env := os.Getenv("ENV")
	return godotenv.Load("./env/" + env + ".env")
}
