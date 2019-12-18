package filesys

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"os/exec"
	"path/filepath"
	"strings"

	"github.com/pkg/errors"
	log "github.com/sirupsen/logrus"

	"private-conda-repo/conda"
	"private-conda-repo/conda/condatypes"
)

type Channel struct {
	name  string
	dir   string
	image string
}

func newChannel(name, dir, image string) (conda.Channel, error) {
	name, err := formatChannel(name)
	if err != nil {
		return nil, err
	}

	return &Channel{
		name:  name,
		dir:   filepath.Join(dir, name),
		image: image,
	}, nil
}

func (c *Channel) Dir() string {
	return c.dir
}

func (c *Channel) Index() error {
	cmd := fmt.Sprintf("docker container run --rm "+
		"--mount type=bind,src=%s,dst=/var/condapkg "+
		"%s index", c.dir, c.image)

	if _, err := exec.Command("/bin/sh", "-c", cmd).Output(); err != nil {
		return errors.Wrapf(err, "could not index channel '%s'", c.name)
	}

	return nil
}

func (c *Channel) GetMetaInfo() (*condatypes.ChannelMetaInfo, error) {
	file := filepath.Join(c.dir, "channeldata.json")

	jsonFile, err := os.Open(file)
	if os.IsNotExist(err) {
		return nil, nil
	}

	var data condatypes.ChannelMetaInfo
	if err = json.NewDecoder(jsonFile).Decode(&data); err != nil {
		return nil, errors.Wrapf(err, "error decoding '%s' meta info", c.name)
	}
	return &data, nil
}

func (c *Channel) AddPackage(file io.Reader, platform, packageName string) error {
	filePath, err := c.getPackagePath(platform, packageName)
	if err != nil {
		return err
	}

	newFile, err := os.Create(filePath)
	if err != nil {
		return errors.Wrapf(err, "error creating package '%s' in channel '%s' for platform '%s'", packageName, c.name, platform)
	}
	defer func() {
		err = newFile.Close()
		log.Printf("Could not close created file: %v\n", err)
	}()

	_, err = io.Copy(newFile, file)
	if err != nil && err != io.EOF {
		return errors.Wrapf(err, "error saving package '%s' in channel '%s' for platform '%s' to disk", packageName, c.name, platform)
	}

	err = c.Index()
	if err != nil {
		return err
	}

	return nil
}

func (c *Channel) RemovePackage(platform, packageName string) error {
	platform = strings.TrimSpace(platform)
	if platform == "" {
		for _, p := range platforms.Values() {
			err := c.removeSinglePackage(p.(string), packageName)
			if err != nil {
				return err
			}
		}

	} else {
		err := c.removeSinglePackage(platform, packageName)
		if err != nil {
			return err
		}
	}

	err := c.Index()
	if err != nil {
		return err
	}

	return nil
}

func (c *Channel) removeSinglePackage(platform, packageName string) error {
	platformFolder, err := c.getPackagePath(platform, packageName)
	if err != nil {
		return err
	}

	files, err := ioutil.ReadDir(platformFolder)
	if err != nil {
		return errors.Errorf("error reading conda directory. Channel: '%s', platform: '%s'", c.name, platform)
	}

	for _, f := range files {
		fileName := strings.ToLower(f.Name())
		parts := strings.Split(fileName, "-")

		if strings.EqualFold(parts[0], packageName) &&
			strings.HasSuffix(packageName, ".tar.bz2") {
			err := os.Remove(filepath.Join(platformFolder, f.Name()))
			if err != nil {
				return errors.Wrapf(err, "error removing package '%s'", f.Name())
			}
			return nil
		}
	}

	return nil
}

func (c *Channel) getPackagePath(platform, packageName string) (string, error) {
	packageName, err := formatPackageName(packageName)
	if err != nil {
		return "", err
	}

	platform, err = formatPlatform(platform)
	if err != nil {
		return "", err
	}

	return filepath.Join(c.dir, platform, packageName), nil
}
