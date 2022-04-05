package client

import (
	"time"

	badger "github.com/dgraph-io/badger/v3"
)

var (
	cache *badger.DB
)

func init() {
	opt := badger.DefaultOptions("").WithInMemory(true)
	opt.Logger = nil
	_cache, err := badger.Open(opt)
	if err != nil {
		panic(err)
	}
	cache = _cache
}

func getFromCache(url string) ([]byte, error) {
	var valCopy []byte
	err := cache.View(func(txn *badger.Txn) error {
		item, err := txn.Get([]byte(url))
		if err != nil {
			return err
		}

		valCopy, err = item.ValueCopy(nil)
		if err != nil {
			return err
		}

		return nil
	})
	if err != nil {
		return nil, err
	}
	return valCopy, nil
}

func writeToCache(url string, data []byte, duration int) error {
	err := cache.Update(func(txn *badger.Txn) error {
		entry := badger.NewEntry([]byte(url), data).WithTTL(time.Second * time.Duration(duration))
		return txn.SetEntry(entry)
	})
	return err
}
