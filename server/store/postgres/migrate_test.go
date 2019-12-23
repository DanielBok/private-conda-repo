package postgres

import (
	"testing"

	"github.com/dhui/dktest"
	_ "github.com/lib/pq"
	"github.com/stretchr/testify/assert"
)

func TestStore_Migrate(t *testing.T) {
	dktest.Run(t, imageName, postgresImageOptions, func(t *testing.T, info dktest.ContainerInfo) {
		db, err := newTestDb()
		assert.NoError(t, err)

		for i := 0; i < 3; i++ {
			err = db.Migrate()
			assert.NoError(t, err, "migration should not raise errors even when migrating database which is at latest revision")
		}
	})
}
