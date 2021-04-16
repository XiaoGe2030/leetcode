package main

import "fmt"

func main() {
	node7 := &TreeNode{Val: 7, Left: nil, Right: nil}
	node15 := &TreeNode{Val: 15, Left: nil, Right: nil}
	node20 := &TreeNode{Val: 20, Left: node15, Right: node7}
	node9 := &TreeNode{Val: 9, Left: nil, Right: nil}
	node3 := &TreeNode{Val: 3, Left: node9, Right: node20}

	fmt.Println(zigzagLevelOrder(node3))
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func zigzagLevelOrder(root *TreeNode) [][]int {
	var res [][]int
	var resNode []*TreeNode
	if root != nil {
		resNode = append(resNode, root)
	}

	n := len(resNode)
	isOdd := true
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

		if !isOdd {
			for i, j := 0, len(tmp)-1; i < j; i, j = i+1, j-1 {
				tmp[i], tmp[j] = tmp[j], tmp[i]
			}
			isOdd = true
		} else {
			isOdd = false
		}

		res = append(res, tmp)
		resNode = resNode[n:]
		n = len(resNode)
	}

	return res
}
