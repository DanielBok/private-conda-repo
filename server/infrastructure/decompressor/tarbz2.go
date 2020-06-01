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

	"private-conda-repo/api/dto"
	"private-conda-repo/domain/enum"
	"private-conda-repo/libs"
)

type TarBz2Decompressor struct{}

func New() *TarBz2Decompressor {
	return &TarBz2Decompressor{}
}

// Retrieves MetaData from the .tar.bz2 file
func (b *TarBz2Decompressor) RetrieveMetadata(f io.ReadCloser) (*MetaData, error) {
	defer libs.IOCloser(f)

	archive, err := savePackageToDisk(f)
	if err != nil {
		return nil, err
	}
	defer libs.IOCloser(archive)

	pkg, err := extractPackageDetail(archive.Name())
	if err != nil {
		return nil, errors.Wrap(err, "could not retrieve package details from index.json in package .tar.gz file")
	}

	return &MetaData{
		Package:  pkg,
		Filepath: archive.Name(),
	}, nil
}

// Saves the incoming file (usually from a HTTP form post to the disk
func savePackageToDisk(f io.ReadCloser) (*os.File, error) {
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

// Walks through the unzipped .tar.bz2 package and extracts information from the various json files
// within it. At the moment, only information from the info/index.json file is extracted
func extractPackageDetail(archivePath string) (*dto.PackageDto, error) {
	var pkg *dto.PackageDto
	var err error

	tarBz2 := archiver.NewTarBz2()
	defer libs.IOCloser(tarBz2)

	err = tarBz2.Walk(archivePath, func(f archiver.File) error {
		if f.Header.(*tar.Header).Name == "info/index.json" {
			defer libs.IOCloser(f)

			pkg, err = readDetailsFromInfoIndexJson(f)
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

// Reads package details from the info/index.json file which is bundled in the conda package
func readDetailsFromInfoIndexJson(f io.Reader) (*dto.PackageDto, error) {
	var err error
	re := regexp.MustCompile(`"([\w\-]+)": "?([\w\-.]+)"?`)
	pkg := &dto.PackageDto{}

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
				if _, err := enum.MapPlatform(matches[2]); err != nil {
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
