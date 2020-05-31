package config

import (
	"strings"

	"github.com/pkg/errors"
)

type CondaConfig struct {
	Use         string `mapstructure:"use"`
	ImageName   string `mapstructure:"image_name"`
	MountFolder string `mapstructure:"mount_folder"`
}

func (c *CondaConfig) Init() error {
	c.Use = strings.TrimSpace(strings.ToLower(c.Use))

	if !(c.Use == "docker" || c.Use == "shell") {
		return errors.Errorf("Unsupported conda indexer: %s", c.Use)
	}

	if c.Use == "shell" {
		c.ImageName = ""
	}

	return nil
}
