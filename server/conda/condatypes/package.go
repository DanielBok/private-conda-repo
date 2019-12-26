package condatypes

import (
	"fmt"
	"strings"

	"github.com/pkg/errors"

	"private-conda-repo/store/models"
)

type Package struct {
	Name        string `json:"name"`
	Version     string `json:"version"`
	BuildString string `json:"build_string"`
	BuildNumber int    `json:"build_number"`
	Platform    string `json:"platform"`
}

// Returns the package's full filename (i.e. perfana-0.0.6-py_0.tar.bz2)
func (p *Package) Filename() string {
	return fmt.Sprintf("%s-%s-%s_%d.tar.bz2", p.Name, p.Version, p.BuildString, p.BuildNumber)
}

func (p *Package) GetPlatform() Platform {
	platform, _ := MapPlatform(p.Platform)
	return platform
}

func (p *Package) Validate() error {
	p.Name = strings.TrimSpace(p.Name)
	if p.Name == "" {
		return errors.New("name cannot be empty")
	}

	_, err := MapPlatform(p.Platform)
	if err != nil {
		return err
	}

	return nil
}

func (p *Package) ToPackageCount(channel string) *models.PackageCount {
	return &models.PackageCount{
		Channel:     channel,
		Package:     p.Name,
		BuildString: p.BuildString,
		BuildNumber: p.BuildNumber,
		Version:     p.Version,
		Platform:    p.Platform,
	}
}
