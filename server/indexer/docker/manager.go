package docker

import (
	"fmt"
	"os/exec"
	"path/filepath"
	"strings"

	"github.com/pkg/errors"
	log "github.com/sirupsen/logrus"

	"private-conda-repo/config"
	"private-conda-repo/indexer"
)

type Manager struct {
	Image string
	repo  string
}

func init() {
	indexer.Register("docker", New)
}

func (m *Manager) Index(dir string) error {
	cmd := []string{
		"container",
		"run",
		"--rm",
		"--mount",
		fmt.Sprintf("type=bind,src=%s,dst=/var/condapkg", dir),
		m.Image,
		"index",
	}

	if _, err := exec.Command("docker", cmd...).Output(); err != nil {
		return errors.Wrapf(err, "could not index channel '%s'", filepath.Base(dir))
	}

	return nil
}

func (m *Manager) Check() error {
	version, err := m.CheckDockerVersion()
	if err != nil {
		return errors.Wrap(err, "could not get docker instance")
	}
	log.Printf("Running docker version: %s", version)
	return nil
}

func (m *Manager) Update() error {
	version, err := m.UpdateImage()
	if err != nil {
		return errors.Wrap(err, "could not update docker image")
	}
	log.Printf("Updated conda image to version: %d", version)
	return nil
}

func New() (indexer.Indexer, error) {
	conf, err := config.New()
	if err != nil {
		return nil, err
	}

	image := conf.Conda.ImageName
	imageParts := strings.Split(image, "/")
	if len(imageParts) != 2 {
		return nil, errors.Errorf("expected conda Image to be in the form <repo>/<Image name> but got '%s' instead", image)
	}

	mgr := &Manager{
		Image: image,
		repo:  imageParts[0],
	}

	return mgr, nil
}

func (m *Manager) CheckDockerVersion() (string, error) {
	cmd := exec.Command("docker", "version", "-f {{.Client.Version}}")
	output, err := cmd.CombinedOutput()
	if err != nil {
		return "", errors.Wrapf(err, "could not get docker client version. Is docker installed?")
	}

	version := strings.TrimSpace(string(output))

	return version, nil
}
