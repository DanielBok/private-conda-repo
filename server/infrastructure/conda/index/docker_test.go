package index_test

import (
	"testing"

	"github.com/stretchr/testify/require"

	. "private-conda-repo/infrastructure/conda/index"
)

// This is mostly a smoke test. Because if the image already exists
// the pullLatestImage method will not be executed
func TestIndexImage_UpdateImage(t *testing.T) {
	assert := require.New(t)

	mgr, err := NewDockerIndex("danielbok/conda-repo-mgr")
	assert.NoError(err)
	err = mgr.Update()
	assert.NoError(err)
}

func TestManager_CheckDockerVersion(t *testing.T) {
	assert := require.New(t)

	mgr, err := NewDockerIndex("danielbok/conda-repo-mgr")
	assert.NoError(err)
	assert.IsType(&DockerIndex{}, mgr)
}
