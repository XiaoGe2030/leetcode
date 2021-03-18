package main

import "fmt"

func main() {
	fmt.Println(permute([]int{1, 2, 3}))
}

//给定一个 没有重复 数字的序列，返回其所有可能的全排列。

var nBool []bool
var result [][]int

func permute(nums []int) [][]int {
	nBool = make([]bool, len(nums))
	tmp := make([]int, 0)
	dfs(nums, 0, tmp)
	return result
}

func dfs(nums []int, index int, res []int) {
	//终止条件
	if len(nums) == index {
		tmp := make([]int, 0)
		tmp = append(tmp, res...)
		result = append(result, tmp)
		return
	}

	for i := 0; i < len(nums); i++ {
		if !nBool[i] {
			res = append(res, nums[i])
			nBool[i] = true
			dfs(nums, index+1, res)
			res = res[:len(res)-1]
			nBool[i] = false
		}
	}
}
