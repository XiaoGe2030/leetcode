package main

import (
	"fmt"
)

func main() {
	node4 := &ListNode{Val: 4, Next: nil}
	node2 := &ListNode{Val: 2, Next: node4}
	node1 := &ListNode{Val: 1, Next: node2}

	node42 := &ListNode{Val: 4, Next: nil}
	node32 := &ListNode{Val: 3, Next: node42}
	node12 := &ListNode{Val: 1, Next: node32}

	p := mergeTwoLists(node12, node1)

	for p != nil {
		fmt.Println(p.Val)
		p = p.Next
	}
}

type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 == nil {
		return l2
	}

	if l2 == nil {
		return l1
	}

	if l1.Val < l2.Val {
		l1.Next = mergeTwoLists(l1.Next, l2)
		return l1
	} else {
		l2.Next = mergeTwoLists(l1, l2.Next)
		return l2
	}
}
