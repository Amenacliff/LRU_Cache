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

func TestCacheHitReshuffle(t *testing.T) {
	lruCacheInit := &LRUCache{}
	cache := lruCacheInit.InitializeCache(3)
	cache.Put("boy1", "AmenaCliff1")
	cache.Put("boy2", "AmenaCliff2")
	cache.Put("boy3", "AmenaCliff3")
	cache.Put("boy4", "AmenaCliff4")
	keyToCheck := "boy1"
	firstItemIn := cache.Get(keyToCheck)
	if firstItemIn != nil && firstItemIn.key != cache.head.key {
		t.Error("A cache hit is suppose to take the node looked up to the head of the cache's, DLL, it didn't")
	}
}

func TestLengthOfList(t *testing.T) {
	lruCacheInit := &LRUCache{}
	cache := lruCacheInit.InitializeCache(3)
	cache.Put("boy1", "AmenaCliff1")
	cache.Put("boy2", "AmenaCliff2")
	cache.Put("boy3", "AmenaCliff3")
	cache.Put("boy4", "AmenaCliff4")
	cache.Put("boy5", "AmenaCliff5")

	if cache.length > cache.maxLength {
		t.Error("The Length of the cache should not be more than, the max size of the cache ")
	}
}

func TestIncreaseCacheSize(t *testing.T) {
	lruCacheInit := &LRUCache{}
	initialSize := 10
	sizeToAdd := 3
	cache := lruCacheInit.InitializeCache(initialSize)
	cache.IncreaseCacheSize(sizeToAdd)

	if cache.maxLength != initialSize+sizeToAdd {
		t.Error("Increase Size Cache Function Failed Test")
	}

}
