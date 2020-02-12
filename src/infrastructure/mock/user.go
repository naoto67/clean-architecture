package mock

import (
	"fmt"

	"github.com/naoto67/clean-architecture/src/domain/model"
)

var users = []model.User{
	model.User{ID: 1, Name: "sample"},
	model.User{ID: 2, Name: "sample2"},
	model.User{ID: 3, Name: "sap"},
}

func (m *Mock) FindByName(name string) (*model.User, error) {
	for _, u := range users {
		if u.Name == name {
			return &u, nil
		}
	}
	return nil, fmt.Errorf("Name: %s not found", name)
}
