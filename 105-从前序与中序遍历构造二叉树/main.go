package main

import "fmt"

func main() {
	node := buildTree([]int{3, 9, 20, 15, 7}, []int{9, 3, 15, 20, 7})
	fmt.Println(printTree(node))
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func buildTree(preorder []int, inorder []int) *TreeNode {
	resMap := make(map[int]int)
	for i, v := range inorder {
		resMap[v] = i
	}
	return recur(0, 0, len(inorder)-1, preorder, resMap)
}

func recur(root, left, right int, preorder []int, resMap map[int]int) *TreeNode {
	if left > right {
		return nil
	}
	node := &TreeNode{Val: preorder[root]}
	index := resMap[preorder[root]]
	node.Left = recur(root+1, left, index-1, preorder, resMap)
	node.Right = recur(root+index-left+1, index+1, right, preorder, resMap)

	return node
}

func printTree(root *TreeNode) [][]int {
	var resNode []*TreeNode
	var res [][]int
	if root != nil {
		resNode = append(resNode, root)
	}

	n := len(resNode)
	for n > 0 {
		var tmp []int
		for _, node := range resNode {
			tmp = append(tmp, node.Val)
			if node.Left != nil {
				resNode = append(resNode, node.Left)
			}
			if node.Right != nil {
				resNode = append(resNode, node.Right)
			}
		}
		res = append(res, tmp)
		resNode = resNode[n:]
		n = len(resNode)
	}
	return res
}
