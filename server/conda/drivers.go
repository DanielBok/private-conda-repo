package conda

import (
	"strings"

	"github.com/pkg/errors"

	"private-conda-repo/conda/condamocks"
	"private-conda-repo/conda/types"
	"private-conda-repo/conda/volume"
	"private-conda-repo/config"
)

func New() (types.Conda, error) {
	conf, err := config.New()
	if err != nil {
		return nil, err
	}

	switch strings.ToLower(strings.TrimSpace(conf.Conda.Type)) {
	case "volume":
		return volume.New()
	case "test":
		return condamocks.New()
	default:
		return nil, errors.Errorf("Unknown conda repository type: '%s'", conf.Conda.Type)
	}
}
