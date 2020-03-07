// +build integration

package config

import (
	"os"
)

var (
	environments map[string]string = map[string]string{
		"DATABASE_USER":     "root",
		"DATABASE_PASSWORD": "",
		"DATABASE_PORT":     "3306",
		"DATABASE_HOST":     "localhost",
		"DATABASE_NAME":     "ca_test",
	}
)

func setup() error {
	for key, value := range environments {
		os.Setenv(key, value)
	}
	return nil
}
