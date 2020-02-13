package infrastructure

import (
	"testing"

	"github.com/naoto67/clean-architecture/src/domain/model"
	"github.com/naoto67/clean-architecture/src/infrastructure/mock"
	"github.com/stretchr/testify/assert"
)

func TestFindByName(t *testing.T) {
	assert := assert.New(t)

	repo := UserRepository{
		DB: mock.New(),
	}

	testUserName := "test user"

	repo.Save(model.User{Name: testUserName})
	user, err := repo.FindByName(testUserName)
	assert.Nil(err)
	assert.Equal(user.Name, testUserName)
}

func TestSave(t *testing.T) {
	assert := assert.New(t)

	repo := UserRepository{
		DB: mock.New(),
	}
	user := model.User{Name: "test user"}
	err := repo.Save(user)
	assert.Nil(err)
}

func TestFindAll(t *testing.T) {
	assert := assert.New(t)

	repo := UserRepository{
		DB: mock.New(),
	}

	repo.Save(model.User{Name: "test user"})

	users, err := repo.FindAll()
	assert.NotNil(users)
	assert.Nil(err)
}
