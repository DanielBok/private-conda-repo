package filesys

import (
	"os"
	"path/filepath"

	"github.com/hashicorp/go-multierror"
	"github.com/pkg/errors"
	log "github.com/sirupsen/logrus"

	"private-conda-repo/conda"
	"private-conda-repo/config"
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

	for _, p := range platforms.Values() {
		path := filepath.Join(chn.Dir(), p.(string))
		if _, err := os.Stat(path); os.IsNotExist(err) {
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
	if _, err := os.Stat(path); os.IsNotExist(err) {
		return nil, errors.Errorf("Channel '%s' does not exist", path)
	}

	return newChannel(channel, c.dir, c.image)
}

func (c *Conda) RemoveChannel(channel string) error {
	chn, err := c.GetChannel(channel)
	if err != nil {
		return err
	}

	err = os.RemoveAll(chn.Dir())
	if err != nil {
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
	} else if _, err := os.Stat(newFolder); os.IsExist(err) {
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
