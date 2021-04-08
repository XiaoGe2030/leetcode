package main

import "fmt"

func main() {
	fmt.Println("vim-go")
}

type ListNode struct {
	Val  int
	Next *ListNode
}

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	mapNode := make(map[int]*ListNode)
	p := head
	var index int
	for head != nil {
		mapNode[index] = head
		head = head.Next
		index++
	}

	var pre, next *ListNode

	if n == index {
		p = p.Next
	} else if n == 1 {
		pre = mapNode[index-n-1]
		pre.Next = nil
	} else {
		pre = mapNode[index-n-1]
		next = mapNode[index-n+1]
		pre.Next = next
	}

	return p
}
