package nimbus

import (
	"github.com/dgraph-io/badger/v3"
)

// type DBOpts struct {
// 	Path   string
// 	Logger string
// }

// func (db *DBOpts) InitDB() (*badger.DB, error) {
// 	opts := badger.DefaultOptions(db.Path)
// 	opts.Logger = nil

// 	db, err := badger.Open(opts)
// 	if err != nil {
// 		return nil, err
// 	}
// 	return db, nil
// }

func InitDB(path string) (*badger.DB, error) {

	opts := badger.DefaultOptions(path)
	opts.Logger = nil

	db, err := badger.Open(opts)
	if err != nil {
		return nil, err
	}

	return db, nil

}
