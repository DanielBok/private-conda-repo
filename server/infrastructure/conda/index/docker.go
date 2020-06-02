package index

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os/exec"
	"path/filepath"
	"strconv"
	"strings"

	"github.com/pkg/errors"
	log "github.com/sirupsen/logrus"

	"private-conda-repo/libs"
)

type DockerIndex struct {
	Image string
	repo  string
}

func NewDockerIndex(imageName string) (*DockerIndex, error) {
	imageParts := strings.Split(imageName, "/")
	if len(imageParts) != 2 {
		return nil, errors.Errorf("expected conda image to be in the form <repo>/<Image name> but got '%s' instead", imageName)
	}

	mgr := &DockerIndex{
		Image: imageName,
		repo:  imageParts[0],
	}

	version, err := mgr.CheckDockerVersion()
	if err != nil {
		return nil, errors.Wrap(err, "could not get docker instance")
	}
	log.Printf("Running docker version: %s", version)

	return mgr, nil
}

func (d *DockerIndex) Index(dir string) error {
	args := []string{
		"container",
		"run",
		"--rm",
		"--mount",
		fmt.Sprintf("type=bind,src=%s,dst=/var/condapkg", dir),
		d.Image,
		"index",
	}

	cmd := exec.Command("docker", args...)
	if err := cmd.Run(); err != nil {
		return errors.Wrapf(err, "could not index channel '%s'", filepath.Base(dir))
	}

	return nil
}

func (d *DockerIndex) Update() error {
	version, err := d.UpdateImage()
	if err != nil {
		return errors.Wrap(err, "could not update docker image")
	}
	log.Printf("Updated conda image to version: %d", version)
	return nil
}

func (d *DockerIndex) CheckDockerVersion() (string, error) {
	cmd := exec.Command("docker", "version", "-f {{.Client.Version}}")
	output, err := cmd.CombinedOutput()
	if err != nil {
		return "", errors.Wrapf(err, "could not get docker client version. Is docker installed?")
	}

	version := strings.TrimSpace(string(output))

	return version, nil
}

func (d *DockerIndex) UpdateImage() (int, error) {
	current, err := d.CheckCurrentVersion()
	if err != nil {
		return -1, err
	}

	latest, err := d.checkLatestVersion()
	if err != nil {
		return -1, err
	}

	if latest > current {
		if err := d.pullLatestImage(latest); err != nil {
			return -1, err
		}
	}

	return latest, nil
}

func (d *DockerIndex) CheckCurrentVersion() (int, error) {
	cmd := exec.Command("docker", "image", "list", "--format", "{{.Tag}}", d.Image)
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
			return -1, errors.Errorf("could not parse Image tag: '%s'", tag)
		}
		if i > current {
			current = i
		}
	}

	return current, nil
}

func (d *DockerIndex) checkLatestVersion() (int, error) {
	resp, err := http.Get(fmt.Sprintf("https://hub.docker.com/v2/repositories/%s/tags", d.Image))
	if err != nil {
		return -1, errors.Wrapf(err, "could not fetch latest Image tags for '%s' from docker hub", d.Image)
	}
	defer libs.IOCloser(resp.Body)

	type ImageInfo struct {
		Results []struct {
			Name string `json:"name"`
		} `json:"results"`
	}

	var output ImageInfo
	if err := json.NewDecoder(resp.Body).Decode(&output); err != nil {
		return -1, errors.Wrap(err, "could not decode Image meta data from docker hub")
	}

	latest := 0
	for _, res := range output.Results {
		if res.Name == "latest" {
			continue
		}
		i, err := strconv.Atoi(res.Name)
		if err != nil {
			return -1, errors.Errorf("could not parse Image tag: '%s'", res.Name)
		}
		if i > latest {
			latest = i
		}
	}

	return latest, nil
}

func (d *DockerIndex) pullLatestImage(tag int) error {
	image := fmt.Sprintf("%s:%d", d.Image, tag)

	cmd := exec.Command("docker", "Image", "pull", image)
	if err := cmd.Run(); err != nil {
		return errors.Wrapf(err, "could not pull Image '%s' from docker hub", image)
	}

	cmd = exec.Command("docker", "Image", "tag", image, d.Image+":latest")
	if err := cmd.Run(); err != nil {
		return errors.Wrapf(err, "could not tag Image '%s' to latest", image)
	}

	return nil
}
