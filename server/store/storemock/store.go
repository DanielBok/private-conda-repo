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
	return &MockStore{}, nil
}
