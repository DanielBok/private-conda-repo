package filesys

import (
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"

	"github.com/hashicorp/go-multierror"
	"github.com/pkg/errors"

	"private-conda-repo/api/interfaces"
	"private-conda-repo/domain/enum"
	"private-conda-repo/libs"
)

type FileSys struct {
	dir string
	ind Indexer
}

type Indexer interface {
	// Indexes the directory via `conda Index`. This should be run whenever a package is added,
	// removed or updated. It will update the current_repodata.json and repodata.json files in
	// the repository so that when we `conda install`, the dependency solver will know how to
	// look for files
	Index(dir string) error

	// Applies a series of fixes to repodata.json and current_repodata.json such as removal of
	// `python_abi` from the dependency lists. A list of instructions (fixes) must be provided
	// to determine which fixes to apply. Presently the supported values are
	//
	// - no-abi : Removes "python_abi *" dependencies from the uploaded package
	FixRepoData(dir string, fixes []string) error
}

// Creates a new Filesys which handles the physical files (conda packages) objects. Basically,
// it is a CRUD handler for channels. The Channel instances are then CRUD handlers for the
// physical conda packages.
//
// The directory parameter is the place where channel folders are created. If using the
// application from docker, this is also the "container path" that should be mounted with the
// host. It is where all the packages are added and indexes are created.
//
// The indexer should be either the ShellIndex or DockerIndex. They are used whenever the
// channel changes due to addition or removal of packages. The indexer will ensure that the
// current_repodata.json and repodata.json files are updated so that "conda install" can
// function properly.
func New(directory string, indexer Indexer) *FileSys {
	return &FileSys{
		dir: directory,
		ind: indexer,
	}
}

func (f *FileSys) RenameChannel(oldName, newName string) (interfaces.Channel, error) {
	var errs error

	oldChn, err := f.GetChannel(oldName)
	if err != nil {
		errs = multierror.Append(errs, err)
	}

	newFolder, err := f.getChannelPath(newName)
	if err != nil {
		errs = multierror.Append(errs, errors.Wrapf(err, "Invalid channel '%s'", newName))
	} else if libs.PathExists(newFolder) {
		errs = multierror.Append(errs, errors.Errorf("Channel '%s' already exists", newName))
	}

	if errs != nil {
		return nil, errs
	}

	err = os.Rename(oldChn.Directory(), newFolder)
	if err != nil {
		return nil, errors.Wrapf(err, "Could not rename channel '%s' to '%s'", oldName, newName)
	}

	return newChannel(newName, newFolder, f.ind)
}

func (f FileSys) CreateChannel(name string) (interfaces.Channel, error) {
	chn, err := newChannel(name, f.dir, f.ind)
	if err != nil {
		return nil, err
	}

	for _, p := range enum.Platforms {
		path := filepath.Join(chn.Directory(), string(p))

		// Create folder if it does not exist
		if !libs.PathExists(path) {
			err = os.MkdirAll(path, os.ModePerm)
			if err != nil {
				return nil, errors.Wrapf(err, "error creating platform '%s' for channel name '%s'", p, name)
			}
		}
	}

	return chn, nil
}

func (f *FileSys) GetChannel(name string) (interfaces.Channel, error) {
	name = strings.TrimSpace(name)
	path, err := f.getChannelPath(name)
	if err != nil {
		return nil, err
	}
	if !libs.PathExists(path) {
		return nil, errors.Errorf("Channel '%s' does not exist", name)
	}

	return newChannel(name, f.dir, f.ind)
}

func (f *FileSys) ListAllChannels() ([]interfaces.Channel, error) {
	dirs, err := ioutil.ReadDir(f.dir)
	if err != nil {
		return nil, errors.Wrap(err, "could not list channel directories")
	}

	var channels []interfaces.Channel
	for _, d := range dirs {
		chn, err := f.GetChannel(d.Name())
		if err != nil {
			return nil, errors.Wrapf(err, "error retrieving channel '%s'", d.Name())
		}
		if valid, err := isCondaDirectory(filepath.Join(f.dir, d.Name())); err != nil {
			return nil, err
		} else if valid {
			channels = append(channels, chn)
		}
	}

	return channels, nil
}

func (f *FileSys) RemoveChannel(name string) error {
	chn, err := f.GetChannel(name)
	if err != nil {
		return err
	}

	if !libs.PathExists(chn.Directory()) {
		return errors.Errorf("channel '%s' does not exist", name)
	}

	if err := os.RemoveAll(chn.Directory()); err != nil {
		return errors.Wrapf(err, "error removing channel '%s'", name)
	}

	return nil
}

func (f *FileSys) getChannelPath(channel string) (string, error) {
	channel = strings.TrimSpace(channel)
	if strings.TrimSpace(channel) == "" {
		return "", errors.New("channel cannot be empty or whitespace")
	}

	return filepath.Join(f.dir, channel), nil
}

func isCondaDirectory(dirname string) (bool, error) {
	dirs, err := ioutil.ReadDir(dirname)
	if err != nil {
		return false, errors.Wrapf(err, "could not verify if directory is a conda dir")
	}

	// get number of unique platforms
	dirMap := make(map[string]int)
	for _, d := range dirs {
		if _, err := enum.MapPlatform(d.Name()); err == nil {
			dirMap[d.Name()] = 0
		}
	}

	// check that all platforms match
	return len(dirMap) == len(enum.Platforms), nil
}
