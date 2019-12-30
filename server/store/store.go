package store

import (
	"strings"

	"github.com/pkg/errors"
	log "github.com/sirupsen/logrus"

	"private-conda-repo/config"
	"private-conda-repo/store/models"
)

type Store interface {
	Migrate() error

	AddUser(channel, password string) (*models.User, error)
	GetUser(channel string) (*models.User, error)
	RemoveUser(channel, password string) error
	GetAllUsers() ([]*models.User, error)

	GetPackageCounts(channel, name string) ([]*models.PackageCount, error)
	CreatePackageCount(pkg *models.PackageCount) (*models.PackageCount, error)
	IncreasePackageCount(channel, name, platform, version string) (*models.PackageCount, error)
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

	name := strings.ToLower(strings.TrimSpace(conf.DB.Type))
	if createStore, ok := stores[name]; !ok {
		return nil, errors.Errorf("Unknown database driver: '%s'. Did you forget to '_ import'?", conf.DB.Type)
	} else {
		return createStore()
	}
}
