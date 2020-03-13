package filesys

import (
	"log"

	"private-conda-repo/config"
	"private-conda-repo/indexer"
)

var indMgr indexer.Indexer

func init() {
	conf, err := config.New()
	if err != nil {
		log.Fatal(err)
	}

	indMgr, err = indexer.New(conf.Conda)
	if err != nil {
		log.Fatal(err)
	}
}
