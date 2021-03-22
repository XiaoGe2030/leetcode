package main

import "fmt"

func main() {
	fmt.Println(permute([]int{1, 2, 3}))
}

<<<<<<< HEAD
var result [][]int
var numBool []bool

func permute(nums []int) [][]int {
	tmp := make([]int, 0)
	numBool = make([]bool, len(nums))
	getNum(nums, 0, tmp)
	return result
}

func getNum(nums []int, index int, tmp []int) {

	if index == len(nums) {
		numTmp := make([]int, 0)
		numTmp = append(numTmp, tmp...)
		result = append(result, numTmp)
		return
	}

	for k, v := range nums {
		if !numBool[k] {
			tmp = append(tmp, v)
			numBool[k] = true
			getNum(nums, index+1, tmp)
			tmp = tmp[:len(tmp)-1]
			numBool[k] = false
		}
	}
	return
=======
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
>>>>>>> 20400e72f6b5c14362dc03729df340c899d1cd9f
}
