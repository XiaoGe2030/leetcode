package main

import (
	"fmt"
	"sort"
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

	var res []int
	for l1 != nil {
		res = append(res, l1.Val)
		l1 = l1.Next
	}

	for l2 != nil {
		res = append(res, l2.Val)
		l2 = l2.Next
	}

	sort.Ints(res)
	head := &ListNode{}
	p := head
	for _, n := range res {
		node := &ListNode{Val: n, Next: nil}
		p.Next = node
		p = p.Next
	}

	return head.Next
}
