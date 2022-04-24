package main

import (
	"testing"
)

func TestEvictionPolicy(t *testing.T) {
	lruCacheInit := &LRUCache{}
	cache := lruCacheInit.InitializeCache(3)
	cache.Put("boy1", "AmenaCliff1")
	cache.Put("boy2", "AmenaCliff2")
	cache.Put("boy3", "AmenaCliff3")
	cache.Put("boy4", "AmenaCliff4")

	firstItemIn := cache.Get("boy1")

	if firstItemIn != nil {
		t.Error("First Item to go in, wasn't evicted from the cache, after size of cache was exceeded")
	}

}