package filesys

import (
	"encoding/json"
	"os"
	"strings"

	"github.com/pkg/errors"

	"private-conda-repo/conda/condatypes"
)

var platforms = []condatypes.Platform{condatypes.LINUX32, condatypes.LINUX64, condatypes.WIN32, condatypes.WIN64, condatypes.OSX64, condatypes.NOARCH}

func formatChannel(channel string) (string, error) {
	channel = strings.TrimSpace(channel)
	if channel == "" {
		return "", errors.New("channel cannot be empty")
	}

	return channel, nil
}

func decodeJsonFile(filepath string, object interface{}) error {
	jsonFile, err := os.Open(filepath)
	if os.IsNotExist(err) {
		return errors.Wrap(err, "json file does not exist")
	}
	defer func() { _ = jsonFile.Close() }()

	if err = json.NewDecoder(jsonFile).Decode(&object); err != nil {
		return err
	}
	return nil
}
