package postgres_test

import (
	"testing"

	"github.com/dhui/dktest"
	_ "github.com/lib/pq"
	"github.com/stretchr/testify/require"
)

func TestPostgres_Migrate(t *testing.T) {
	dktest.Run(t, imageName, postgresImageOptions, func(t *testing.T, info dktest.ContainerInfo) {
		db, err := newTestDb(info)
		require.NoError(t, err)

		for i := 0; i < 3; i++ {
			err = db.Migrate()
			require.NoError(t, err, "migration should not raise errors even when migrating database which is at latest revision")
		}
	})
}
