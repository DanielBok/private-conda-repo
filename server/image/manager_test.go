package image

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestManager_CheckDockerVersion(t *testing.T) {
	assert := require.New(t)

	mgr, err := New()
	assert.NoError(err)
	version, err := mgr.CheckDockerVersion()
	assert.NotEmpty(version)
}
