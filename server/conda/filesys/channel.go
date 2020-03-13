package filesys

import (
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"

	"github.com/hashicorp/go-multierror"
	"github.com/pkg/errors"
	log "github.com/sirupsen/logrus"

	"private-conda-repo/conda"
	"private-conda-repo/conda/condatypes"
	"private-conda-repo/libs"
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

func (c *Channel) Name() string {
	return c.name
}

func (c *Channel) Dir() string {
	return c.dir
}

func (c *Channel) Index() error {
	return indMgr.Index(c.dir)
}

func (c *Channel) GetMetaInfo() (*condatypes.ChannelMetaInfo, error) {
	var data condatypes.ChannelMetaInfo

	if err := decodeJsonFile(filepath.Join(c.dir, "channeldata.json"), &data); err != nil {
		return nil, errors.Wrapf(err, "error decoding '%s' meta info", c.name)
	}
	return &data, nil
}

func (c *Channel) AddPackage(file io.Reader, pkg *condatypes.Package) (*condatypes.Package, error) {
	if c.packageExists(pkg) {
		err := c.RemoveSinglePackage(pkg)
		if err != nil {
			return nil, errors.New("could not remove existing package")
		}
	}

	destPath := c.packagePath(pkg)
	newFile, err := os.Create(destPath)
	if err != nil {
		return nil, errors.Wrapf(err, "error creating package '%s' in channel '%s' for platform '%s'", pkg.Name, c.name, pkg.Platform)
	}

	defer func() {
		if err := newFile.Close(); err != nil {
			log.Println(errors.Wrapf(err, "Could not close created file: %v\n", destPath))
		}
	}()

	_, err = io.Copy(newFile, file)
	if err != nil && err != io.EOF {
		return nil, errors.Wrapf(err, "error saving package '%s' in channel '%s' for platform '%s' to disk", pkg.Name, c.name, pkg.Platform)
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
	if err := c.Index(); err != nil {
		return count, err
	}

	if err := c.removeEntryFromMetaInfo(name); err != nil {
		return count, err
	}

	return count, nil
}

func (c *Channel) packagePath(pkg *condatypes.Package) string {
	return filepath.Join(c.dir, pkg.Platform, pkg.Filename())
}

func (c *Channel) packageExists(pkg *condatypes.Package) bool {
	return libs.PathExists(c.packagePath(pkg))
}

// This should be called whenever an entire package is removed from the channel.
// The function will remove the entry from the channeldata.json file
func (c *Channel) removeEntryFromMetaInfo(name string) error {
	meta, err := c.GetMetaInfo()
	if err != nil {
		return errors.Wrap(err, "could not rewrite channel metadata")
	}

	if _, exists := meta.Packages[name]; exists {
		delete(meta.Packages, name)

		err := meta.Write(filepath.Join(c.dir, "channeldata.json"))
		if err != nil {
			return err
		}
	}
	return nil
}

func (c *Channel) retrievePackageFamily(name string) ([]string, error) {
	var packages []string

	for _, p := range platforms {
		dir := filepath.Join(c.dir, string(p))
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
