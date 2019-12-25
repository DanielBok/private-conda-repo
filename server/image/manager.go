package image

import (
	"os/exec"
	"strings"

	"github.com/pkg/errors"

	"private-conda-repo/config"
)

type Manager struct {
	image string
	repo  string
}

func New() (*Manager, error) {
	conf, err := config.New()
	if err != nil {
		return nil, err
	}

	image := conf.Conda.ImageName
	imageParts := strings.Split(image, "/")
	if len(imageParts) != 2 {
		return nil, errors.Errorf("expected conda image to be in the form <repo>/<image name> but got '%s' instead", image)
	}

	mgr := &Manager{
		image: image,
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
