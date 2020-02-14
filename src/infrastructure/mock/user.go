package mock

import (
	"fmt"

	"github.com/naoto67/clean-architecture/src/domain/model"
	"github.com/naoto67/clean-architecture/src/domain/repository"
)

type UserRepository struct {
	DB *MemoryMock
}

func NewUserRepository() repository.UserRepository {
	return UserRepository{
		DB: NewMemoryMock(),
	}
}

var users = []model.User{
	model.User{ID: 1, Name: "sample"},
	model.User{ID: 2, Name: "sample2"},
	model.User{ID: 3, Name: "sap"},
}

func (repository UserRepository) FindByName(name string) (*model.User, error) {
	for _, u := range users {
		if u.Name == name {
			return &u, nil
		}
	}
	return nil, fmt.Errorf("Name: %s not found", name)
}

func (repository UserRepository) Save(user model.User) error {
	user.ID = users[len(users)-1].ID + 1
	users = append(users, user)
	return nil
}

func (repository UserRepository) FindAll() ([]model.User, error) {
	return users, nil
}
