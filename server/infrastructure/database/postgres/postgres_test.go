package postgres_test

import (
	"context"
	"database/sql"
	"strconv"
	"time"

	"github.com/dhui/dktest"
	"github.com/pkg/errors"

	"private-conda-repo/config"
	. "private-conda-repo/infrastructure/database/postgres"
	"private-conda-repo/libs"
)

const (
	dbUid  = "user"
	dbPwd  = "password"
	dbName = "pcrdb"
)

var (
	imageName            = "postgres:12-alpine"
	postgresImageOptions = dktest.Options{
		ReadyFunc:    dbReady,
		PortRequired: true,
		ReadyTimeout: 5 * time.Minute,
		Env: map[string]string{
			"POSTGRES_USER":     dbUid,
			"POSTGRES_PASSWORD": dbPwd,
			"POSTGRES_DB":       dbName,
		},
		PullTimeout: 7.5 * 60 * time.Second,
	}
)

func newDbConfig(c dktest.ContainerInfo) (*config.DbConfig, error) {
	host, port, err := c.FirstPort()
	if err != nil {
		return nil, err
	}

	portNo, err := strconv.Atoi(port)
	if err != nil {
		return nil, err
	}

	return &config.DbConfig{
		Host:     host,
		Port:     portNo,
		User:     dbUid,
		Password: dbPwd,
		DbName:   dbName,
	}, nil
}

func dbReady(ctx context.Context, c dktest.ContainerInfo) bool {
	conf, err := newDbConfig(c)
	if err != nil {
		return false
	}

	db, err := sql.Open("postgres", conf.ConnectionString())
	if err != nil {
		return false
	}
	defer libs.IOCloser(db)

	return db.PingContext(ctx) == nil
}

func newTestDb(c dktest.ContainerInfo) (*Postgres, error) {
	conf, err := newDbConfig(c)
	if err != nil {
		return nil, errors.Wrap(err, "could not create connection string from docker info")
	}

	store, err := New(conf)
	if err != nil {
		return nil, err
	}

	if err = store.Migrate(); err != nil {
		return nil, err
	}

	return store, nil
}
