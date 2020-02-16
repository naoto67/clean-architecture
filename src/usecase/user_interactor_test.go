package usecase

import (
	"testing"

	"github.com/naoto67/clean-architecture/src/domain/model"
	"github.com/naoto67/clean-architecture/src/infrastructure/mock"
	"github.com/stretchr/testify/assert"
)

func TestAdd(t *testing.T) {
	assert := assert.New(t)

	interactor := NewUserInteractor(mock.NewUserRepository())

	err := interactor.Add(model.User{Name: "test"})
	assert.Nil(err)
}
