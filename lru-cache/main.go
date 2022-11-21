package main

type LRUCache struct {
	Name   string
	Keys   []int
	Values []int
}

func Constructor(capacity int) LRUCache {
	k := make([]int, capacity)
	v := make([]int, capacity)
	return LRUCache{
		Name:   "",
		Keys:   k,
		Values: v,
	}
}

func (this *LRUCache) Get(key int) int {
	if len(this.Keys) < key {
		return -1
	}
	return this.Values[key]
}

func (this *LRUCache) Put(key int, value int) {
	if len(this.Keys) <= key {

	}
}
