package store

import (
	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/pkg/errors"
)

func (s *Store) Migrate() error {
	driver, err := postgres.WithInstance(s.db.DB(), &postgres.Config{})
	if err != nil {
		return errors.Wrap(err, "could not create database driver")
	}

	m, err := migrate.NewWithDatabaseInstance("file://migrations", "postgres", driver)
	if err != nil {
		return errors.Wrap(err, "could not create migration instance")
	}

	if err := m.Up(); err != nil {
		return errors.Wrap(err, "could not apply migrations")
	}
	return nil
}
