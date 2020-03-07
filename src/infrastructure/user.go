package infrastructure

import (
	"github.com/naoto67/clean-architecture/src/domain/model"
	"github.com/naoto67/clean-architecture/src/domain/repository"
	"github.com/naoto67/clean-architecture/src/infrastructure/database"
)

type userRepository struct {
	// Memory *memory.Memory
	db *database.DB
}

func NewUserRepository(db *database.DB) repository.UserRepository {
	return userRepository{db: db}
}

func (repository userRepository) FindByName(name string) (*model.User, error) {
	return repository.db.FindUserByName(name)
}

func (repository userRepository) Save(user model.User) error {
	return repository.db.SaveUser(user)
}

func (repository userRepository) FindAll() ([]model.User, error) {
	return repository.db.FindUsers()
}
