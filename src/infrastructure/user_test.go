package infrastructure

import (
	"testing"

	"github.com/naoto67/clean-architecture/src/config"
	"github.com/naoto67/clean-architecture/src/domain/model"
	"github.com/naoto67/clean-architecture/src/domain/repository"
	"github.com/naoto67/clean-architecture/src/infrastructure/database"
	"github.com/naoto67/clean-architecture/test"

	"github.com/stretchr/testify/assert"
)

var (
	userRepo repository.UserRepository
	db       *database.DB
)

// before all and after all
func TestMain(m *testing.M) {
	config.Setup()
	setupUserRepository()
	// seed
	userRepo.Save(model.User{Name: "username"})
	m.Run()
	test.TruncateUsersTable()
}

func TestFindByName(t *testing.T) {
	assert := assert.New(t)

	t.Run("succeeded to find by username is username", func(t *testing.T) {
		username := "username"
		user, err := userRepo.FindByName(username)
		assert.Nil(err)
		assert.Equal(user.Name, username)
	})

	t.Run("failed to find by username is test", func(t *testing.T) {
		user, err := userRepo.FindByName("test")
		assert.NotNil(err)
		assert.Nil(user)
	})
}

func TestSave(t *testing.T) {
	assert := assert.New(t)

	t.Run("succeeded to save user when username is test", func(t *testing.T) {
		username := "test"
		err := userRepo.Save(model.User{Name: username})
		assert.Nil(err)
	})

	t.Run("failed to save user when username is username because duplicated name", func(t *testing.T) {
		username := "username"
		err := userRepo.Save(model.User{Name: username})
		assert.NotNil(err)
	})
}

func TestFindAll(t *testing.T) {
	assert := assert.New(t)

	users, err := userRepo.FindAll()
	assert.NotNil(users)
	assert.Nil(err)
}

func setupUserRepository() {
	if db == nil {
		db = database.New()
	}
	if userRepo == nil {
		userRepo = NewUserRepository(db)
	}
}
