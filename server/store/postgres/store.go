package postgres

import (
	"time"

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

	// waiting for db to be ready
	wait := 1
	for i := 1; i < 10; i++ {
		db, err := gorm.Open("postgres", cs)
		if err == nil {
			return &Store{db: db, conf: conf}, nil
		}
		wait += i
		time.Sleep(time.Duration(wait) * time.Second)

	}

	return nil, errors.Wrapf(err, "could not connect to database with '%s'", cs)
}
