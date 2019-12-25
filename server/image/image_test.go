package image

import (
	"testing"

	"github.com/stretchr/testify/require"
)

// This is mostly a smoke test. Because if the image already exists
// the pullLatestImage method will not be executed
func TestIndexImage_UpdateImage(t *testing.T) {
	assert := require.New(t)

	mgr, err := New()
	assert.NoError(err)
	err = mgr.UpdateImage()
	assert.NoError(err)
}
