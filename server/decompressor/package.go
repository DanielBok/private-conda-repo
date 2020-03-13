package decompressor

import (
	"os"

	"private-conda-repo/conda/condatypes"
	"private-conda-repo/libs"
)

type Package struct {
	Package  *condatypes.Package
	Filepath string
	file     *os.File
}

func (p *Package) Open() (*os.File, error) {
	var err error
	p.file, err = os.Open(p.Filepath)

	return p.file, err
}

func (p *Package) Close() {
	if p.file != nil {
		_ = p.file.Close()
	}

	if libs.PathExists(p.Filepath) {
		_ = os.Remove(p.Filepath)
	}
}
