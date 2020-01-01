package shell

import (
	"os/exec"
	"regexp"
	"strings"

	"github.com/pkg/errors"
)

const PKG = "conda-build"

func (m *Manager) Update() error {
	exists, err := m.condaBuildExist()
	if err != nil {
		return errors.Wrap(err, " could not install conda build")
	}

	if !exists {
		return m.installCondaBuild()
	}

	return nil
}

func (m *Manager) condaBuildExist() (bool, error) {
	cmd := []string{"list", "-f", PKG}
	output, err := exec.Command("conda", cmd...).CombinedOutput()
	if err != nil {
		return false, errors.Wrapf(err, "could not execute conda command. Is it installed?")
	}

	return regexp.MatchString("conda-build", strings.TrimSpace(string(output)))
}

func (m *Manager) installCondaBuild() error {
	cmd := []string{"install", "-y", PKG}
	if _, err := exec.Command("conda", cmd...).Output(); err != nil {
		return errors.Wrap(err, "could not install conda-build package")
	}
	return nil
}
