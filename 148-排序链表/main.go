package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println("vim-go")
}

type ListNode struct {
	Val  int
	Next *ListNode
}

func sortList(head *ListNode) *ListNode {
	var res []int
	for head != nil {
		res = append(res, head.Val)
		head = head.Next
	}

	sort.Ints(res)
	pre := &ListNode{}
	q := pre
	for _, v := range res {
		node := &ListNode{Val: v, Next: nil}
		pre.Next = node
		pre = pre.Next
	}

	return q.Next
}
