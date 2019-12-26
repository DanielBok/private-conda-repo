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
	s := &MockStore{}

	// Returns are matched by the actual method
	s.On("AddUser", mock.AnythingOfType("string"), mock.AnythingOfType("string"))
	s.On("RemoveUser", mock.AnythingOfType("string"), mock.AnythingOfType("string"))

	return s, nil
}
