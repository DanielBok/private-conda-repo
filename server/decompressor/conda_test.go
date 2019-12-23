package decompressor

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"

	"private-conda-repo/testutils"
)

func TestTarBz2Decompressor_RetrieveMetadata(t *testing.T) {
	t.Parallel()
	dcp := tarBz2Decompressor{}

	test := func(details testutils.TestPackage) {
		f, err := os.Open(details.Path)
		assert.NoError(t, err)
		defer func() { _ = f.Close() }()

		pkg, err := dcp.RetrieveMetadata(f)
		assert.NoError(t, err)
		assert.Equal(t, details.Filename, pkg.Package.Filename())
		assert.Equal(t, details.Platform, pkg.Package.Platform)
		pkg.Close()
	}

	for _, details := range testutils.GetTestPackages() {
		test(details)
	}
}
