package main

import "fmt"

func main() {
	node := buildTree([]int{3, 9, 20, 15, 7}, []int{9, 3, 15, 20, 7})
	printTreeNode(node)
	fmt.Println(res)
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

var res []int
var pre []int
var dic map[int]int

func buildTree(preorder []int, inorder []int) *TreeNode {

	dic = make(map[int]int)
	pre = preorder
	for i := 0; i < len(inorder); i++ {
		dic[inorder[i]] = i
	}

	return recur(0, 0, len(inorder)-1)
}

func recur(root, left, right int) *TreeNode {
	if left > right {
		return nil
	}

	node := &TreeNode{Val: pre[root]}
	i := dic[pre[root]]

	node.Left = recur(root+1, left, i-1)
	node.Right = recur(root+i-left+1, i+1, right)
	return node
}

func printTreeNode(root *TreeNode) {
	if root == nil {
		return
	}

	res = append(res, root.Val)
	printTreeNode(root.Left)
	printTreeNode(root.Right)

}
