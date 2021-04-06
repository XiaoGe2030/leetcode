package main

import "fmt"

func main() {
	fmt.Println("vim-go")
}

func levelOrder(root *TreeNode) [][]int {
	var res [][]int
	var resNode []*TreeNode

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
