package postgres

import (
	"testing"

	"github.com/dhui/dktest"
	"github.com/stretchr/testify/require"

	"private-conda-repo/store/models"
)

func TestStore_PackageCountOperations(t *testing.T) {
	assert := require.New(t)
	channel := "daniel-counts"
	packageName := "perfana"
	platform := "noarch"

	dktest.Run(t, imageName, postgresImageOptions, func(t *testing.T, info dktest.ContainerInfo) {
		store, err := newTestDb()
		assert.NoError(err)

		counts, err := store.GetPackageCounts(channel, packageName)
		assert.NoError(err)
		assert.Len(counts, 0)

		pkg, err := store.CreatePackageCount(&models.PackageCount{
			Channel:     channel,
			Package:     packageName,
			BuildString: "py",
			BuildNumber: 0,
			Version:     "0.0.1",
			Platform:    platform,
		})
		assert.NoError(err)

		count, err := store.IncreasePackageCount(pkg)
		assert.NoError(err)
		assert.EqualValues(1, count.Count)

		counts, err = store.GetPackageCounts(channel, packageName)
		assert.NoError(err)
		assert.Len(counts, 1)

		err = store.RemovePackageCount(pkg)
		assert.NoError(err)
	})
}
