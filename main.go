package main

import (
	"log"
)

type Node struct {
	key  string
	data string
	prev *Node
	next *Node
}

type LRUCache struct {
	head      *Node
	tail      *Node
	maxLength int
	length    int
	cacheMap  map[string]*Node
}

func (cache *LRUCache) InitializeCache(maxSize int) *LRUCache {
	newLruCache := LRUCache{
		maxLength: maxSize,
		length:    0,
		cacheMap:  make(map[string]*Node),
	}

	return &newLruCache
}

func (cache *LRUCache) Put(key string, data string) {
	newCacheItem := &Node{
		key:  key,
		data: data,
	}
	if cache.head == nil {
		cache.head = newCacheItem
		cache.tail = newCacheItem
	} else {
		newCacheItem = cache.AddItemToFront(newCacheItem)
	}
	cacheMap := cache.cacheMap
	cacheMap[key] = newCacheItem
	cache.cacheMap = cacheMap

	if cache.length == cache.maxLength {
		cache.RemoveLastCacheItem()
	}

	cache.length = cache.length + 1

}

func (cache *LRUCache) RemoveItem(cacheItem *Node) {

	if cache.head == nil {
		return
	}

	if cache.head == cacheItem {
		cache.head = cacheItem.next
	}

	if cacheItem.prev != nil {
		previousItem := cacheItem.prev
		previousItem.next = cacheItem.next
		if cacheItem.next != nil {
			nextItem := cacheItem.next
			nextItem.prev = previousItem
		}
	}

}

func (cache *LRUCache) RemoveLastCacheItem() {
	lastItem := cache.tail
	if lastItem != nil {
		previousItem := lastItem.prev
		cache.tail = previousItem
		previousItem.next = nil
		delete(cache.cacheMap, lastItem.key)
	}
	cache.length = cache.length - 1
}

func (cache *LRUCache) AddItemToFront(cacheItem *Node) *Node {
	if cache.head == nil {
		cache.head = cacheItem
		cache.tail = cacheItem
		cacheItem.prev = nil
		cacheItem.next = nil
	}

	if cache.head != nil {
		currentHead := cache.head
		cacheItem.next = currentHead
		cacheItem.prev = nil
		currentHead.prev = cacheItem
		cache.head = cacheItem
	}

	return cacheItem

}

func (cache *LRUCache) Get(key string) *Node {
	if value, ok := cache.cacheMap[key]; ok {
		if value != cache.head {
			cache.RemoveItem(value)
			cache.AddItemToFront(value)
		}
		return value
	} else {
		return nil
	}
}

func main() {
	lruCache := LRUCache{}
	cache := lruCache.InitializeCache(6)
	cache.Put("boy", "Kevwe")
	cache.Put("boy1", "Kevwe1")
	cache.Put("boy2", "Kevwe2")
	cache.Put("boy3", "Kevwe3")
	cache.Put("boy4", "Kevwe4")
	cache.Put("boy5", "Kevwe5")
	cache.Put("boy6", "Kevwe6")
	cache.Put("boy7", "Amenannnnaa")
	log.Println(cache.head.key, cache.tail.key, cache.cacheMap)
}
