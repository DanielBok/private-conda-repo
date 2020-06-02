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
	assert := require.New(t)

	_, file, _, _ := runtime.Caller(0)
	file = filepath.Join(filepath.Dir(file), "repodata.json")

	data, hasChanges, err := index.FixRepoData(file, []string{"abi"})
	assert.NoError(err)
	assert.True(hasChanges)

	for _, p := range data.Packages {
		for _, d := range p.Depends {
			// check that abi is removed
			assert.False(strings.HasPrefix(strings.ToLower(d), "python_abi"))
		}
	}
}

func TestFixRepoData_NoFixes(t *testing.T) {
	assert := require.New(t)

	_, file, _, _ := runtime.Caller(0)
	file = filepath.Join(filepath.Dir(file), "repodata.json")

	data, hasChanges, err := index.FixRepoData(file, nil)
	assert.NoError(err)
	assert.False(hasChanges)

	atLeastOnePackageHasPythonAbi := false
	for _, p := range data.Packages {
		for _, d := range p.Depends {
			// check that at least one package hsa python abi

			if strings.HasPrefix(strings.ToLower(d), "python_abi") {
				atLeastOnePackageHasPythonAbi = true
				break
			}
		}
	}
	assert.True(atLeastOnePackageHasPythonAbi)
}
