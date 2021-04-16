package main

import "fmt"

func main() {

	node7 := &ListNode{Val: 5, Next: nil}
	node8 := &ListNode{Val: 3, Next: node7}
	l3 := reverseBetween(node8, 1, 2)
	var res3 []int
	for l3 != nil {
		res3 = append(res3, l3.Val)
		l3 = l3.Next
	}

	fmt.Println(res3)

}

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseBetween(head *ListNode, left int, right int) *ListNode {
	var pre *ListNode
	for index, cur := 1, head; cur != nil; pre, cur, index = cur, cur.Next, index+1 {
		if index == left {
			tmpHead, nextL := reverseList(cur, right-left+1)
			if pre != nil {
				pre.Next = tmpHead
			} else {
				head = tmpHead
			}
			cur.Next = nextL
			break
		}
	}
	return head
}

func reverseList(root *ListNode, num int) (*ListNode, *ListNode) {
	var pre, nextL *ListNode
	cur := root
	for cur != nil && num > 0 {
		next := cur.Next
		nextL = next
		cur.Next = pre
		pre = cur
		cur = next
		num--
	}
	return pre, nextL
}
