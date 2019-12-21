package filesys

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"

	"private-conda-repo/conda/condatypes"
)

func TestConda_CRUDPackage(t *testing.T) {
	t.Parallel()

	var assert = assert.New(t)
	repo, cleanup := newTestConda()
	defer cleanup()

	chn, err := repo.CreateChannel("test-channel")
	assert.NoError(err)

	testPkg := testPackages["perfana-0.0.6-py_0.tar.bz2"]

	file, err := os.Open(testPkg.Path)
	assert.NoError(err)
	defer func() { _ = file.Close() }()

	platform := string(condatypes.NOARCH)
	pkg, err := chn.AddPackage(file, platform, "perfana")
	assert.Error(err)

	pkg, err = chn.AddPackage(file, platform, testPkg.Filename)
	assert.NoError(err)

	meta, err := chn.GetMetaInfo()
	assert.NoError(err)

	assert.Len(meta.Packages, 1)
	assert.NotNil(meta.Packages["perfana"])

	err = chn.RemoveSinglePackage(pkg)
	assert.NoError(err)
}
