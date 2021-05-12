package main

import "fmt"

func main() {
	fmt.Println("vim-go")
}

type ListNode struct {
	Val  int
	Next *ListNode
}

func getIntersectionNode(headA, headB *ListNode) *ListNode {
	nodeMap := make(map[*ListNode]*ListNode)
	pA := headA
	pB := headB
	for pA != nil {
		nodeMap[pA] = pA
		pA = pA.Next
	}

	for pB != nil {
		if v, ok := nodeMap[pB]; ok {
			return v
		}
	}

	return nil
}
