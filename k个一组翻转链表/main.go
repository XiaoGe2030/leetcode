package main

import "fmt"

func main() {
	node5 := &ListNode{Val: 5, Next: nil}
	node4 := &ListNode{Val: 4, Next: node5}
	node3 := &ListNode{Val: 3, Next: node4}
	node2 := &ListNode{Val: 2, Next: node3}
	node1 := &ListNode{Val: 1, Next: node2}

	fmt.Println(reverseKGroup(node1, 3))
}

//给你一个链表，每 k 个节点一组进行翻转，请你返回翻转后的链表。

//k 是一个正整数，它的值小于或等于链表的长度。

//如果节点总数不是 k 的整数倍，那么请将最后剩余的节点保持原有顺序。

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseKGroup(head *ListNode, k int) *ListNode {
	var res []int
	for head != nil {
		res = append(res, head.Val)
		head = head.Next
	}

	n := len(res)
	t := n / k //t = 1
	for i := 0; i < t*k; i = i + k {
		fmt.Println(res[i : i+k])
		for m, n := i, i+k-1; m < n; m, n = m+1, n-1 {
			res[m], res[n] = res[n], res[m]
		}
		fmt.Println(res[i : i+k])
	}
	fmt.Println(res)
	phead := &ListNode{}
	p := phead

	for _, n := range res {
		node := &ListNode{Val: n, Next: nil}
		p.Next = node
		p = p.Next
	}

	return phead.Next
}
