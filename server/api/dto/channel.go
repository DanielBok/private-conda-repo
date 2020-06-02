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
	CreatedOn time.Time `json:"createdOn"`
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

// Creates a new ChannelDto object from the channel entity where the password is automatically masked
func NewChannelDto(channel *entity.Channel) *ChannelDto {
	return &ChannelDto{
		Id:        channel.Id,
		Channel:   channel.Channel,
		Email:     channel.Email,
		CreatedOn: channel.CreatedOn,
	}
}

// Contains information about the channel in general. This file is located at the root of the channel
// folder. It does not inform anything about packages in the channel
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

// A DTO containing the channel's credentials and an optional package specification
// This DTO is used for removal of packages. If 'package' is not specified, it'll means
// to remove all packages in the channel. If 'package' is specified, it means to remove
// the specific package (version, build, etc) in the channel.
type ChannelPackage struct {
	Channel  string      `json:"channel"`
	Password string      `json:"password"`
	Package  *PackageDto `json:"package,omitempty"`
}
