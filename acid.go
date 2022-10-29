package nimbus

import (
	badger "github.com/dgraph-io/badger/v3"
)

func AddKV(db *badger.DB, key string, value string) error {
	txn := db.NewTransaction(true)
	defer txn.Discard()

	err := txn.Set([]byte(key), []byte(value))
	if err != nil {
		return err
	}

	if err := txn.Commit(); err != nil {
		return err
	}
	return nil

}

func ListKV(db *badger.DB) (map[string]string, error) {

	var values = make(map[string]string)

	err := db.View(func(txn *badger.Txn) error {
		opts := badger.DefaultIteratorOptions
		opts.PrefetchSize = 10
		it := txn.NewIterator(opts)
		defer it.Close()
		for it.Rewind(); it.Valid(); it.Next() {
			item := it.Item()
			k := item.Key()
			err := item.Value(func(v []byte) error {
				key := string(k[:])
				val := string(v[:])
				values[key] = val
				return nil
			})
			if err != nil {
				return err
			}
		}
		return nil
	})
	if err != nil {
		return nil, err
	}
	return values, nil
}

func ReadKV(db *badger.DB, key string) (string, error) {

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
			return err
		}
		return nil
	})
	if err != nil {
		return "", err
	}
	return value, nil
}

func RemoveKV(db *badger.DB, key string) error {

	err := db.Update(func(txn *badger.Txn) error {
		opts := badger.DefaultIteratorOptions
		opts.AllVersions = false

		// get key to validate it exists in database before delete
		_, err := txn.Get([]byte(key))
		if err != nil {
			return err
		}

		err = txn.Delete([]byte(key))
		if err != nil {
			return err
		}

		return nil
	})
	if err != nil {
		return err
	}
	return nil
}
