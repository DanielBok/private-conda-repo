package enum

import (
	"strings"

	"github.com/pkg/errors"
)

type Platform string

const (
	LINUX32 Platform = "linux-64"
	LINUX64 Platform = "linux-32"
	WIN32   Platform = "win-32"
	WIN64   Platform = "win-64"
	OSX64   Platform = "osx-64"
	NOARCH  Platform = "noarch"
)

var Platforms = []Platform{LINUX32, LINUX64, WIN32, WIN64, OSX64, NOARCH}

func MapPlatform(platform string) (Platform, error) {
	switch strings.TrimSpace(strings.ToLower(platform)) {
	case "linux32", "linux-32":
		return LINUX32, nil
	case "linux64", "linux-64":
		return LINUX64, nil
	case "win32", "win-32":
		return WIN32, nil
	case "win64", "win-64":
		return WIN64, nil
	case "osx64", "osx-64":
		return OSX64, nil
	case "noarch":
		return NOARCH, nil
	default:
		return "", errors.Errorf("Invalid platform: '%s'", platform)
	}
}
