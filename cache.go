package coc

import (
	"errors"
	"sync"
	"time"
)

var (
	useCache = true

	cache     = make(map[string][]byte)
	cacheLock sync.RWMutex
)

// Caching is enabled by default, use this to disable it
func UseCache(b bool) {
	useCache = b
}

// getFromCache returns the data from the cache map if it exists, otherwise nil
func getFromCache(key string) ([]byte, error) {
	cacheLock.RLock()
	data, ok := cache[key]
	cacheLock.RUnlock()
	if ok {
		return data, nil
	}
	return nil, errors.New("not in cache")
}

// writeToCache writes the data to the cache map
func writeToCache(key string, data []byte, duration int) error {
	cacheLock.Lock()
	cache[key] = data
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
