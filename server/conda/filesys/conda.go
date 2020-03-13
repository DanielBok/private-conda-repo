package filesys

import (
	"io/ioutil"
	"os"
	"path/filepath"

	"github.com/hashicorp/go-multierror"
	"github.com/pkg/errors"
	log "github.com/sirupsen/logrus"

	"private-conda-repo/conda"
	"private-conda-repo/conda/condatypes"
	"private-conda-repo/config"
	"private-conda-repo/libs"
)

type Conda struct {
	dir   string
	image string
}

func init() {
	conf, err := config.New()
	if err != nil {
		log.Fatalln(err)
	}

	conda.Register("filesys", &Conda{
		dir:   conf.Conda.MountFolder,
		image: conf.Conda.ImageName,
	})
}

func (c *Conda) getChannelPath(channel string) (string, error) {
	channel, err := formatChannel(channel)
	if err != nil {
		return "", err
	}

	return filepath.Join(c.dir, channel), nil
}

func (c Conda) CreateChannel(channel string) (conda.Channel, error) {
	chn, err := newChannel(channel, c.dir, c.image)
	if err != nil {
		return nil, err
	}

	for _, p := range platforms {
		path := filepath.Join(chn.Dir(), string(p))
		if libs.PathExists(path) {
			err = os.MkdirAll(path, os.ModePerm)
			if err != nil {
				return nil, errors.Wrapf(err, "error creating platform '%s' for channel '%s'", p, channel)
			}
		}
	}

	return chn, nil
}

func (c *Conda) GetChannel(channel string) (conda.Channel, error) {
	path, err := c.getChannelPath(channel)
	if err != nil {
		return nil, err
	}
	if libs.PathExists(path) {
		return nil, errors.Errorf("Channel '%s' does not exist", path)
	}

	return newChannel(channel, c.dir, c.image)
}

func (c *Conda) RemoveChannel(channel string) error {
	chn, err := c.GetChannel(channel)
	if err != nil {
		return err
	}

	if libs.PathExists(chn.Dir()) {
		return errors.Errorf("channel '%s' does not exist", channel)
	}

	if err := os.RemoveAll(chn.Dir()); err != nil {
		return errors.Wrapf(err, "error removing channel '%s'", channel)
	}

	return nil
}

func (c *Conda) ChangeChannelName(oldChannel, newChannel string) (conda.Channel, error) {
	var _errors error

	oldChn, err := c.GetChannel(oldChannel)
	if err != nil {
		_errors = multierror.Append(_errors, err)
	}

	newFolder, err := c.getChannelPath(newChannel)
	if err != nil {
		_errors = multierror.Append(_errors, errors.Wrapf(err, "Invalid channel '%s'", newChannel))
	} else if libs.PathExists(newFolder) {
		_errors = multierror.Append(_errors, errors.Wrapf(err, "Channel '%s' already exists", newChannel))
	}

	if _errors != nil {
		return nil, _errors
	}

	err = os.Rename(oldChn.Dir(), newFolder)
	if err != nil {
		return nil, errors.Wrapf(err, "Could not rename channel '%s' to '%s'", oldChannel, newChannel)
	}
	return &Channel{
		name: newChannel,
		dir:  newFolder,
	}, nil
}

func (c *Conda) ListAllChannels() ([]conda.Channel, error) {
	dirs, err := ioutil.ReadDir(c.dir)
	if err != nil {
		return nil, errors.Wrap(err, "could not list channel directories")
	}
	var channels []conda.Channel
	for _, d := range dirs {
		chn, err := c.GetChannel(d.Name())
		if err != nil {
			return nil, errors.Wrapf(err, "error retrieving channel '%s'", d.Name())
		}
		if valid, err := isCondaDirectory(filepath.Join(c.dir, d.Name())); err != nil {
			return nil, err
		} else if valid {
			channels = append(channels, chn)
		}
	}

	return channels, nil
}

func isCondaDirectory(dirname string) (bool, error) {
	dirs, err := ioutil.ReadDir(dirname)
	if err != nil {
		return false, errors.Wrapf(err, "could not verify if directory is a conda dir")
	}

	// get number of unique platforms
	dirMap := make(map[string]int)
	for _, d := range dirs {
		if _, err := condatypes.MapPlatform(d.Name()); err == nil {
			dirMap[d.Name()] = 0
		}
	}

	// check that all platforms match
	return len(dirMap) == len(platforms), nil
}
