package database

import (
	"fmt"

	"github.com/naoto67/clean-architecture/src/config"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

type DB struct {
	conn *sqlx.DB
}

func New() *DB {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true",
		config.Config.DatabaseUser,
		config.Config.DatabasePassword,
		config.Config.DatabaseHost,
		config.Config.DatabasePort,
		config.Config.DatabaseName,
	)
	db, err := sqlx.Open(
		"mysql",
		dsn,
	)
	if err != nil {
		panic(err)
	}
	if err := db.Ping(); err != nil {
		panic(err)
	}

	return &DB{conn: db}
}
