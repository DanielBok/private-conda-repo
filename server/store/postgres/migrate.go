package postgres

import (
	"fmt"
	"os"
	"path/filepath"
	"runtime"
	"strings"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/golang-migrate/migrate/v4/source/github"
	"github.com/pkg/errors"
)

func (s *Store) Migrate() error {
	driver, err := postgres.WithInstance(s.db.DB(), &postgres.Config{})
	if err != nil {
		return errors.Wrap(err, "could not create database driver")
	}
	sourceUrl, err := getMigrationSourceUrl()
	if err != nil {
		return err
	}

	m, err := migrate.NewWithDatabaseInstance(sourceUrl, "postgres", driver)
	if err != nil {
		return errors.Wrap(err, "could not create migration instance")
	}

	if err := m.Up(); err != nil && err != migrate.ErrNoChange {
		return errors.Wrap(err, "could not apply migrations")
	}
	return nil
}

func getMigrationSourceUrl() (string, error) {
	pathExists := func(fp string) bool {
		_, err := os.Stat(fp)
		return !os.IsNotExist(err)
	}

	formatFolderPath := func(folder string) string {
		sourceUrl := "file://" + folder
		if runtime.GOOS == "windows" {
			cwd, _ := os.Getwd()
			sourceUrl = strings.Replace(strings.Replace(sourceUrl, cwd, ".", 1), `\`, "/", -1)
		}
		return sourceUrl
	}

	// search from source executable (which is usually the case for Docker images
	root, err := os.Executable()
	if err != nil {
		return "", err
	}
	mgDir := filepath.Join(filepath.Dir(root), "store", "migrations")
	if pathExists(mgDir) {
		return formatFolderPath(mgDir), nil
	}

	// search from local file path, (which is usually the case during development)
	_, file, _, _ := runtime.Caller(0)
	mgDir = filepath.Join(filepath.Dir(file), "..", "migrations")
	if pathExists(mgDir) {
		return formatFolderPath(mgDir), nil
	}

	username := "danielbok"
	publicRepoReadonlyToken := ""
	repo := "private-conda-repo"
	folderPath := "server/store/migrations"
	return fmt.Sprintf("github://%s:%s@%s/%s/%s", username, publicRepoReadonlyToken, username, repo, folderPath), nil
}
