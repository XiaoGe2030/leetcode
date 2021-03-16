package main

import "fmt"

func main() {
	node17 := &ListNode{Val: 9, Next: nil}
	node16 := &ListNode{Val: 9, Next: node17}
	node15 := &ListNode{Val: 9, Next: node16}
	node14 := &ListNode{Val: 9, Next: node15}
	node13 := &ListNode{Val: 9, Next: node14}
	node12 := &ListNode{Val: 9, Next: node13}
	node11 := &ListNode{Val: 9, Next: node12}

	node24 := &ListNode{Val: 9, Next: nil}
	node23 := &ListNode{Val: 9, Next: node24}
	node22 := &ListNode{Val: 9, Next: node23}
	node21 := &ListNode{Val: 9, Next: node22}

	fmt.Println(addTwoNumbers(node11, node21))
}

type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {

	var result []int
	isJ := false
	for l1 != nil && l2 != nil {
		n1 := l1.Val
		n2 := l2.Val
		tmp := n1 + n2
		if isJ {
			tmp++
		}

		if tmp > 9 {
			isJ = true
			tmp -= 10
		} else {
			isJ = false
		}
		result = append(result, tmp)
		l1 = l1.Next
		l2 = l2.Next
	}

	for l1 != nil {
		tmp := l1.Val
		if isJ {
			tmp++
		}

		if tmp > 9 {
			isJ = true
			tmp -= 10
		} else {
			isJ = false
		}
		result = append(result, tmp)

		l1 = l1.Next
	}
	for l2 != nil {
		tmp := l2.Val
		if isJ {
			tmp++
		}

		if tmp > 9 {
			isJ = true
			tmp -= 10
		} else {
			isJ = false
		}
		result = append(result, tmp)
		l2 = l2.Next
	}

	if isJ {
		result = append(result, 1)
	}
	head := &ListNode{}
	p := head
	for _, n := range result {
		node := &ListNode{Val: n, Next: nil}
		p.Next = node
		p = p.Next
	}
	fmt.Println(result)
	return head.Next
}
