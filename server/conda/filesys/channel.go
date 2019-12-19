package filesys

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
	"os/exec"
	"path/filepath"

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
	cmd := []string{
		"container",
		"run",
		"--rm",
		"--mount",
		fmt.Sprintf("type=bind,src=%s,dst=/var/condapkg", c.dir),
		c.image,
		"index",
	}

	if _, err := exec.Command("docker", cmd...).Output(); err != nil {
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
	defer jsonFile.Close()

	var data condatypes.ChannelMetaInfo
	if err = json.NewDecoder(jsonFile).Decode(&data); err != nil {
		return nil, errors.Wrapf(err, "error decoding '%s' meta info", c.name)
	}
	return &data, nil
}

func (c *Channel) AddPackage(file io.Reader, platform string, name string) (*condatypes.Package, error) {
	pkg, err := condatypes.PackageFromFileName(name, platform)
	if err != nil {
		return nil, err
	}

	if c.packageExists(pkg) {
		err := c.RemoveSinglePackage(pkg)
		if err != nil {
			return nil, errors.New("could not remove existing package")
		}
	}

	newFile, err := os.Create(c.packagePath(pkg))
	if err != nil {
		return nil, errors.Wrapf(err, "error creating package '%s' in channel '%s' for platform '%s'", pkg.Name, c.name, platform)
	}

	defer func() {
		err = newFile.Close()
		log.Printf("Could not close created file: %v\n", err)
	}()

	_, err = io.Copy(newFile, file)
	if err != nil && err != io.EOF {
		return nil, errors.Wrapf(err, "error saving package '%s' in channel '%s' for platform '%s' to disk", pkg.Name, c.name, platform)
	}

	err = c.Index()
	if err != nil {
		return nil, err
	}

	return pkg, nil
}

func (c *Channel) RemoveSinglePackage(pkg *condatypes.Package) error {
	if !c.packageExists(pkg) {
		return errors.New("package specified does not exist")
	}

	if err := os.Remove(c.packagePath(pkg)); err != nil {
		return errors.Wrapf(err, "error removing package '%+v'", pkg)
	}

	if err := os.Remove(filepath.Join(c.dir, "channeldata.json")); err != nil {
		return errors.Wrap(err, "could not replace channeldata and refresh index")
	}

	if err := c.Index(); err != nil {
		return err
	}

	return nil
}

func (c *Channel) RemovePackageAllVersions(name string) error {
	panic("Implement me")
}

func (c *Channel) packagePath(pkg *condatypes.Package) string {
	return filepath.Join(c.dir, pkg.Platform, pkg.Filename())
}

func (c *Channel) packageExists(pkg *condatypes.Package) bool {
	if _, err := os.Stat(c.packagePath(pkg)); os.IsNotExist(err) {
		return false
	}
	return true
}
