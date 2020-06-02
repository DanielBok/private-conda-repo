package config

import (
	"strings"

	"github.com/pkg/errors"
)

type IndexerConfig struct {
	Type        string `mapstructure:"type"`
	ImageName   string `mapstructure:"image_name"`
	MountFolder string `mapstructure:"mount_folder"`
	Update      bool   `mapstructure:"update"`
}

func (c *IndexerConfig) Init() error {
	c.Type = strings.TrimSpace(strings.ToLower(c.Type))

	if !(c.Type == "docker" || c.Type == "shell") {
		return errors.Errorf("Unsupported conda indexer: %s", c.Type)
	}

	if c.Type == "shell" {
		c.ImageName = ""
	}

	return nil
}
