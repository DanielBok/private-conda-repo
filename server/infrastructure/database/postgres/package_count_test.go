package postgres_test

import (
	"testing"

	"github.com/dhui/dktest"
	"github.com/stretchr/testify/require"

	"private-conda-repo/domain/entity"
)

func TestPostgres_PackageCountOperations(t *testing.T) {
	assert := require.New(t)
	packageName := "perfana"
	platform := "noarch"

	dktest.Run(t, imageName, postgresImageOptions, func(t *testing.T, info dktest.ContainerInfo) {
		store, err := newTestDb(info)
		assert.NoError(err)

		chn, err := store.CreateChannel("counts-channel", "password", "salt")
		assert.NoError(err)

		counts, err := store.GetPackageCounts(chn.Id, packageName)
		assert.NoError(err)
		assert.Len(counts, 0)

		pkg, err := store.CreatePackageCount(&entity.PackageCount{
			ChannelId:   chn.Id,
			Package:     packageName,
			BuildString: "py",
			BuildNumber: 0,
			Version:     "0.0.1",
			Platform:    platform,
		})
		assert.NoError(err)
		assert.Equal(0, pkg.Count)

		pkg, err = store.IncreasePackageCount(pkg)
		assert.NoError(err)
		assert.EqualValues(1, pkg.Count)

		counts, err = store.GetPackageCounts(chn.Id, packageName)
		assert.NoError(err)
		assert.Len(counts, 1)

		err = store.RemovePackageCount(pkg)
		assert.NoError(err)
	})
}
