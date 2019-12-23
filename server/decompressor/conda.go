package decompressor

import (
	"archive/tar"
	"bufio"
	"io"
	"io/ioutil"
	"os"
	"regexp"
	"strconv"
	"strings"

	"github.com/mholt/archiver/v3"
	"github.com/pkg/errors"

	"private-conda-repo/conda/condatypes"
)

type tarBz2Decompressor struct{}

func (b *tarBz2Decompressor) RetrieveMetadata(f io.ReadCloser) (*Package, error) {
	defer func() { _ = f.Close() }()
	archive, err := b.savePackageToDisk(f)
	if err != nil {
		return nil, err
	}
	defer func() { _ = archive.Close() }()

	pkg, err := b.extractPackageDetail(archive.Name())
	if err != nil {
		return nil, errors.Wrap(err, "could not retrieve package details from index.json in package .tar.gz file")
	}

	return &Package{
		Package:  pkg,
		Filepath: archive.Name(),
	}, nil
}

func (b *tarBz2Decompressor) savePackageToDisk(f io.ReadCloser) (*os.File, error) {
	tmpFile, err := ioutil.TempFile("", "conda-package-*.tar.bz2")
	if err != nil {
		return nil, errors.Wrap(err, "could not create temp file")
	}
	_, err = io.Copy(tmpFile, f)
	if err != nil {
		return nil, errors.Wrap(err, "could not save package to disk")
	}

	return tmpFile, err
}

func (b *tarBz2Decompressor) extractPackageDetail(archivePath string) (*condatypes.Package, error) {
	var pkg *condatypes.Package
	var err error

	tarBz2 := archiver.NewTarBz2()
	defer func() { _ = tarBz2.Close() }()

	err = tarBz2.Walk(archivePath, func(f archiver.File) error {
		if f.Header.(*tar.Header).Name == "info/index.json" {
			defer func() { _ = f.Close() }()

			pkg, err = b.packageDetail(f)
			if err != nil {
				return err
			}
		}
		return nil
	})
	if err != nil {
		return nil, errors.Wrap(err, "could not retrieve package details from index.json in package .tar.gz file")
	}

	return pkg, nil
}

func (b *tarBz2Decompressor) packageDetail(f io.Reader) (*condatypes.Package, error) {
	var err error
	re := regexp.MustCompile(`"([\w\-]+)": "?([\w\-.]+)"?`)
	pkg := &condatypes.Package{}

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		matches := re.FindStringSubmatch(scanner.Text())
		if len(matches) == 3 {
			switch matches[1] {
			case "name":
				if name := strings.TrimSpace(matches[2]); name == "" {
					return nil, errors.New("package name cannot be empty")
				} else {
					pkg.Name = name
				}

			case "build":
				switch parts := strings.Split(matches[2], "_"); {
				case len(parts) == 0:
					return nil, errors.Errorf("no build string specified")
				case len(parts) == 1:
					pkg.BuildString = parts[0]
				default:
					pkg.BuildString = strings.Join(parts[:len(parts)-1], "_")
				}

			case "build_number":
				pkg.BuildNumber, err = strconv.Atoi(matches[2])
				if err != nil {

				}

			case "subdir":
				if _, err := condatypes.MapPlatform(matches[2]); err != nil {
					return nil, errors.Wrapf(err, "invalid platform '%s' specified in package", matches[2])
				}
				pkg.Platform = strings.ToLower(strings.TrimSpace(matches[2]))

			case "version":
				pkg.Version = matches[2]
				pkg.Filename()
			}
		}

	}

	return pkg, nil
}
