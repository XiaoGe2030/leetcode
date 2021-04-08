package main

import "fmt"

func main() {
	node := ListNode{Val:1,Next:}
}

type ListNode struct {
	Val  int
	Next *ListNode
}

func hasCycle(head *ListNode) bool {
	
	if head == nil || head.Next == nil {
		return false
	}

	slow := head.Next
	fast := head.Next.Next

	for slow !=fast {
		if fast  == nil || fast.Next == nil {
			return false
		}

		slow = slow.Next
		fast= fast.Next.Next
	}
	return true 
}
