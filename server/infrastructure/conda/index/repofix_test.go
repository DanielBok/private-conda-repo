package index_test

import (
	"path/filepath"
	"runtime"
	"strings"
	"testing"

	"github.com/stretchr/testify/require"

	"private-conda-repo/infrastructure/conda/index"
)

func TestFixRepoData(t *testing.T) {
	t.Parallel()
	assert := require.New(t)

	_, file, _, _ := runtime.Caller(0)
	file = filepath.Join(filepath.Dir(file), "repodata.json")

	data, hasChanges, err := index.FixRepoData(file)
	assert.NoError(err)
	assert.True(hasChanges)

	for _, p := range data.Packages {
		for _, d := range p.Depends {
			// check that abi is removed
			assert.False(strings.HasPrefix(strings.ToLower(d), "python_abi"))
		}
	}
}
