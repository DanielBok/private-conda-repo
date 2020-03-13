package indexer

import (
	"strings"

	"github.com/pkg/errors"

	"private-conda-repo/config"
)

type Indexer interface {
	Index(dir string) error
	Check() error
	Update() error
}

func New(conf config.CondaConfig) (Indexer, error) {
	switch strings.ToLower(conf.Use) {
	case "docker":
		return NewDockerIndexer(conf.ImageName)
	case "shell":
		return NewShellManager(), nil
	default:
		return nil, errors.Errorf("Unsupported Conda manager type: %s", conf.Use)
	}
}
