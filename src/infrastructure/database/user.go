package database

import (
	"github.com/naoto67/clean-architecture/src/domain/model"
)

func (db *DB) FindUserByName(name string) (*model.User, error) {
	var user model.User
	err := db.conn.Get(&user, "SELECT * FROM users WHERE name = ?", name)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (db *DB) FindUserById(id int) (*model.User, error) {
	var user model.User
	err := db.conn.Get(&user, "SELECT * FROM users WHERE id = ?", id)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (db *DB) FindUsers() ([]model.User, error) {
	var users []model.User
	err := db.conn.Select(&users, "SELECT * FROM users")
	if err != nil {
		return nil, err
	}
	return users, nil
}

func (db *DB) InsertUser(user model.User) error {
	sql := "INSERT INTO users (name) VALUES (?)"
	_, err := db.conn.Exec(sql, user.Name)
	return err
}

func (db *DB) UpdateUser(user model.User) error {
	sql := "UPDATE users SET name = ? WHERE id = ?"
	_, err := db.conn.Exec(sql, user.Name, user.ID)
	return err
}
