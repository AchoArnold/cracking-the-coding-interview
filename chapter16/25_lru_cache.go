package chapter16

import (
	"container/list"
	"sync"

	"github.com/pkg/errors"
)

// LRUCache is the data structure for an LRU cache
type LRUCache struct {
	maxCacheSize  int
	lock          sync.Mutex
	cacheMap      map[int]*list.Element
	frequencyList *list.List
}

type freqListNode struct {
	key   int
	value string
}

var (
	errCacheMiss = errors.New("cache miss")
)

// NewLRUCache creates a new LRU cache
func NewLRUCache(size int) *LRUCache {
	return &LRUCache{
		maxCacheSize:  size,
		lock:          sync.Mutex{},
		cacheMap:      map[int]*list.Element{},
		frequencyList: list.New(),
	}
}

// GetValue returns the value from the cache
func (cache *LRUCache) GetValue(key int) (val string, err error) {
	item, ok := cache.cacheMap[key]
	if !ok {
		return val, errCacheMiss
	}

	cache.lock.Lock()
	// move to front of the list to mark as most recently used
	if item != cache.frequencyList.Front() {
		cache.removeFromFrequencyList(item)
		cache.insertAtFrontOfFrequencyList(item.Value.(*freqListNode))
	}
	cache.lock.Unlock()

	return item.Value.(*freqListNode).value, err
}

func (cache *LRUCache) removeKey(key int) {
	element, ok := cache.cacheMap[key]
	if !ok {
		return
	}

	cache.removeFromFrequencyList(element)
	delete(cache.cacheMap, key)
}

// SetKeyValue sets the value into the front of the cache.
func (cache *LRUCache) SetKeyValue(key int, value string) {
	cache.lock.Lock()

	// if key is already present, remove it
	cache.removeKey(key)

	// if full, remove least recently used item from cache
	if len(cache.cacheMap) >= cache.maxCacheSize && cache.frequencyList.Back() != nil {
		cache.removeKey(cache.frequencyList.Back().Value.(*freqListNode).key)
	}

	// insert new node
	node := &freqListNode{key, value}
	element := cache.insertAtFrontOfFrequencyList(node)
	cache.cacheMap[key] = element

	cache.lock.Unlock()
}

func (cache *LRUCache) removeFromFrequencyList(item *list.Element) {
	cache.frequencyList.Remove(item)
}

func (cache *LRUCache) insertAtFrontOfFrequencyList(val *freqListNode) *list.Element {
	return cache.frequencyList.PushFront(val)
}
