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
	
	res := make(map[*ListNode]int)
	for head != nil {
		if _,ok := res[head];ok {
			return true
		}else{
			res[head]=head.Val
		}
		head = head.Next
	}

	return false

}
