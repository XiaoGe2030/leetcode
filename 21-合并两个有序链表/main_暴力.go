package main

import (
	"fmt"
	"math"
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

	prepre := &ListNode{Val: math.MinInt8, Next: nil}
	pre := prepre

	for l1 != nil && l2 != nil {
		n1 := l1.Val
		n2 := l2.Val
		if n1 < n2 {
			pre.Next = l1
			pre = l1
			l1 = l1.Next
		} else {
			pre.Next = l2
			pre = l2
			l2 = l2.Next
		}
	}

	if l1 != nil {
		pre.Next = l1
	}

	if l2 != nil {
		pre.Next = l2
	}

	return prepre.Next
}
