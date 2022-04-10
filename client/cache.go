package client

import (
	"bytes"
	"errors"
	"io/ioutil"
	"sync"
	"time"

	"github.com/andybalholm/brotli"
	// badger "github.com/dgraph-io/badger/v3"
) // fuck badger

var (
	cache = make(map[string][]byte)
	cacheLock sync.RWMutex
)

// func init() {
// 	opt := badger.DefaultOptions("").WithInMemory(true)
// 	opt.Logger = nil
// 	_cache, err := badger.Open(opt)
// 	if err != nil {
// 		panic(err)
// 	}
// 	cache = _cache
// }

func getFromCache(key string) ([]byte, error) {
	cacheLock.RLock()
	data, ok := cache[key]
	cacheLock.RUnlock()
	if ok {
		data, err := readFromBrotli(data)
		return data, err
	}
	return nil, errors.New("not in cache")
}

func writeToCache(key string, data []byte, duration int) error {
	brot, err := writeToBrotli(data)
	if err != nil {
		return err 
	}
	cacheLock.Lock()
	cache[key] = brot
	cacheLock.Unlock()
	if duration > 0 {
		go func() {
			time.Sleep(time.Second * time.Duration(duration))
			cacheLock.Lock()
			delete(cache, key)
			cacheLock.Unlock()
		}()
	}
	return nil
}

func writeToBrotli(data []byte) ([]byte, error) {
	buf := bytes.NewBuffer(nil)
	writer := brotli.NewWriter(buf)
	_, err := writer.Write(data)
	if err != nil {
		return nil, err
	}
	writer.Close()
	return buf.Bytes(), nil
}

func readFromBrotli(brotData []byte) ([]byte, error) {
	buf := bytes.NewReader(brotData)
	reader := brotli.NewReader(buf)
	return ioutil.ReadAll(reader)
}


// func getFromCache(url string) ([]byte, error) {
// 	var valCopy []byte
// 	err := cache.View(func(txn *badger.Txn) error {
// 		item, err := txn.Get([]byte(url))
// 		if err != nil {
// 			return err
// 		}

// 		valCopy, err = item.ValueCopy(nil)
// 		if err != nil {
// 			return err
// 		}

// 		return nil
// 	})
// 	if err != nil {
// 		return nil, err
// 	}
// 	return valCopy, nil
// }

// func writeToCache(url string, data []byte, duration int) error {
// 	err := cache.Update(func(txn *badger.Txn) error {
// 		entry := badger.NewEntry([]byte(url), data).WithTTL(time.Second * time.Duration(duration))
// 		return txn.SetEntry(entry)
// 	})
// 	return err
// }
