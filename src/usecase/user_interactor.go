package usecase

import (
	"github.com/naoto67/clean-architecture/src/domain/model"
	"github.com/naoto67/clean-architecture/src/domain/repository"
)

type UserInteractor interface {
	Add(user model.User) error
}

type userInteractor struct {
	userRepository repository.UserRepository
}

func NewUserInteractor(repository repository.UserRepository) UserInteractor {
	return &userInteractor{
		userRepository: repository,
	}
}

func (interactor *userInteractor) Add(user model.User) error {
	return interactor.userRepository.Create(user)
}
