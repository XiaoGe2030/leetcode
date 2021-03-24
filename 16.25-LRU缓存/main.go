package main

import "fmt"

func main() {
	fmt.Println("vim-go")
}

type LRUCache struct {
	data      map[int]int
	index_key []int
	lurCap    int
}

func Constructor(capacity int) LRUCache {
	lru := LRUCache{}
	lru.data = make(map[int]int, capacity)
	lru.index_key = make([]int, 0, capacity)
	lru.lurCap = capacity
	return lru
}

func (this *LRUCache) Get(key int) int {
	if v, ok := this.data[key]; ok {
		//查询到的数据需要往前移动
		kIndex := findIndex(this.index_key, key)
		tmp := this.index_key[:kIndex]
		tmp = append(tmp, this.index_key[kIndex+1:]...)
		tmp = append(tmp, key)

		this.index_key = tmp
		return v
	} else {
		return -1
	}
}

func (this *LRUCache) Put(key int, value int) {
	if _, ok := this.data[key]; !ok {
		if len(this.index_key) == this.lurCap {
			delete(this.data, this.index_key[0])
			this.index_key = this.index_key[1:]
			this.index_key = append(this.index_key, key)
			this.data[key] = value
		} else {
			this.index_key = append(this.index_key, key)
			this.data[key] = value
		}
		return
	}
	//如果存在，也是要移动到slice的尾部
	kIndex := findIndex(this.index_key, key)
	tmp := this.index_key[:kIndex]
	tmp = append(tmp, this.index_key[kIndex+1:]...)
	tmp = append(tmp, key)

	this.index_key = tmp
}

func findIndex(nums []int, n int) int {
	for i := 0; i < len(nums); i++ {
		if n == nums[i] {
			return i
		}
	}
	return -1
}
