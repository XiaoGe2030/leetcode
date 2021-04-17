package main

import "fmt"

func main() {
	fmt.Println("vim-go")
}

type ListNode struct {
	Val  int
	Next *ListNode
}

func detectCycle(head *ListNode) *ListNode {
	f, s := head, head
	for {
		if f == nil || f.Next == nil {
			return nil
		}

		f = f.Next.Next
		s = s.Next
		if f == s {
			break
		}
	}

	f = head
	for f != s {
		f = f.Next
		s = s.Next
	}

	return s
}
