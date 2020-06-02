package filesys

import (
	"encoding/json"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"

	"github.com/hashicorp/go-multierror"
	"github.com/pkg/errors"

	"private-conda-repo/api/dto"
	"private-conda-repo/domain/condatypes"
	"private-conda-repo/domain/enum"

	"private-conda-repo/libs"
)

type Channel struct {
	name string
	dir  string
	ind  Indexer
}

func newChannel(name, dir string, indexer Indexer) (*Channel, error) {
	name = strings.TrimSpace(name)
	if name == "" {
		return nil, errors.New("channel cannot be empty")
	}

	return &Channel{
		name: name,
		dir:  filepath.Join(dir, name),
		ind:  indexer,
	}, nil
}

// Adds a package to the channel
func (c *Channel) AddPackage(file io.Reader, pkg *dto.PackageDto, fixes []string) (*dto.PackageDto, error) {
	if c.packageExists(pkg) {
		err := c.removePackage(pkg)
		if err != nil {
			return nil, errors.New("could not remove existing package")
		}
	}

	destPath := c.packagePath(pkg)
	newFile, err := os.Create(destPath)
	if err != nil {
		return nil, errors.Wrapf(err, "error creating package '%s' in channel '%s' for platform '%s'", pkg.Name, c.name, pkg.Platform)
	}

	defer libs.IOCloserf(newFile, "Could not close created file")

	_, err = io.Copy(newFile, file)
	if err != nil && err != io.EOF {
		return nil, errors.Wrapf(err, "error saving package '%s' in channel '%s' for platform '%s' to disk", pkg.Name, c.name, pkg.Platform)
	}

	err = c.Index(fixes)
	if err != nil {
		return nil, err
	}

	return pkg, nil
}

// Returns the absolute path of channel
func (c *Channel) Directory() string {
	return c.dir
}

// Returns channel's channeldata.json file. This is useful for debugging the current state of the
// channel
func (c *Channel) GetChannelData() (*condatypes.ChannelData, error) {
	var data condatypes.ChannelData

	jsonFile, err := os.Open(c.channelDataFilePath())
	if os.IsNotExist(err) {
		return nil, errors.Wrap(err, "could not find channeldata.json file")
	}
	defer libs.IOCloser(jsonFile)

	if err = json.NewDecoder(jsonFile).Decode(&data); err != nil {
		return nil, errors.Wrap(err, "error decoding channeldata.json")
	}

	return &data, nil
}

// Reindex the channel folder. This should be called whenever there are changes to the packages
// in the channel.
func (c *Channel) Index(fixes []string) error {
	err := c.ind.Index(c.Directory())
	if err != nil {
		return err
	}

	return c.ind.FixRepoData(c.Directory(), fixes)
}

func (c *Channel) Name() string {
	return c.name
}

// Removes a single package from the channel
func (c *Channel) RemoveSinglePackage(pkg *dto.PackageDto) error {
	err := c.removePackage(pkg)
	if err != nil {
		return err
	}

	if err := c.Index(nil); err != nil {
		return err
	}

	packages, err := c.retrievePackageFamily(pkg.Name)
	if err != nil {
		return errors.Wrapf(err, "could not check other directories for '%s' package", pkg.Name)
	}

	if len(packages) == 0 {
		if err := c.removeEntryFromMetaInfo(pkg.Name); err != nil {
			return err
		}
	}

	return nil
}

// Removes all packages of the same name from the channel That is if you have a package called
// numpy with different versions, this method will remove all versions of 'numpy'. Other packages
// like 'scipy' will remain intact. Returns the number of packages removed
func (c *Channel) RemovePackageAllVersions(name string) (int, error) {
	var errs error

	// Remove all matching packages (packages match by the name prefix)
	packages, err := c.retrievePackageFamily(name)
	if err != nil {
		return 0, err
	}

	if len(packages) == 0 {
		return 0, nil
	}

	count := 0
	for _, fp := range packages {
		if err = os.Remove(fp); err != nil {
			errs = multierror.Append(errs, errors.Wrapf(err, "could not remove file at '%s'", fp))
		} else {
			count++
		}
	}
	if errs != nil {
		return count, errs
	}

	// reindex the subdirectories
	if err := c.Index(nil); err != nil {
		return count, err
	}

	if err := c.removeEntryFromMetaInfo(name); err != nil {
		return count, err
	}

	return count, nil
}

func (c *Channel) channelDataFilePath() string {
	return filepath.Join(c.Directory(), "channeldata.json")
}

func (c *Channel) packagePath(pkg *dto.PackageDto) string {
	return filepath.Join(c.Directory(), pkg.Platform, pkg.Filename())
}

func (c *Channel) packageExists(pkg *dto.PackageDto) bool {
	return libs.PathExists(c.packagePath(pkg))
}

// This should be called whenever an entire package is removed from the channel.
// The function will remove the entry from the channeldata.json file
func (c *Channel) removeEntryFromMetaInfo(name string) error {
	data, err := c.GetChannelData()
	if err != nil {
		return errors.Wrap(err, "could not rewrite channeldata.json")
	}

	if _, exists := data.Packages[name]; exists {
		delete(data.Packages, name)

		data, err := json.MarshalIndent(data, "", "  ")
		if err != nil {
			return errors.Wrap(err, "could not marshal ChannelData to json")
		}

		err = ioutil.WriteFile(c.channelDataFilePath(), data, 0644)
		if err != nil {
			return errors.Wrap(err, "could not overwrite ChannelData data")
		}
	}
	return nil
}

func (c *Channel) retrievePackageFamily(name string) ([]string, error) {
	var packages []string

	for _, p := range enum.Platforms {
		dir := filepath.Join(c.Directory(), string(p))
		files, err := ioutil.ReadDir(dir)
		if err != nil {
			return nil, errors.Wrap(err, "could not read files in directory")
		}

		for _, f := range files {
			if strings.HasPrefix(f.Name(), name+"-") {
				fp := filepath.Join(dir, f.Name())
				packages = append(packages, fp)
			}
		}
	}
	return packages, nil
}

// Removes the package without indexing the channel
func (c *Channel) removePackage(pkg *dto.PackageDto) error {
	if !c.packageExists(pkg) {
		return ErrPackageNotFound
	}

	if err := os.Remove(c.packagePath(pkg)); err != nil {
		return errors.Wrapf(err, "error removing package '%+v'", pkg)
	}

	if err := os.Remove(c.channelDataFilePath()); err != nil {
		return errors.Wrap(err, "could not replace channeldata.json")
	}

	return nil
}
