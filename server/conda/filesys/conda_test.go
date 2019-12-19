package filesys

import (
	"io/ioutil"
	"os"
	"path/filepath"
	"runtime"
	"testing"

	"github.com/pkg/errors"
	log "github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
)

var perfana = testPackagePath("perfana-0.0.6-py_0.tar.bz2")

func testPackagePath(pkg string) string {
	_, filename, _, _ := runtime.Caller(1)
	path, _ := filepath.Abs(filepath.Join(filepath.Dir(filename), "testpackages", pkg))

	return path
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
