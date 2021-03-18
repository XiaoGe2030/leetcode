package main

import "fmt"

func main() {
	// node5 := &ListNode{Val: 5, Next: nil}
	//node4 := &ListNode{Val: 4, Next: node5}
	//node3 := &ListNode{Val: 3, Next: node4}
	//node2 := &ListNode{Val: 2, Next: node3}
	//node1 := &ListNode{Val: 1, Next: node2}
	//reorderList(node1)

	//for node1 != nil {
	//fmt.Println(node1.Val)
	//node1 = node1.Next
	//}

	node14 := &ListNode{Val: 4, Next: nil}
	node13 := &ListNode{Val: 3, Next: node14}
	node12 := &ListNode{Val: 2, Next: node13}
	node11 := &ListNode{Val: 1, Next: node12}
	reorderList(node11)

	for node11 != nil {
		fmt.Println(node11.Val)
		node11 = node11.Next
	}
}

type ListNode struct {
	Val  int
	Next *ListNode
}

func reorderList(head *ListNode) {
	var res []int

	p := head
	for head != nil {
		res = append(res, head.Val)
		head = head.Next
	}

	n := len(res)

	p1 := p
	i, j := 1, n-1
	for ; i < j; i, j = i+1, j-1 {
		nodei := &ListNode{Val: res[i], Next: nil}
		nodej := &ListNode{Val: res[j], Next: nil}

		p1.Next = nodej
		nodej.Next = nodei
		p1 = nodei
	}

	if n&1 == 0 {
		nodei := &ListNode{Val: res[i], Next: nil}
		p1.Next = nodei
	}
	head = p
}
