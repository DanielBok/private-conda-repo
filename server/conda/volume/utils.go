package volume

import (
	"strings"

	"github.com/emirpasic/gods/sets/linkedhashset"
	"github.com/pkg/errors"
)

var platforms = linkedhashset.New("linux-64", "linux-32", "osx-64", "win-64", "win-32", "noarch")

func formatPlatform(platform string) (string, error) {
	platform = strings.TrimSpace(strings.ToLower(platform))
	if platforms.Contains(platform) {
		return platform, nil
	}

	return "", errors.Errorf("Unknown platform: %s", platform)
}

func formatChannel(channel string) (string, error) {
	channel = strings.TrimSpace(channel)
	if channel == "" {
		return "", errors.New("channel cannot be empty")
	}

	return channel, nil
}

func formatPackageName(name string) (string, error) {
	name = strings.TrimSpace(name)
	if name == "" {
		return "", errors.New("package name cannot be empty")
	}

	return name, nil
}
