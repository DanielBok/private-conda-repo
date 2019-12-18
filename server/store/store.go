package store

import (
	"strings"

	"github.com/pkg/errors"
	log "github.com/sirupsen/logrus"

	"private-conda-repo/config"
	"private-conda-repo/store/models"
	"private-conda-repo/store/postgres"
)

type Store interface {
	Migrate() error

	AddUser(name, password string) (*models.User, error)
	GetUser(name string) (*models.User, error)
	RemoveUser(name, password string) error
}

var stores = make(map[string]func() (Store, error))

func Register(name string, creator func() (Store, error)) {
	if _, dup := stores[name]; dup {
		log.Fatalf("%s store type called twice.", name)
	}
	stores[name] = creator
}

func New() (Store, error) {
	conf, err := config.New()
	if err != nil {
		return nil, err
	}
	switch strings.ToLower(conf.DB.Type) {
	case "postgres":
		return postgres.New()
	default:
		return nil, errors.Errorf("Unknown store type: '%s'", conf.DB.Type)
	}
}
