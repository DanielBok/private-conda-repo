package dto

import (
	"fmt"
	"strings"
	"time"

	"github.com/pkg/errors"

	"private-conda-repo/domain/entity"
	"private-conda-repo/domain/enum"
)

// Package details that is received from
type PackageDto struct {
	Name        string `json:"name"`
	Version     string `json:"version"`
	BuildString string `json:"buildString"`
	BuildNumber int    `json:"buildNumber"`
	Platform    string `json:"platform"`
}

// Returns the package's full filename (i.e. perfana-0.0.6-py_0.tar.bz2)
func (p *PackageDto) Filename() string {
	return fmt.Sprintf("%s-%s-%s_%d.tar.bz2", p.Name, p.Version, p.BuildString, p.BuildNumber)
}

func (p *PackageDto) GetPlatform() enum.Platform {
	platform, _ := enum.MapPlatform(p.Platform)
	return platform
}

func (p *PackageDto) Validate() error {
	p.Name = strings.TrimSpace(p.Name)
	if p.Name == "" {
		return errors.New("name cannot be empty")
	}

	_, err := enum.MapPlatform(p.Platform)
	if err != nil {
		return err
	}

	return nil
}

func (p *PackageDto) ToPackageCount(channelId int) *entity.PackageCount {
	return &entity.PackageCount{
		ChannelId:   channelId,
		Package:     p.Name,
		BuildString: p.BuildString,
		BuildNumber: p.BuildNumber,
		Version:     p.Version,
		Platform:    p.Platform,
		UploadDate:  time.Now(),
	}
}

// A DTO giving information about the package in the channel. For example, the channel can be
// 'EISR' and the package can be 'numpy'. This DTO will then provide all the details for this
// specification for all versions of the package.
type PackageDetails struct {
	Channel string                 `json:"channel"`
	Package string                 `json:"package"`
	Details []*entity.PackageCount `json:"details"`
	Latest  *ChannelData           `json:"latest"`
}
