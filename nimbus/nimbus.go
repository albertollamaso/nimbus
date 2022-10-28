package nimbus

import (
	"log"

	badger "github.com/dgraph-io/badger/v3"
)

func InitDB(path string) *badger.DB {

	opts := badger.DefaultOptions(path)
	opts.Logger = nil

	db, err := badger.Open(opts)
	if err != nil {
		log.Fatal(err)
	}

	return db

}

func AddKV(db *badger.DB, key string, value string) {
	txn := db.NewTransaction(true)
	defer txn.Discard()

	err := txn.Set([]byte(key), []byte(value))
	if err != nil {
		log.Fatal(err)
	}

	if err := txn.Commit(); err != nil {
		log.Fatal(err)
	}

}

func ReadKV(db *badger.DB, key string) string {

	var value string

	err := db.View(func(txn *badger.Txn) error {
		item, err := txn.Get([]byte(key))
		if err != nil {
			return err
		}

		err = item.Value(func(val []byte) error {

			value = string(val[:])

			return nil
		})
		if err != nil {
			log.Fatal(err)
		}
		return nil
	})
	if err != nil {
		log.Fatal(err)
	}
	return value
}
