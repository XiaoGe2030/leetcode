package main

import "fmt"

func main() {

	node3 := &TreeNode{Val: 3, Left: nil, Right: nil}
	node2 := &TreeNode{Val: 2, Left: nil, Right: node3}
	node5 := &TreeNode{Val: 5, Left: nil, Right: nil}
	node4 := &TreeNode{Val: 4, Left: node2, Right: node5}
	fmt.Println(kthSmallest(node4, 1))
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

var res []int

func kthSmallest(root *TreeNode, k int) int {
	qianXiang(root)
	fmt.Println(res)
	return res[k-1]
}

func qianXiang(root *TreeNode) {
	if root == nil {
		return
	}
	if root.Left != nil {
		qianXiang(root.Left)
	}
	res = append(res, root.Val)

	if root.Right != nil {
		qianXiang(root.Right)
	}
}

var num int

func kthSmallest(root *TreeNode, k int) int {
	dfs(root, &k)
	return num
}

func dfs(root *TreeNode, k *int) {
	if root != nil {
		dfs(root.Left, &k)
		*k--
		if *k == 0 {
			num = root.Val
		}
		dfs(root.Right, &k)
	}
}
