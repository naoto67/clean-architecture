package infrastructure

import (
	"github.com/naoto67/clean-architecture/src/domain/model"
	"github.com/naoto67/clean-architecture/src/domain/repository"
	"github.com/naoto67/clean-architecture/src/infrastructure/memory"
)

type UserRepository struct {
	Memory *memory.Memory
}

func NewUserRepository(m *memory.Memory) repository.UserRepository {
	return UserRepository{Memory: m}
}

func (repository UserRepository) FindByName(name string) (*model.User, error) {
	return repository.Memory.FindByName(name)
}

func (repository UserRepository) Save(user model.User) error {
	return repository.Memory.Save(user)
}

func (repository UserRepository) FindAll() ([]model.User, error) {
	return repository.Memory.FindAll()
}
