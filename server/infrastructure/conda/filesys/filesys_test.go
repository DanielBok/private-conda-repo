package filesys_test

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"testing"
	"time"

	log "github.com/sirupsen/logrus"
	"github.com/stretchr/testify/require"

	"private-conda-repo/api/interfaces"
	. "private-conda-repo/infrastructure/conda/filesys"
	"private-conda-repo/infrastructure/conda/index"
	"private-conda-repo/testutils"
)

var tmpDir string

func TestMain(m *testing.M) {
	dir, err := ioutil.TempDir("", "conda")
	if err != nil {
		log.Fatal(err)
	}
	tmpDir = dir
	code := m.Run()

	time.Sleep(1 * time.Second)
	err = os.RemoveAll(tmpDir)
	if err != nil {
		log.Fatal(err)
	}

	os.Exit(code)
}

func NewFileSys() *FileSys {
	return New(tmpDir, &index.DockerIndex{Image: "danielbok/conda-repo-mgr:latest"})
}

func newPreloadedChannel(name string) (interfaces.Channel, error) {
	repo := NewFileSys()

	chn, err := repo.CreateChannel(name)
	if err != nil {
		return nil, err
	}

	for _, details := range testutils.GetTestPackages() {
		f, err := os.Open(details.Path)
		if err != nil {
			return nil, err
		}

		_, err = chn.AddPackage(f, details.ToPackageDto(), nil)
		if err != nil {
			return nil, err
		}
		_ = f.Close()
	}

	return chn, err
}

func TestFileSys_CRUDChannel(t *testing.T) {
	t.Parallel()
	assert := require.New(t)
	channelName := "crud-conda-channel"
	channelNewName := "crud-conda-channel-new-channel-name"

	repo := NewFileSys()

	_, err := repo.GetChannel(channelName)
	assert.Error(err)

	chn, err := repo.CreateChannel(channelName)
	assert.NoError(err)

	chn, err = repo.GetChannel(channelName)
	assert.NoError(err)
	assert.EqualValues(channelName, filepath.Base(chn.Directory()))

	_, err = repo.RenameChannel(channelName, channelName)
	assert.Error(err, "should not be able to replace existing channel")

	chn, err = repo.RenameChannel(channelName, channelNewName)
	assert.NoError(err)
	assert.EqualValues(channelNewName, filepath.Base(chn.Directory()))

	err = repo.RemoveChannel(channelName)
	assert.Error(err, "channel does not exist and should raise error if trying to remove")

	err = repo.RemoveChannel(channelNewName)
	assert.NoError(err)

	// Test listing all channels
	numChannels := 10
	for i := 0; i < numChannels; i++ {
		_, err := repo.CreateChannel(fmt.Sprintf("test-channel-%d", i))
		assert.NoError(err)
	}

	allChannels, err := repo.ListAllChannels()
	assert.NoError(err)
	assert.GreaterOrEqual(len(allChannels), numChannels)
}
