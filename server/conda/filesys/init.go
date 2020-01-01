package filesys

import (
	"log"

	"private-conda-repo/indexer"
	_ "private-conda-repo/indexer/docker"
	_ "private-conda-repo/indexer/shell"
)

var indMgr indexer.Indexer

func init() {
	_id, err := indexer.New()
	if err != nil {
		log.Fatal(err)
	}
	indMgr = _id
}
