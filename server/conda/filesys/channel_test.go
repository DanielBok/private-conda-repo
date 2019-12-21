package filesys

import (
	"os"
	"path/filepath"
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

	file, err := os.Open(perfana)
	assert.NoError(err)
	defer func() { _ = file.Close() }()

	platform := string(condatypes.NOARCH)
	pkg, err := chn.AddPackage(file, platform, "perfana")
	assert.Error(err)

	pkg, err = chn.AddPackage(file, platform, filepath.Base(perfana))
	assert.NoError(err)

	meta, err := chn.GetMetaInfo()
	assert.NoError(err)

	assert.Len(meta.Packages, 1)
	assert.NotNil(meta.Packages["perfana"])

	err = chn.RemoveSinglePackage(pkg)
	assert.NoError(err)
}
