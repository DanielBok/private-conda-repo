package indexer_test

import (
	"testing"

	"github.com/stretchr/testify/require"

	. "private-conda-repo/indexer"
)

// This is mostly a smoke test. Because if the image already exists
// the pullLatestImage method will not be executed
func TestIndexImage_UpdateImage(t *testing.T) {
	assert := require.New(t)

	mgr, err := NewDockerIndexer("danielbok/conda-repo-mgr")
	assert.NoError(err)
	err = mgr.Update()
	assert.NoError(err)
}

func TestManager_CheckDockerVersion(t *testing.T) {
	assert := require.New(t)

	mgr, err := NewDockerIndexer("danielbok/conda-repo-mgr")
	assert.NoError(err)
	err = mgr.Check()
	assert.NoError(err)
}
