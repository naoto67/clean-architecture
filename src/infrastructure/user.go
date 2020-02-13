package infrastructure

import (
	"github.com/naoto67/clean-architecture/src/domain/model"
	"github.com/naoto67/clean-architecture/src/infrastructure/mock"
)

type UserRepository struct {
	DB *mock.Mock
}

func (repository UserRepository) FindByName(name string) (*model.User, error) {
	return repository.DB.FindByName(name)
}

func (repository UserRepository) Save(user model.User) error {
	return repository.DB.Save(user)
}

func (repository UserRepository) FindAll() ([]model.User, error) {
	return repository.DB.FindAll()
}
