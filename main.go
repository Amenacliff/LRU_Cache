package LRU_Cache

type Node struct {
	key  string
	data interface{}
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
	}
	cacheMap := cache.cacheMap
	cacheMap[key] = newCacheItem
	cache.cacheMap = cacheMap
}

func (cache *LRUCache) RemoveItem(cacheItem *Node) {

	if cache.head == nil || cacheItem == nil {
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

func (cache *LRUCache) Get(key string) interface{} {
	if value, ok := cache.cacheMap[key]; ok {
		return value.data
	} else {
		return -1
	}
}
