package decompressor

import (
	"os"

	"private-conda-repo/api/dto"
	"private-conda-repo/libs"
)

// Meta data that is derived when unzipping the tar.bz2 package
type MetaData struct {
	Package  *dto.PackageDto
	Filepath string
	file     *os.File
}

func (m *MetaData) Open() (*os.File, error) {
	var err error
	m.file, err = os.Open(m.Filepath)

	return m.file, err
}

func (m *MetaData) Close() {
	if m.file != nil {
		_ = m.file.Close()
	}

	if libs.PathExists(m.Filepath) {
		_ = os.Remove(m.Filepath)
	}
}
