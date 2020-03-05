package indexer

import (
	"os/exec"
	"path/filepath"
	"regexp"
	"strings"

	"github.com/pkg/errors"
	log "github.com/sirupsen/logrus"
)

type ShellManager struct {
}

func NewShellManager() *ShellManager {
	return &ShellManager{}
}

func (m *ShellManager) Check() error {
	if version, err := exec.Command("conda", "--version").Output(); err != nil {
		return errors.Wrapf(err, "conda not installed")
	} else {
		log.Printf("Using: %s", version)
	}

	return nil
}

func (m *ShellManager) Index(dir string) error {
	cmd := []string{"index", dir}

	if _, err := exec.Command("conda", cmd...).Output(); err != nil {
		return errors.Wrapf(err, "could not index channel '%s'", filepath.Base(dir))
	}

	return nil
}
func (m *ShellManager) Update() error {
	exists, err := m.condaBuildExist()
	if err != nil {
		return errors.Wrap(err, " could not install conda build")
	}

	if !exists {
		return m.installCondaBuild()
	}

	return nil
}

func (m *ShellManager) condaBuildExist() (bool, error) {
	cmd := []string{"list", "-f", "conda-build"}
	output, err := exec.Command("conda", cmd...).CombinedOutput()
	if err != nil {
		return false, errors.Wrapf(err, "could not execute conda command. Is it installed?")
	}

	return regexp.MatchString("conda-build", strings.TrimSpace(string(output)))
}

func (m *ShellManager) installCondaBuild() error {
	cmd := []string{"install", "-y", "conda-build"}
	if _, err := exec.Command("conda", cmd...).Output(); err != nil {
		return errors.Wrap(err, "could not install conda-build package")
	}
	return nil
}
