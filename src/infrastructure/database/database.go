package database

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

type DB struct {
	conn *sqlx.DB
}

func New() *DB {
	db, err := sqlx.Open(
		"mysql",
		"root@/ca_development",
	)
	if err != nil {
		panic(err)
	}
	if db.Ping() != nil {
		panic(err)
	}

	return &DB{conn: db}
}
