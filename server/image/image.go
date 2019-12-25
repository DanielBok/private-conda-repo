package image

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os/exec"
	"strconv"
	"strings"

	"github.com/pkg/errors"
)

type DockerImageInfo struct {
	Results []struct {
		Name string `json:"name"`
	} `json:"results"`
}

func (m *Manager) UpdateImage() (int, error) {
	current, err := m.checkCurrentVersion()
	if err != nil {
		return -1, err
	}

	latest, err := m.checkLatestVersion()
	if err != nil {
		return -1, err
	}

	if latest > current {
		if err := m.pullLatestImage(latest); err != nil {
			return -1, err
		}
	}

	return latest, nil
}

func (m *Manager) checkCurrentVersion() (int, error) {
	cmd := exec.Command("docker", "image", "list", "--format", "{{.Tag}}", m.image)
	output, err := cmd.CombinedOutput()
	if err != nil {
		return -1, err
	}
	out := strings.TrimSpace(string(output))
	current := 0

	if len(out) == 0 {
		return current, nil
	}

	for _, tag := range strings.Split(out, "\n") {
		tag = strings.TrimSpace(tag)
		if tag == "latest" {
			continue
		}

		i, err := strconv.Atoi(tag)
		if err != nil {
			return -1, errors.Errorf("could not parse image tag: '%s'", tag)
		}
		if i > current {
			current = i
		}
	}

	return current, nil
}

func (m *Manager) checkLatestVersion() (int, error) {
	resp, err := http.Get(fmt.Sprintf("https://hub.docker.com/v2/repositories/%s/tags", m.image))
	if err != nil {
		return -1, errors.Wrapf(err, "could not fetch latest image tags for '%s' from docker hub", m.image)
	}
	defer func() { _ = resp.Body.Close() }()

	var output DockerImageInfo
	if err := json.NewDecoder(resp.Body).Decode(&output); err != nil {
		return -1, errors.Wrap(err, "could not decode image meta data from docker hub")
	}

	latest := 0
	for _, res := range output.Results {
		if res.Name == "latest" {
			continue
		}
		i, err := strconv.Atoi(res.Name)
		if err != nil {
			return -1, errors.Errorf("could not parse image tag: '%s'", res.Name)
		}
		if i > latest {
			latest = i
		}
	}

	return latest, nil
}

func (m *Manager) pullLatestImage(tag int) error {
	image := fmt.Sprintf("%s:%d", m.image, tag)

	cmd := exec.Command("docker", "image", "pull", image)
	if err := cmd.Run(); err != nil {
		return errors.Wrapf(err, "could not pull image '%s' from docker hub", image)
	}

	cmd = exec.Command("docker", "image", "tag", image, m.image+":latest")
	if err := cmd.Run(); err != nil {
		return errors.Wrapf(err, "could not tag image '%s' to latest", image)
	}

	return nil
}
