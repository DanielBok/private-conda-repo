package postgres

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/pkg/errors"

	"private-conda-repo/config"
	"private-conda-repo/store"
)

type Store struct {
	db   *gorm.DB
	conf *config.AppConfig
}

func init() {
	store.Register("postgres", New)
}

func New() (store.Store, error) {
	conf, err := config.New()
	if err != nil {
		return nil, err
	}

	cs := conf.DB.ConnectionString()
	db, err := gorm.Open("postgres", cs)
	if err != nil {
		return nil, errors.Wrapf(err, "could not connect to database with '%s'", cs)
	}

	return &Store{db: db, conf: conf}, nil
}
