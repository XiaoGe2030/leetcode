package main

func main() {
	node7 := &TreeNode{Val: 7, Left: nil, Right: nil}
	node15 := &TreeNode{Val: 15, Left: nil, Right: nil}
	node20 := &TreeNode{Val: 20, Left: node15, Right: node7}
	node9 := &TreeNode{Val: 9, Left: nil, Right: nil}
	node3 := &TreeNode{Val: 3, Left: node9, Right: node20}

	levelOrder(node3)
}

///var res []int

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func levelOrder(root *TreeNode) []int {
	var res []int
	var resNode []*TreeNode
	if root != nil {
		resNode = append(resNode, root)
	}

	n := len(resNode)

	for n > 0 {
		for _, node := range resNode {
			res = append(res, node.Val)
			if node.Left != nil {
				resNode = append(resNode, node.Left)
			}

			if node.Right != nil {
				resNode = append(resNode, node.Right)
			}
		}
		resNode = resNode[n:]
		n = len(resNode)
	}
	return res
}
