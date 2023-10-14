package lrucache

import (
	"testing"
)

func TestLRU(t *testing.T) {
	cache := NewLRUCache(3)
	cache.Set(&CacheElem{
		Key: "1",
		Val: "foo",
	})
	cache.Set(&CacheElem{
		Key: "2",
		Val: "bar",
	})
	cache.Set(&CacheElem{
		Key: "3",
		Val: "foobar",
	})
	cache.Print()
	cache.Get("2")
	cache.Set(&CacheElem{
		Key: "4",
		Val: "foobar",
	})
	cache.Print()
}
