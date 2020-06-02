package index

import (
	"os/exec"
	"path/filepath"
	"regexp"
	"strings"

	"github.com/pkg/errors"
	log "github.com/sirupsen/logrus"
)

type ShellIndex struct {
}

func NewShellIndex() (*ShellIndex, error) {
	if version, err := exec.Command("conda", "--version").Output(); err != nil {
		return nil, errors.Wrapf(err, "conda not installed")
	} else {
		log.Printf("Using: %s", version)
	}

	return &ShellIndex{}, nil
}

func (s *ShellIndex) Index(dir string) error {
	cmd := []string{"index", dir}

	if _, err := exec.Command("conda", cmd...).Output(); err != nil {
		return errors.Wrapf(err, "could not index channel '%s'", filepath.Base(dir))
	}

	return nil
}
func (s *ShellIndex) Update() error {
	exists, err := s.condaBuildExist()
	if err != nil {
		return errors.Wrap(err, " could not install conda build")
	}

	if !exists {
		return s.installCondaBuild()
	}

	return nil
}

func (s *ShellIndex) condaBuildExist() (bool, error) {
	cmd := []string{"list", "-f", "conda-build"}
	output, err := exec.Command("conda", cmd...).CombinedOutput()
	if err != nil {
		return false, errors.Wrapf(err, "could not execute conda command. Is it installed?")
	}

	return regexp.MatchString("conda-build", strings.TrimSpace(string(output)))
}

func (s *ShellIndex) installCondaBuild() error {
	cmd := []string{"install", "-y", "conda-build"}
	if _, err := exec.Command("conda", cmd...).Output(); err != nil {
		return errors.Wrap(err, "could not install conda-build package")
	}
	return nil
}
