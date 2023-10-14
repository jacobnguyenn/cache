package lrucache

import (
	"sync"

	log "github.com/sirupsen/logrus"
)

type LRUCache struct {
	store   DoublyLinkedList
	lock    sync.Mutex
	hashMap map[string]*ListNode
	size    int
}

func NewLRUCache(size int) ICache {
	return &LRUCache{
		store:   DoublyLinkedList{},
		hashMap: make(map[string]*ListNode),
		size:    size,
	}
}

func (c *LRUCache) Get(key string) *CacheElem {
	c.lock.Lock()
	defer c.lock.Unlock()
	log.WithFields(log.Fields{
		"key": key,
	}).Info("Getting from cache")
	if _, exist := c.hashMap[key]; !exist {
		log.WithFields(log.Fields{
			"key": key,
		}).Info("Not exist in cache")
		return nil
	}
	node := c.hashMap[key]
	c.store.MoveToEnd(node)
	return c.hashMap[key].Val.(*CacheElem)
}

func (c *LRUCache) Set(elem *CacheElem) {
	c.lock.Lock()
	defer c.lock.Unlock()
	node := &ListNode{Val: elem}
	if _, exist := c.hashMap[elem.Key]; exist {
		log.WithFields(log.Fields{
			"key": elem.Key,
		}).Info("Existed in cache")
		c.store.MoveToEnd(node)
		return
	}
	log.WithFields(log.Fields{
		"key": elem.Key,
	}).Info("Not exist in cache")
	if c.store.Size() == c.size {
		log.Info("Cache size reached, evicting")
		c.Evict()
		delete(c.hashMap, elem.Key)
	}
	c.store.PushBack(node)
	c.hashMap[elem.Key] = node
}

func (c *LRUCache) Evict() {
	c.store.PopHead()
	c.size = c.store.Size()
}

func (c *LRUCache) Print() {
	c.store.Print()
}
