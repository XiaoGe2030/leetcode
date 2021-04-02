package main

import "fmt"

func main() {
	fmt.Println("vim-go")
}

type ListNode struct {
	Val  int
	Next *ListNode
}

/*
	1、O(n)时间复杂度
	2、O(1)的空间复杂度

	解决思路
	1、利用快慢指针找到链表的中间位置，将链表分为两部分
	2、反转后面部分
	3、对比前面部分和反转后的后面部分
	4、再次反转反转后的后面部分链表
*/

func reverseList(node *ListNode) *ListNode {

	var pre *ListNode
	cur := node
	for cur != nil {
		next := cur.Next
		cur.Next = pre
		pre = cur
		cur = next

	}
	return pre
}

func endofFirstHalf(head *ListNode) *ListNode {
	fast, slow := head, head
	for fast.Next != nil && fast.Next.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}
	return slow

}
func isPalindrome(head *ListNode) bool {
	if head == nil {
		return false
	}

	firstHalfEnd := endofFirstHalf(head)
	secondHalfStart := reverseList(firstHalfEnd.Next)

	//判断是否回文
	p1 := head
	p2 := secondHalfStart
	result := true

	for result && p2 != nil {
		if p1.Val != p2.Val {
			return false
		}

		p1 = p1.Next
		p2 = p2.Next
	}

	firstHalfEnd.Next = reverseList(secondHalfStart)
	return result
}
