package memory

import (
	"fmt"

	"github.com/naoto67/clean-architecture/src/domain/model"
)

var users = []model.User{}

func (m Memory) FindByName(name string) (*model.User, error) {
	for _, u := range users {
		if u.Name == name {
			return &u, nil
		}
	}
	return nil, fmt.Errorf("Name: %s not found", name)
}

func (m Memory) FindById(id int) (*model.User, error) {
	for _, u := range users {
		if u.ID == id {
			return &u, nil
		}
	}
	return nil, fmt.Errorf("ID: %d not found", id)
}

func (m Memory) Create(user model.User) error {
	user.ID = users[len(users)-1].ID + 1
	users = append(users, user)
	return nil
}

func (m Memory) Update(user model.User) error {
	for _, u := range users {
		if u.ID == user.ID {
			u.Name = user.Name
		}
	}
	return nil
}

func (m Memory) FindAll() ([]model.User, error) {
	return users, nil
}
