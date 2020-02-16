package database

import (
	"github.com/naoto67/clean-architecture/src/domain/model"
)

func (db *DB) FindUserByName(name string) (*model.User, error) {
	var user model.User
	err := db.conn.Get(&user, "SELECT id, name FROM users WHERE name = ?", name)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (db *DB) FindUsers() ([]model.User, error) {
	var users []model.User
	err := db.conn.Select(&users, "SELECT id, name FROM users")
	if err != nil {
		return nil, err
	}
	return users, nil
}

func (db *DB) SaveUser(user model.User) error {
	sql := "INSERT INTO users (name) VALUES (?)"
	_, err := db.conn.Exec(sql, user.Name)
	return err
}
