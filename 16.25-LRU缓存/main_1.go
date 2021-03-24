package main

func main() {
}

/*
	解题关键，
	1、数据结构map[key]node, node节点是双向链表的节点，且节点的值为，key ,value
	2、实现addToHead、removeNode、removeTail、moveToHead方法
*/
type LRUCache struct {
	size     int
	capCache int
	data     map[int]*LRUNode
	head     *LRUNode
	tail     *LRUNode
}

type LRUNode struct {
	key   int
	value int
	pre   *LRUNode
	next  *LRUNode
}

func Constructor(capacity int) LRUCache {
	lru := LRUCache{}

	lru.data = make(map[int]*LRUNode)
	lru.capCache = capacity
	lru.head = &LRUNode{}
	lru.tail = &LRUNode{}

	lru.head.next = lru.tail
	lru.tail.pre = lru.head

	return lru
}

func (this *LRUCache) Get(key int) int {
	if v, ok := this.data[key]; ok {
		//key对应的节点移动到双向链表头
		this.moveToHead(v)
		return v.value
	} else {
		return -1
	}
}

func (this *LRUCache) Put(key int, value int) {
	if v, ok := this.data[key]; ok {
		v.Value = value
		this.moveToHead(v)
	} else {
		tmp := &LRUNode{value: value, key: key}
		this.data[key] = tmp
		size++
		this.addToHead(tmp)
		if this.size > this.capCache {
			node := this.removeTail()
			delete(this.data, node.key)
			this.size--
		}
	}
}

func (this *LRUCache) addToHead(node *LRUNode) *LRUNode {
	node.next = this.head.next
	node.pre = this.head
	this.head.next.pre = node
	thi.head.next = node
}
func (this *LRUCache) removeNode(node *LRUNode) {
	node.pre.next = node.next
	node.next.pre = node.pre
}

func (this *LRUCache) moveToHead(node *LRUNode) {
	this.removeNode(node)
	this.addToHead(node)
}
func (this *LRUCache) removeTail() LRUNode {
	node := this.tail.pre
	this.removeNode(node)
	return node
}
