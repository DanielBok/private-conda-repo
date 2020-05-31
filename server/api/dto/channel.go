package dto

import (
	"errors"
	"regexp"
	"strings"
	"time"

	"github.com/hashicorp/go-multierror"

	"private-conda-repo/domain/condatypes"
	"private-conda-repo/domain/entity"
)

type ChannelDto struct {
	Id        int       `json:"id"`
	Channel   string    `json:"channel"`
	Password  string    `json:"password,omitempty"`
	Email     string    `json:"email"`
	CreatedOn time.Time `json:"created_on"`
}

func (c *ChannelDto) IsValid() error {
	var err error
	nameRegex := regexp.MustCompile(`^\w[\w\-]{2,50}$`)

	c.Channel = strings.TrimSpace(c.Channel)
	c.Password = strings.TrimSpace(c.Password)
	c.Email = strings.TrimSpace(c.Email)

	if !nameRegex.MatchString(c.Channel) {
		err = multierror.Append(err, errors.New("channel name length must be between [2, 50] characters and can only be alphanumeric with dashes"))
	}
	if len(c.Password) < 4 {
		err = multierror.Append(err, errors.New("password must be >= 4 characters"))
	}

	if !strings.Contains(c.Email, "@") {
		err = multierror.Append(err, errors.New("email does not seem valid"))
	}

	return err
}

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
