package decompressor

import (
	"os"
	"testing"

	"github.com/stretchr/testify/require"

	"private-conda-repo/testutils"
)

func TestTarBz2Decompressor_RetrieveMetadata(t *testing.T) {
	t.Parallel()
	assert := require.New(t)
	dcp := tarBz2Decompressor{}

	test := func(details testutils.TestPackage) {
		f, err := os.Open(details.Path)
		assert.NoError(err)
		defer func() { _ = f.Close() }()

		pkg, err := dcp.RetrieveMetadata(f)
		assert.NoError(err)
		assert.Equal(details.Filename, pkg.Package.Filename())
		assert.Equal(details.Platform, pkg.Package.Platform)
		pkg.Close()
	}

	for _, details := range testutils.GetTestPackages() {
		test(details)
	}
}
