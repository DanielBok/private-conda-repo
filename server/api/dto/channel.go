package dto

import (
	"errors"
	"strings"

	"private-conda-repo/domain/condatypes"
	"private-conda-repo/domain/entity"
)

type ChannelData struct {
	Channel     string   `json:"channel"`
	Platforms   []string `json:"platforms"`
	Version     *string  `json:"version"`
	Description *string  `json:"description"`
	DevUrl      *string  `json:"dev_url"`
	DocUrl      *string  `json:"doc_url"`
	Home        *string  `json:"home"`
	License     *string  `json:"license"`
	Summary     *string  `json:"summary"`
	Timestamp   uint64   `json:"timestamp"`
	Name        string   `json:"name"`
}

// Returns the ChannelMetaInfo as an array of normalized data
func ToChannelDataDto(channelData *condatypes.ChannelData, channel string) []*ChannelData {
	var output []*ChannelData
	for name, p := range channelData.Packages {
		output = append(output, &ChannelData{
			Channel:     channel,
			Platforms:   p.Subdirs,
			Version:     p.Version,
			Description: p.Description,
			DevUrl:      p.DevUrl,
			DocUrl:      p.DocUrl,
			Home:        p.Home,
			License:     p.License,
			Summary:     p.Summary,
			Timestamp:   p.Timestamp,
			Name:        name,
		})
	}

	return output
}

type ChannelPackageDetails struct {
	Channel string                 `json:"channel"`
	Package string                 `json:"package"`
	Details []*entity.PackageCount `json:"details"`
	Latest  *ChannelData           `json:"latest"`
}

type ChannelDetails struct {
	Channel  string      `json:"channel"`
	Password string      `json:"password"`
	Package  *PackageDto `json:"package"`
}

func (c *ChannelDetails) Validate() error {
	c.Channel = strings.TrimSpace(c.Channel)
	if c.Channel == "" {
		return errors.New("channel name cannot be empty or whitespaces")
	}
	return nil
}
