package usecase

import (
	"testing"

	"github.com/naoto67/clean-architecture/src/config"
	"github.com/naoto67/clean-architecture/src/registry"
	"github.com/naoto67/clean-architecture/test"
)

var (
	userUsecase UserInteractor
	reg         registry.Registry
)

func TestMain(m *testing.M) {
	config.Setup()
	setupUserInteractor()
	m.Run()
	test.TruncateUsersTable()
}

func setupUserInteractor() {
	if reg == nil {
		reg = registry.New()
	}
	if userUsecase == nil {
		userUsecase = NewUserInteractor(reg.NewUserRepository())
	}
}
