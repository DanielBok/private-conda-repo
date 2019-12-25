package application

import (
	"private-conda-repo/conda"
	"private-conda-repo/decompressor"
	"private-conda-repo/image"
	"private-conda-repo/store"
)

var (
	db   store.Store
	repo conda.Conda
	dcp  decompressor.Decompressor
	mgr  *image.Manager
)

func initStore() error {
	_db, err := store.New()
	if err != nil {
		return err
	}

	_repo, err := conda.New()
	if err != nil {
		return err
	}

	_dcp, err := decompressor.New()
	if err != nil {
		return err
	}

	_mgr, err := image.New()
	if err != nil {
		return err
	}

	db = _db
	repo = _repo
	dcp = _dcp
	mgr = _mgr

	return nil
}
