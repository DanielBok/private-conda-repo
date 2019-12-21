package storemock

import (
	"github.com/stretchr/testify/mock"

	"private-conda-repo/store"
)

type MockStore struct {
	mock.Mock
}

func init() {
	store.Register("mock", New)
}

func New() (store.Store, error) {
	store := &MockStore{}

	// Returns are matched by the actual method
	store.On("AddUser", mock.AnythingOfType("string"), mock.AnythingOfType("string"))
	store.On("RemoveUser", mock.AnythingOfType("string"), mock.AnythingOfType("string"))

	return store, nil
}
