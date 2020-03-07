package registry

import (
	"github.com/naoto67/clean-architecture/src/domain/repository"
	"github.com/naoto67/clean-architecture/src/infrastructure"
	"github.com/naoto67/clean-architecture/src/infrastructure/database"
	"github.com/naoto67/clean-architecture/src/infrastructure/memory"
)

type Registry interface {
	NewUserRepository() repository.UserRepository
}

type registry struct {
	memory *memory.Memory
	db     *database.DB
}

func init() {
	New()
}

func New() Registry {
	return &registry{
		memory: memory.New(),
		db:     database.New(),
	}
}

func (reg *registry) NewUserRepository() repository.UserRepository {
	return infrastructure.NewUserRepository(reg.db)
}
