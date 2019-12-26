package fileserver

import (
	"private-conda-repo/store"
)

var db store.Store

func initStore() error {
	_db, err := store.New()
	if err != nil {
		return err
	}

	db = _db

	return nil
}
