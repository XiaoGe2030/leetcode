package main

import "fmt"

func main() {
	lru := Constructor(2)
	lru.Put(1, 1)
	lru.Put(2, 2)
	fmt.Println(lru.Get(1))
	lru.Put(3, 3)
	fmt.Println(lru.Get(2))
	lru.Put(4, 4)
	fmt.Println(lru.Get(1))
	fmt.Println(lru.Get(3))
	fmt.Println(lru.Get(4))

}

type LRUCache struct {
	Size   int
	CapLru int
	MapLru map[int]*LRUNode
	Head   *LRUNode
	Tail   *LRUNode
}

type LRUNode struct {
	KeyLru   int
	ValueLru int
	PreNode  *LRUNode
	NextNode *LRUNode
}

func Constructor(capacity int) LRUCache {
	lru := LRUCache{}
	lru.CapLru = capacity
	lru.Tail = &LRUNode{}
	lru.Head = &LRUNode{}
	lru.MapLru = make(map[int]*LRUNode, 0)
	lru.Head.NextNode = lru.Tail
	lru.Tail.PreNode = lru.Head
	return lru
}

func (this *LRUCache) Get(key int) int {
	if v, ok := this.MapLru[key]; ok {
		//把v移动到双向链表的头部
		this.moveToHead(v)
		return v.ValueLru
	} else {
		return -1
	}
}

func (this *LRUCache) Put(key int, value int) {
	if v, ok := this.MapLru[key]; ok {
		v.ValueLru = value
		this.moveToHead(v)
	} else {
		node := &LRUNode{KeyLru: key, ValueLru: value}
		this.MapLru[key] = node
		this.addToHead(node)
		this.Size++
		if this.Size > this.CapLru {
			node := this.removeTail()
			this.Size--
			delete(this.MapLru, node.KeyLru)
		}
	}
}

func (this *LRUCache) moveToHead(node *LRUNode) {
	this.removeNode(node)
	this.addToHead(node)
}

func (this *LRUCache) addToHead(node *LRUNode) {
	node.NextNode = this.Head.NextNode
	node.PreNode = this.Head

	this.Head.NextNode.PreNode = node
	this.Head.NextNode = node
}

func (this *LRUCache) removeNode(node *LRUNode) {
	node.PreNode.NextNode = node.NextNode
	node.NextNode.PreNode = node.PreNode
}

func (this *LRUCache) removeTail() *LRUNode {
	node := this.Tail.PreNode
	this.removeNode(node)
	return node
}
