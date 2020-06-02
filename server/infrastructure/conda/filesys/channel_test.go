package filesys_test

import (
	"os"
	"testing"

	"github.com/stretchr/testify/require"

	"private-conda-repo/api/dto"
	"private-conda-repo/libs"
	"private-conda-repo/testutils"
)

func TestChannel_CRUDPackage(t *testing.T) {
	t.Parallel()

	var assert = require.New(t)
	repo := NewFileSys()

	chn, err := repo.CreateChannel("crud-channel-packages")
	assert.NoError(err)

	testPkg := testutils.GetTestPackages()["perfana-0.0.6-py_0.tar.bz2"]

	file, err := os.Open(testPkg.Path)
	assert.NoError(err)
	defer libs.IOCloser(file)

	pkg, err := chn.AddPackage(file, testPkg.ToPackageDto(), nil)
	assert.NoError(err)

	channelData, err := chn.GetChannelData()
	assert.NoError(err)

	assert.Len(channelData.Packages, 1)
	assert.NotNil(channelData.Packages["perfana"])

	err = chn.RemoveSinglePackage(pkg)
	assert.NoError(err)
}

func TestChannel_GetChannelData(t *testing.T) {
	t.Parallel()

	var assert = require.New(t)
	chn, err := newPreloadedChannel("get-channel-data")
	assert.NoError(err)

	// both packages (copulae and perfana) are registered
	channelData, err := chn.GetChannelData()
	assert.NoError(err)
	assert.Len(channelData.Packages, 2)
	assert.EqualValues("0.4.3", *channelData.Packages["copulae"].Version)
	assert.EqualValues("0.0.6", *channelData.Packages["perfana"].Version)

	// Remove package updates indices correctly
	err = chn.RemoveSinglePackage(&dto.PackageDto{
		Name:        "perfana",
		Version:     "0.0.6",
		BuildString: "py",
		Platform:    "noarch",
	})

	assert.NoError(err)
	channelData, err = chn.GetChannelData()
	assert.NoError(err)
	assert.EqualValues("0.0.5", *channelData.Packages["perfana"].Version)
}

func TestChannel_RemovePackageAllVersions(t *testing.T) {
	t.Parallel()

	var assert = require.New(t)
	chn, err := newPreloadedChannel("remove-package-all-versions-channel")
	assert.NoError(err)

	n, err := chn.RemovePackageAllVersions("copulae")
	assert.NoError(err)
	assert.EqualValues(6, n)

	channelData, err := chn.GetChannelData()
	assert.NoError(err)
	assert.Len(channelData.Packages, 1)
}
