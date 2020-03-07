package test

import (
	"fmt"

	"github.com/naoto67/clean-architecture/src/config"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

var dbx *sqlx.DB

func fetchDBX() {
	if dbx == nil {
		dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s",
			config.Config.DatabaseUser,
			config.Config.DatabasePassword,
			config.Config.DatabaseHost,
			config.Config.DatabasePort,
			config.Config.DatabaseName,
		)

		dbx, _ = sqlx.Open(
			"mysql",
			dsn,
		)
	}
}

func TruncateUsersTable() {
	fetchDBX()
	_, err := dbx.Exec("TRUNCATE TABLE users")
	if err != nil {
		panic(err)
	}
}
