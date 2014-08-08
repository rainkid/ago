package cache

import (
	"fmt"
	"sync"
	"time"
)

//cache item
type Item struct {
	value   interface{}
	expired int64
}

type LItem struct {
	key     string
	expired int64
}

type Cache struct {
	lock  sync.RWMutex
	items map[string]Item
}

func NewCaches() *Cache {
	cache := &Cache{items: make(map[string]Item)}
	go cache.runtime()
	return cache
}

//set cache
func (cache *Cache) Set(key string, value interface{}, expired int64) {
	fmt.Println("set cache : ", key, expired)
	//lock
	cache.lock.Lock()
	defer cache.lock.Unlock()

	expired += time.Now().Unix()
	cache.items[key] = Item{expired: expired, value: value}
}

//get cache
func (cache *Cache) Get(key string) interface{} {
	fmt.Println("get cache : ", key)
	item, ok := cache.items[key]
	if !ok {
		return nil
	}
	now := time.Now().Unix()

	//check item if expired and delete
	if item.expired <= now {
		cache.Delete(key)
		return nil
	}
	return item.value
}

//delete cache with key
func (cache *Cache) Delete(key string) {
	fmt.Println("delete cache : ", key)
	cache.lock.Lock()
	defer cache.lock.Unlock()
	//delete key
	delete(cache.items, key)

}

func (cache *Cache) runtime() {
	duration, _ := time.ParseDuration(fmt.Sprintf("%ds", 1))
	for {
		<-time.After(duration)
		cache.expired()
	}
}

//
func (cache *Cache) expired() {
	now := time.Now().Unix()
	for name, item := range cache.items {
		if item.expired <= now {
			cache.Delete(name)
		}
	}
}
