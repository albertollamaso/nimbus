package nimbus

import (
	"github.com/dgraph-io/badger/v3"
)

func InitDB(path string) (*badger.DB, error) {

	opts := badger.DefaultOptions(path)
	opts.Logger = nil

	db, err := badger.Open(opts)
	if err != nil {
		return nil, err
	}

	return db, nil
}
