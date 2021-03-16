package main

import "fmt"

func main() {
	fmt.Println(maxSubArray2([]int{-2, 1, -3, 4, -1, 2, 1, -5, 4}))
}

/*
给定一个整数数组nums,找到一个具有最大和的连续子序列（子序列最少包含一个元素），返回其最大和。
*/

//这里是列出所有的子序列，然后获取最大的值
func maxSubArray(nums []int) int {
	result := nums[0]

	for i := 0; i < len(nums); i++ {
		count := 0
		for j := i; j < len(nums); j++ {
			count += nums[j]
			if result < count {
				result = count
			}
		}
	}

	return result
}

//贪心算法
func maxSubArray1(nums []int) int {

	var sum int
	result := nums[0]
	for i := 0; i < len(nums); i++ {
		sum += nums[i]
		if sum >= result {
			result = sum
		}

		if sum < 0 {
			sum = 0
		}
	}

	return result
}

//贪心算法
// fmt.Println(maxSubArray2([]int{-2, 1, -3, 4, -1, 2, 1, -5, 4}))
func maxSubArray2(nums []int) []int {

	var sum int
	result := nums[0]
	var res []int
	for i := 0; i < len(nums); i++ {
		sum += nums[i]
		if sum >= result {
			result = sum
		}
		if sum < 0 {
			sum = 0
			res = make([]int, 0)
		} else {
			res = append(res, nums[i])
		}
	}
	return res
}
