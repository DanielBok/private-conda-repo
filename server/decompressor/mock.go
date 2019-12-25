package decompressor

import (
	"io"

	"github.com/stretchr/testify/mock"

	"private-conda-repo/conda/condatypes"
)

type mockDecompressor struct {
	mock.Mock
}

func (m mockDecompressor) RetrieveMetadata(file io.ReadCloser) (*Package, error) {
	m.Called(file)
	return &Package{
		Package: &condatypes.Package{
			Name:        "package",
			Version:     "0.0.1",
			BuildString: "py",
			BuildNumber: 0,
			Platform:    "noarch",
		},
		Filepath: "mock-path",
		file:     nil,
	}, nil
}
