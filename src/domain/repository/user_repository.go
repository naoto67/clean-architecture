package repository

import (
	"github.com/naoto67/clean-architecture/src/domain/model"
)

type UserRepository interface {
	FindByName(name string) (*model.User, error)
	FindAll() ([]model.User, error)
	Create(user model.User) error
	Update(user model.User) error
}
