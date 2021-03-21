package main

import "fmt"

func main() {
	node5 := &ListNode{Val: 5, Next: nil}
	node4 := &ListNode{Val: 4, Next: node5}
	node3 := &ListNode{Val: 3, Next: node4}
	node2 := &ListNode{Val: 2, Next: node3}
	node1 := &ListNode{Val: 1, Next: node2}

	fmt.Println(getKthFromEnd(node1, 2))
}

type ListNode struct {
	Val  int
	Next *ListNode
}

func getKthFromEnd(head *ListNode, k int) *ListNode {
	var res []int

	for head != nil {
		res = append(res, head.Val)
		head = head.Next
		if len(res) > k {
			res = res[1:]
		}
	}

	p := &ListNode{}
	q = p
	for _, n := range res {
		node := &ListNode{Val: n, Next: nil}
		q.Next = node
		q = q.Next

	}
	return p.Next
}
