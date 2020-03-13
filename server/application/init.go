package application

import (
	"private-conda-repo/conda"
	"private-conda-repo/config"
	"private-conda-repo/decompressor"
	"private-conda-repo/store"
)

var (
	db   store.Store
	repo conda.Conda
	dcp  decompressor.Decompressor
)

func initStore(conf *config.AppConfig) error {
	_db, err := store.New(conf)
	if err != nil {
		return err
	}

	_repo, err := conda.New(conf.Conda.Type)
	if err != nil {
		return err
	}

	_dcp := decompressor.New(conf.Decompressor.Type)

	db = _db
	repo = _repo
	dcp = _dcp

	return nil
}
