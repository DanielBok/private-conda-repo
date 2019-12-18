package postgres

import (
	"fmt"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/github"
	"github.com/pkg/errors"
)

func (s *Store) Migrate() error {
	driver, err := postgres.WithInstance(s.db.DB(), &postgres.Config{})
	if err != nil {
		return errors.Wrap(err, "could not create database driver")
	}

	username := "danielbok"
	publicRepoReadonlyToken := "258c18256df7f3e77aff672c61cf33a36cc64546"
	repo := "private-conda-repo"
	folderPath := "server/store/migrations"
	sourceUrl := fmt.Sprintf("github://%s:%s@%s/%s/%s", username, publicRepoReadonlyToken, username, repo, folderPath)

	m, err := migrate.NewWithDatabaseInstance(sourceUrl, "postgres", driver)
	if err != nil {
		return errors.Wrap(err, "could not create migration instance")
	}

	if err := m.Up(); err != nil {
		return errors.Wrap(err, "could not apply migrations")
	}
	return nil
}
