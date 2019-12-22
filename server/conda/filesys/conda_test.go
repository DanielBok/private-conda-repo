package filesys

import (
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"path/filepath"
	"runtime"
	"strings"
	"sync"
	"testing"

	"github.com/pkg/errors"
	log "github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"

	"private-conda-repo/conda"
)

type testPackage struct {
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
	testPackages = make(map[string]testPackage)
)

func packageFolder() string {
	_, filename, _, _ := runtime.Caller(1)
	pkgDir, _ := filepath.Abs(filepath.Join(filepath.Dir(filename), "testpackages"))

	return pkgDir
}

func appendTestPackage(url string, wg *sync.WaitGroup, m *sync.Mutex) {
	parts := strings.Split(url, "/")
	filename := parts[len(parts)-1]
	platform := parts[len(parts)-2]

	pkgPath := filepath.Join(packageFolder(), filename)
	if _, err := os.Stat(pkgPath); os.IsNotExist(err) {
		resp, err := http.Get(url)
		if err != nil {
			log.Fatalln(errors.Wrapf(err, "could not download '%s' from '%s'", filename, url))
		}
		defer func() { _ = resp.Body.Close() }()

		file, err := os.Create(pkgPath)
		if err != nil {
			log.Fatalln(errors.Wrapf(err, "could not create '%s' to path '%s'", filename, pkgPath))
		}
		defer func() { _ = file.Close() }()

		_, err = io.Copy(file, resp.Body)
		if err != nil {
			log.Fatalln(errors.Wrapf(err, "could not write (copy) '%s' to path '%s'", filename, pkgPath))
		}
	}

	if filename != "" {
		m.Lock()
		testPackages[filename] = testPackage{
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

func newTestConda() (*Conda, func()) {
	tmpdir, err := ioutil.TempDir("", "conda")
	if err != nil {
		log.Fatalln(errors.Wrap(err, "could not create temp directory for conda test"))
	}

	return &Conda{
			dir:   tmpdir,
			image: "danielbok/conda-repo-mgr",
		}, func() {
			_ = os.RemoveAll(tmpdir)
		}
}

func newPreloadedChannel(name string) (conda.Channel, func(), error) {
	repo, cleanup := newTestConda()

	chn, err := repo.CreateChannel(name)
	if err != nil {
		return nil, nil, err
	}

	for _, details := range testPackages {
		f, err := os.Open(details.Path)
		if err != nil {
			return nil, nil, err
		}
		defer func() { _ = f.Close() }()

		_, err = chn.AddPackage(f, details.Platform, details.Filename)
		if err != nil {
			return nil, nil, err
		}
	}

	return chn, cleanup, err
}

func TestConda_CRUDChannel(t *testing.T) {
	t.Parallel()
	assert := assert.New(t)
	channelName := "create-conda-channel-test"
	channelNewName := "new-channel-name"

	repo, cleanup := newTestConda()
	defer cleanup()

	_, err := repo.GetChannel(channelName)
	assert.Error(err)

	chn, err := repo.CreateChannel(channelName)
	assert.NoError(err)

	chn, err = repo.GetChannel(channelName)
	assert.NoError(err)
	assert.EqualValues(channelName, filepath.Base(chn.Dir()))

	_, err = repo.ChangeChannelName(channelName, channelName)
	assert.Error(err, "should not be able to replace existing channel")

	chn, err = repo.ChangeChannelName(channelName, channelNewName)
	assert.NoError(err)
	assert.EqualValues(channelNewName, filepath.Base(chn.Dir()))

	err = repo.RemoveChannel(channelName)
	assert.Error(err, "channel should not exist")

	err = repo.RemoveChannel(channelNewName)
	assert.NoError(err)
}
