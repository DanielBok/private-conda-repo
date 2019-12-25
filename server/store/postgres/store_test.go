package postgres

import (
	"context"
	"database/sql"
	"strconv"
	"time"

	"github.com/dhui/dktest"
	"github.com/jinzhu/gorm"
	"github.com/pkg/errors"
	"github.com/spf13/viper"

	"private-conda-repo/config"
)

var (
	imageName            = "postgres:12-alpine"
	postgresImageOptions = dktest.Options{
		ReadyFunc:    dbReady,
		PortRequired: true,
		ReadyTimeout: 5 * time.Minute,
		Env: map[string]string{
			"POSTGRES_USER":     "user",
			"POSTGRES_PASSWORD": "password",
			"POSTGRES_DB":       "pcrdb",
		},
	}
)

func dbReady(ctx context.Context, c dktest.ContainerInfo) bool {
	err := setupConfig(c)
	if err != nil {
		return false
	}
	conf, err := config.New()
	if err != nil {
		return false
	}

	db, err := sql.Open("postgres", conf.DB.ConnectionString())
	if err != nil {
		return false
	}
	defer func() { _ = db.Close() }()

	return db.PingContext(ctx) == nil
}

func setupConfig(c dktest.ContainerInfo) error {
	ip, port, err := c.FirstPort()
	if err != nil {
		return err
	}

	if portNo, err := strconv.Atoi(port); err != nil {
		return err
	} else {
		viper.Set("DB.PORT", portNo)
		viper.Set("DB.HOST", ip)
	}

	return nil
}

func newTestDb() (*Store, error) {
	conf, err := config.New()
	if err != nil {
		return nil, err
	}

	cs := conf.DB.ConnectionString()
	db, err := gorm.Open("postgres", cs)
	if err != nil {
		return nil, errors.Wrapf(err, "could not connect to database with '%s'", cs)
	}

	store := &Store{db: db, conf: conf}

	if err = store.Migrate(); err != nil {
		return nil, err
	}

	return store, nil
}
