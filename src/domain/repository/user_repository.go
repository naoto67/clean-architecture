package repository

import (
	"github.com/naoto67/clean-architecture/src/domain/model"
)

type UserRepository interface {
	FindByName(name string) (*model.User, error)
	FindAll() ([]model.User, error)
	Save(user model.User) error
}
