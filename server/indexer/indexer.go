package indexer

import (
	"strings"

	"github.com/pkg/errors"
	log "github.com/sirupsen/logrus"

	"private-conda-repo/config"
)

type Indexer interface {
	Index(dir string) error
	Check() error
	Update() error
}

var creators = make(map[string]func() (Indexer, error))

func Register(name string, creator func() (Indexer, error)) {
	if _, dup := creators[name]; dup {
		log.Fatalf("%s store type called twice.", name)
	}
	creators[name] = creator
}

func New() (Indexer, error) {
	conf, err := config.New()
	if err != nil {
		return nil, err
	}

	name := strings.ToLower(strings.TrimSpace(conf.Conda.Use))
	if creator, ok := creators[name]; !ok {
		return nil, errors.Errorf("Unknown indexer: '%s'. Did you forget to '_ import'?", conf.Conda.Use)
	} else {
		return creator()
	}
}
