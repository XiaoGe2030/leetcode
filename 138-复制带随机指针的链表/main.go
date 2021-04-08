package main

func main() {
	node1 := &Node{Val: 1, Next: nil, Random: node7}
	node10 := &Node{Val: 10, Next: node1, Random: node11}
	node11 := &Node{Val: 11, Next: node10, Random: node1}
	node13 := &Node{Val: 13, Next: node11, Random: node7}
	node7 := &Node{Val: 7, Next: node13, Random: nil}
}

type Node struct {
	Val    int
	Next   *Node
	Random *Node
}

func copyRandomList(head *Node) *Node {
	p := head

	resMap := make(map[*Node]*Node, 0)
	for head != nil {
		nNode := &Node{Val: head.Val}
		resMap[head] = nNode
		head = head.Next
	}

	res := p
	for p != nil {
		tmpNode := resMap[p]
		tmpNode.Next = resMap[p.Next]
		tmpNode.Random = resMap[p.Random]
		p = p.Next
	}

	return resMap[res]
}
