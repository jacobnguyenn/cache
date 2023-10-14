package lrucache

type CacheElem struct {
	Key string
	Val any
}
type ICache interface {
	Set(e *CacheElem)
	Get(key string) *CacheElem
	Print()
}
