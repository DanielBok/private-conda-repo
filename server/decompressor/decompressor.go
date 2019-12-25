package decompressor

import (
	"io"

	"private-conda-repo/config"
)

type Decompressor interface {
	RetrieveMetadata(file io.ReadCloser) (*Package, error)
}

func New() (Decompressor, error) {
	conf, err := config.New()
	if err != nil {
		return nil, err
	}

	switch conf.Decompressor.Type {
	case "mock":
		return &mockDecompressor{}, nil
	default:
		return &tarBz2Decompressor{}, nil
	}
}
