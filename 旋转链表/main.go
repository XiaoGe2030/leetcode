package main

import "fmt"

func main() {
	node5 := &ListNode{Val: 5, Next: nil}
	node4 := &ListNode{Val: 4, Next: node5}
	node3 := &ListNode{Val: 3, Next: node4}
	node2 := &ListNode{Val: 2, Next: node3}
	node1 := &ListNode{Val: 1, Next: node2}

	rotateRight(node1, 2)
}

type ListNode struct {
	Val  int
	Next *ListNode
}

func rotateRight(head *ListNode, k int) *ListNode {
	result := make([]int, 0)
	p := head
	for p != nil {
		result = append(result, p.Val)
		p = p.Next
	}
	fmt.Println(result)

	length := len(result)
	r := make([]int, length)
	for j := length - 1; j >= 0; j-- {
		r[(j+k)%length] = result[j]
	}

	fmt.Println(r)
	newHead := new(ListNode)
	p1 := newHead
	for i := 0; i < length; i++ {
		p1.Next = &ListNode{Val: r[i], Next: nil}
		p1 = p1.Next
	}
	return newHead.Next
}
