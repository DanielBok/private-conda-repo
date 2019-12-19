package filesys

import (
	"strings"

	"github.com/emirpasic/gods/sets/linkedhashset"
	"github.com/pkg/errors"

	"private-conda-repo/conda/condatypes"
)

var platforms = linkedhashset.New(condatypes.LINUX32, condatypes.LINUX64, condatypes.WIN32, condatypes.WIN64, condatypes.OSX64, condatypes.NOARCH)

func formatChannel(channel string) (string, error) {
	channel = strings.TrimSpace(channel)
	if channel == "" {
		return "", errors.New("channel cannot be empty")
	}

	return channel, nil
}
