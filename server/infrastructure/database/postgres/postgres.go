package postgres

import (
	"time"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/pkg/errors"

	"private-conda-repo/config"
)

type Postgres struct {
	db   *gorm.DB
	salt string
}

func New(config *config.DbConfig, salt string) (*Postgres, error) {
	var err error

	// waiting for db to be ready
	wait := 1
	for i := 1; i < 10; i++ {
		db, err := gorm.Open("postgres", config.ConnectionString())
		if err == nil {
			db.SingularTable(true)
			return &Postgres{
				db:   db,
				salt: salt,
			}, nil
		}
		wait += i
		time.Sleep(time.Duration(wait) * time.Second)
	}

	return nil, errors.Wrapf(err, "could not connect to database with '%s'", config.MaskedConnectionString())
}
