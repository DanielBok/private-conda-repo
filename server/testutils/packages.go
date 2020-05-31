package testutils

import (
	"io"
	"net/http"
	"os"
	"path/filepath"
	"regexp"
	"runtime"
	"strconv"
	"strings"
	"sync"

	"github.com/pkg/errors"
	log "github.com/sirupsen/logrus"

	"private-conda-repo/api/dto"
	"private-conda-repo/libs"
)

type TestPackage struct {
	Filename string
	Path     string
	Platform string
}

var (
	urls = [8]string{
		"https://anaconda.org/danielbok/perfana/0.0.6/download/noarch/perfana-0.0.6-py_0.tar.bz2",
		"https://anaconda.org/danielbok/perfana/0.0.5/download/noarch/perfana-0.0.5-py_0.tar.bz2",
		"https://anaconda.org/conda-forge/copulae/0.4.3/download/win-64/copulae-0.4.3-py38hfa6e2cd_1.tar.bz2",
		"https://anaconda.org/conda-forge/copulae/0.4.3/download/osx-64/copulae-0.4.3-py38h0b31af3_1.tar.bz2",
		"https://anaconda.org/conda-forge/copulae/0.4.3/download/linux-64/copulae-0.4.3-py38h516909a_1.tar.bz2",
		"https://anaconda.org/conda-forge/copulae/0.4.2/download/osx-64/copulae-0.4.2-py36h01d97ff_1.tar.bz2",
		"https://anaconda.org/conda-forge/copulae/0.4.2/download/linux-64/copulae-0.4.2-py37h516909a_1.tar.bz2",
		"https://anaconda.org/conda-forge/copulae/0.4.2/download/win-64/copulae-0.4.2-py36hfa6e2cd_1.tar.bz2",
	}
	packages = make(map[string]TestPackage)
)

func packageFolder() string {
	_, filename, _, _ := runtime.Caller(1)
	pkgDir, _ := filepath.Abs(filepath.Join(filepath.Dir(filename), "test-packages"))

	return pkgDir
}

func appendTestPackage(url string, wg *sync.WaitGroup, m *sync.Mutex) {
	parts := strings.Split(url, "/")
	filename := parts[len(parts)-1]
	platform := parts[len(parts)-2]

	pkgPath := filepath.Join(packageFolder(), filename)
	if libs.PathExists(pkgPath) {
		resp, err := http.Get(url)
		if err != nil {
			log.Fatalln(errors.Wrapf(err, "could not download '%s' from '%s'", filename, url))
		}
		defer libs.IOCloser(resp.Body)

		file, err := os.Create(pkgPath)
		if err != nil {
			log.Fatalln(errors.Wrapf(err, "could not create '%s' to path '%s'", filename, pkgPath))
		}
		defer libs.IOCloser(file)

		_, err = io.Copy(file, resp.Body)
		if err != nil {
			log.Fatalln(errors.Wrapf(err, "could not write (copy) '%s' to path '%s'", filename, pkgPath))
		}
	}

	if filename != "" {
		m.Lock()
		packages[filename] = TestPackage{
			Filename: filename,
			Path:     pkgPath,
			Platform: platform,
		}
		m.Unlock()
	}

	wg.Done()
}

func init() {
	var wg sync.WaitGroup
	var m sync.Mutex

	for _, url := range urls {
		wg.Add(1)
		go appendTestPackage(url, &wg, &m)
	}

	wg.Wait()
}

func GetTestPackages() map[string]TestPackage {
	data := make(map[string]TestPackage, len(packages))

	for k, v := range packages {
		data[k] = v
	}
	return data
}

func GetPackageN(n int) (TestPackage, error) {
	for _, v := range packages {
		if n == 0 {
			return v, nil
		}
		n -= 1
	}
	return TestPackage{}, errors.Errorf("n must be between 0 and %d", len(urls)-1)
}

func (t *TestPackage) ToPackageDto() *dto.PackageDto {
	re := regexp.MustCompile(`([\w\-]+)-([\w.]+)-(\w+)_(\d+)\.tar\.bz2`)
	matches := re.FindStringSubmatch(t.Filename)
	n, _ := strconv.Atoi(matches[4])

	return &dto.PackageDto{
		Name:        matches[1],
		Version:     matches[2],
		BuildString: matches[3],
		BuildNumber: n,
		Platform:    t.Platform,
	}
}
