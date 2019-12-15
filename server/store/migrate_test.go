package store

import (
	"testing"

	"github.com/dhui/dktest"
	_ "github.com/lib/pq"
	"github.com/stretchr/testify/assert"
)

func TestStore_Migrate(t *testing.T) {
	dktest.Run(t, imageName, postgresImageOptions, func(t *testing.T, info dktest.ContainerInfo) {
		_, err := newTestDb()
		assert.NoError(t, err)
	})
}

func newTestDb() (*Store, error) {
	store, err := New()
	if err != nil {
		return nil, err
	}

	err = store.Migrate()
	if err != nil {
		return nil, err
	}
	return store, nil
}
