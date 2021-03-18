package main

import "fmt"

func main() {
	node3 := &ListNode{Val: 3, Next: nil}
	node1 := &ListNode{Val: 1, Next: node3}
	node2 := &ListNode{Val: 2, Next: node1}
	node4 := &ListNode{Val: 4, Next: node2}

	insertionSortList(node4)
}

type ListNode struct {
	Val  int
	Next *ListNode
}

func insertionSortList(head *ListNode) *ListNode {

	//利用数组接收链表中的元素，进入数组前找到合适的位置点
	var res []int
	if head != nil {
		res = append(res, head.Val)
		head = head.Next
	}

	fmt.Println(res)
	for head != nil {
		n := head.Val
		var tmp []int
		for i, v := range res {
			if n < v {
				tmp = append(tmp, res[0:i]...)
				tmp = append(tmp, n)
				tmp = append(tmp, res[i:]...)
			}
		}
		if len(tmp) <= len(res) {
			res = append(res, n)
		} else {
			res = tmp
		}

		head = head.Next
	}

	fmt.Println(res)

	return nil
}
