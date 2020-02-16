package infrastructure

import (
	"github.com/naoto67/clean-architecture/src/domain/model"
	"github.com/naoto67/clean-architecture/src/domain/repository"
	"github.com/naoto67/clean-architecture/src/infrastructure/memory"
)

type userRepository struct {
	Memory *memory.Memory
}

func NewUserRepository(m *memory.Memory) repository.UserRepository {
	return userRepository{Memory: m}
}

func (repository userRepository) FindByName(name string) (*model.User, error) {
	return repository.Memory.FindByName(name)
}

func (repository userRepository) Save(user model.User) error {
	return repository.Memory.Save(user)
}

func (repository userRepository) FindAll() ([]model.User, error) {
	return repository.Memory.FindAll()
}
