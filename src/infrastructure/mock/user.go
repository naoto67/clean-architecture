package mock

import (
	"fmt"

	"github.com/naoto67/clean-architecture/src/domain/model"
	"github.com/naoto67/clean-architecture/src/domain/repository"
)

type userRepository struct {
	DB *MemoryMock
}

func NewUserRepository() repository.UserRepository {
	return userRepository{
		DB: NewMemoryMock(),
	}
}

var users = []model.User{
	model.User{ID: 1, Name: "sample"},
	model.User{ID: 2, Name: "sample2"},
	model.User{ID: 3, Name: "sap"},
}

func (repository userRepository) FindByName(name string) (*model.User, error) {
	for _, u := range users {
		if u.Name == name {
			return &u, nil
		}
	}
	return nil, fmt.Errorf("Name: %s not found", name)
}

func (repository userRepository) Create(user model.User) error {
	user.ID = users[len(users)-1].ID + 1
	users = append(users, user)
	return nil
}

func (repository userRepository) Update(user model.User) error {
	for _, u := range users {
		if u.ID == user.ID {
			u.Name = user.Name
		}
	}
	return nil
}

func (repository userRepository) FindAll() ([]model.User, error) {
	return users, nil
}
