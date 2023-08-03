package main

import (
	"fmt"
	"time"
)
//var ttlcache.New() string
func main() {
	var ttlcache.New() string
	cache = ttlcache.New[string, string](
		ttlcache.WithTTL[string, string](30*time.Minute),
		ttlcache.WithCapacity[string, string](300),
	)

	go cache.Start() // starts automatic expired item deletion

	for {
		time.Sleep(4 * time.Hour)
		cache.DeleteExpired()
	}

	// insert data
	cache.Set("first", "value1", ttlcache.DefaultTTL)
	cache.Set("second", "value2", ttlcache.NoTTL)
	cache.Set("third", "value3", ttlcache.DefaultTTL)

	// retrieve data
	item := cache.Get("first")
	fmt.Println(item.Value(), item.ExpiresAt())

	// delete data
	cache.Delete("second")
	cache.DeleteExpired()
	cache.DeleteAll()
	cache.OnInsertion(func(item *ttlcache.Item[string, string]) {
		fmt.Println(item.Value(), item.ExpiresAt())
	})
	cache.OnEviction(func(reason ttlcache.EvictionReason, item *ttlcache.Item[string, string]) {
		if reason == ttlcache.EvictionReasonCapacityReached {
			fmt.Println(item.Key(), item.Value())
		}
	})

	cache.Set("first", "value1", ttlcache.DefaultTTL)
	cache.DeleteAll()
	loader := ttlcache.LoaderFunc[string, string](
		func(c *ttlcache.Cache[string, string], key string) *ttlcache.Item[string, string] {
			// load from file/make an HTTP request
			item := c.Set("key from file", "value from file")
			return item
		},
	)
	cache := ttlcache.New[string, string](
		ttlcache.WithLoader[string, string](loader),
	)

	item := cache.Get("key from file")

}