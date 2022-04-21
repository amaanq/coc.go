package coc

import (
	"errors"
	"sync"
	"time"
)

var (
	cache     = make(map[string][]byte)
	cacheLock sync.RWMutex
)

func getFromCache(key string) ([]byte, error) {
	cacheLock.RLock()
	data, ok := cache[key]
	cacheLock.RUnlock()
	if ok {
		return data, nil
	}
	return nil, errors.New("not in cache")
}

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