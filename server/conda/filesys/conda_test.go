package filesys

import (
	"io/ioutil"
	"os"
	"path/filepath"
	"testing"

	"github.com/pkg/errors"
	log "github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"

	"private-conda-repo/conda"
	"private-conda-repo/testutils"
)

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

	for _, details := range testutils.GetTestPackages() {
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
