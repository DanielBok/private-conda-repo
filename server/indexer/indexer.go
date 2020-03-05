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

func New(conf *config.AppConfig) (Indexer, error) {
	switch strings.ToLower(conf.Conda.Type) {
	case "docker":
		return NewDockerIndexer(conf.Conda.ImageName)
	case "shell":
		return NewShellManager(), nil
	default:
		return nil, errors.Errorf("Unsupported Conda manager type: %s", conf.Conda.Type)
	}
}
