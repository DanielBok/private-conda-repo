package condatypes

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"

	"github.com/pkg/errors"
)

type Package struct {
	Name        string `json:"name"`
	Version     string `json:"version"`
	BuildString string `json:"build_string"`
	BuildNumber int    `json:"build_number"`
	Platform    string `json:"platform"`
}

var packageParseRegex = regexp.MustCompile(`([\w\-]+)-([\w.]+)-(\w+)_(\d+)\.tar\.bz2`)

func PackageFromFileName(name string, platform string) (*Package, error) {
	name = strings.ToLower(strings.TrimSpace(name))
	matches := packageParseRegex.FindStringSubmatch(name)

	if len(matches) != 5 {
		return nil, errors.Errorf("Name must be in format " +
			"<package-name: string>-<version: string>-<build string: string>-<build number: int>.tar.bz2")
	}
	matches = matches[1:]

	if matches[0] == "" {
		return nil, errors.New("package name cannot be empty")
	}

	n, err := strconv.Atoi(matches[3])
	if err != nil {
		return nil, errors.New("Build number must be a positive integer")
	}

	platform = strings.ToLower(strings.TrimSpace(platform))
	if _, err := MapPlatform(platform); err != nil {
		return nil, err
	}

	return &Package{
		Name:        matches[0],
		Version:     matches[1],
		BuildString: matches[2],
		BuildNumber: n,
		Platform:    platform,
	}, nil
}

// Returns the package's full filename (i.e. perfana-0.0.6-py_0.tar.bz2)
func (p *Package) Filename() string {
	return fmt.Sprintf("%s-%s-%s_%d.tar.bz2", p.Name, p.Version, p.BuildString, p.BuildNumber)
}

func (p *Package) GetPlatform() Platform {
	platform, _ := MapPlatform(p.Platform)
	return platform
}
