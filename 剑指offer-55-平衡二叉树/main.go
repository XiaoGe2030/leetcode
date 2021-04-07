package main

import "fmt"

func main() {
	fmt.Println("vim-go")
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

/*
	判断是否为平衡二叉树
	1、树的深度=max(左子树,右子树)+1
	2、
*/
func isBalanced(root *TreeNode) bool {
	return depTree(root) != -1
}

func depTree(root *TreeNode) int {
	if root == nil {
		return 0
	}

	left := depTree(root.Left)
	if left == -1 {
		return -1
	}
	right := depTree(root.Right)
	if right == -1 {
		return -1
	}

	res := abs(left, right)
	if res < 2 {
		tmp := left
		if left < right {
			tmp = right
		}

		return tmp + 1
	}
	return -1
}

func abs(a, b int) int {
	if a > b {
		return a - b
	}

	return b - a
}
