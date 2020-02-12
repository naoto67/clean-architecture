package infrastructure

import (
	"github.com/naoto67/clean-architecture/src/domain/model"
	"github.com/naoto67/clean-architecture/src/domain/repository"
	"github.com/naoto67/clean-architecture/src/infrastructure/mock"
)

type userRepository struct {
	mock *mock.Mock
}

func NewUserRepository() repository.UserRepository {
	return userRepository{
		mock: mock.NewMock(),
	}
}

func (repository userRepository) FindByName(name string) (*model.User, error) {
	return repository.mock.FindByName(name)
}
