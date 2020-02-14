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

func (m Memory) Save(user model.User) error {
	user.ID = users[len(users)-1].ID + 1
	users = append(users, user)
	return nil
}

func (m Memory) FindAll() ([]model.User, error) {
	return users, nil
}
