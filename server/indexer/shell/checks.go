package shell

import (
	"os/exec"

	"github.com/pkg/errors"
	log "github.com/sirupsen/logrus"
)

func (m *Manager) Check() error {
	if version, err := exec.Command("conda", "--version").Output(); err != nil {
		return errors.Wrapf(err, "conda not installed")
	} else {
		log.Printf("Using: %s", version)
	}

	return nil
}
