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

}
