package decompressor

import (
	"io"
	"strings"
)

type Decompressor interface {
	RetrieveMetadata(file io.ReadCloser) (*Package, error)
}

func New(decompressor string) Decompressor {
	switch strings.ToLower(strings.TrimSpace(decompressor)) {
	case "mock":
		return &mockDecompressor{}
	default:
		return &tarBz2Decompressor{}
	}
}
