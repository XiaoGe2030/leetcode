package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

//var resMap map[int][]int
var resSlice [][]int
var res []int

func main() {
	node5 := &TreeNode{Val: 5, Left: nil, Right: nil}
	node4 := &TreeNode{Val: 4, Left: nil, Right: nil}
	node3 := &TreeNode{Val: 3, Left: nil, Right: node4}
	node2 := &TreeNode{Val: 2, Left: nil, Right: node5}

	node1 := &TreeNode{Val: 1, Left: node2, Right: node3}

	rightSideView(node1)
	fmt.Println(res)
}

func rightSideView(root *TreeNode) []int {
	//resMap = make(map[int][]int)
	zhongXuBianli(root, 0)

	for i := 0; i < len(resSlice); i++ {
		tmp := resSlice[i]
		res = append(res, tmp[len(tmp)-1])
	}
	return res
}

func zhongXuBianli(root *TreeNode, depth int) {
	if root != nil {
		//resMap[depth] = append(resMap[depth], root.Val)
		if len(resSlice) < depth+1 {
			resSlice = append(resSlice, []int{})
		}
		resSlice[depth] = append(resSlice[depth], root.Val)
		zhongXuBianli(root.Left, depth+1)
		zhongXuBianli(root.Right, depth+1)
	}

}
