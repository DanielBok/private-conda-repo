package shell

import (
	"os/exec"
	"path/filepath"

	"github.com/pkg/errors"

	"private-conda-repo/indexer"
)

type Manager struct {
}

func init() {
	indexer.Register("shell", func() (indexer.Indexer, error) {
		return &Manager{}, nil
	})
}

func (m *Manager) Index(dir string) error {
	cmd := []string{"index", dir}

	if _, err := exec.Command("conda", cmd...).Output(); err != nil {
		return errors.Wrapf(err, "could not index channel '%s'", filepath.Base(dir))
	}

	return nil
}
