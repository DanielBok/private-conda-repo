package postgres

import (
	"time"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/pkg/errors"
	log "github.com/sirupsen/logrus"

	"private-conda-repo/config"
)

type Postgres struct {
	db *gorm.DB
}

func New(config *config.DbConfig) (*Postgres, error) {
	var err error

	// waiting for db to be ready
	wait := 1
	for i := 1; i < 10; i++ {
		db, err := gorm.Open("postgres", config.ConnectionString())
		if err == nil {
			db.SingularTable(true)
			return &Postgres{db: db}, nil
		}
		wait += i
		log.Infof("waiting %d seconds to retry connection with database at %s", wait, config.Host)
		time.Sleep(time.Duration(wait) * time.Second)
	}

	return nil, errors.Wrapf(err, "could not connect to database with '%s'", config.MaskedConnectionString())
}
