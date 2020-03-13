package fileserver

import (
	"private-conda-repo/config"
	"private-conda-repo/store"
)

var db store.Store

func initStore(conf *config.AppConfig) error {
	_db, err := store.New(conf)
	if err != nil {
		return err
	}

	db = _db

	return nil
}
